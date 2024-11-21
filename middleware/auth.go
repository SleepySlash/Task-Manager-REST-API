package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CreateToken(userid string) (string, error) {
	log.Println("Creating token for user:", userid)
	secretKey := os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userid": userid,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetIdFromContext(ctx context.Context) (string, error) {
	id, ok := ctx.Value("userid").(string)
	if !ok {
		return "", fmt.Errorf("user id not found in the context")
	}
	return id, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Forbidden", http.StatusUnauthorized)
			return
		}

		secretKey := os.Getenv("SECRET_KEY")
		log.Println("Verifying token:", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			log.Println("Token parsing error:", err)
			http.Error(w, "Forbidden", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Println("Token is invalid")
			http.Error(w, "Forbidden", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Println("Invalid token claims")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract userid from claims
		userID, ok := claims["userid"].(string)
		if !ok {
			log.Println("UserID missing in token claims")
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		// Add userid to context
		ctx := context.WithValue(r.Context(), "userid", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func DataSource() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading the env")
		return nil
	}
	log.Println("Loaded the env")

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("error in connecting to the db")
	}
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}
	log.Println("Ping Succesfull")
	log.Println("Connection Established with MongoDB")
	return mongoClient
}
