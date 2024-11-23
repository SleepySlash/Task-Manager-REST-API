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

// Create a jwt token while user login and write it onto the header for authentication purposes
func CreateToken(userid string) (string, error) {
	log.Println("Creating token for user:", userid)
	secretKey := os.Getenv("SECRET_KEY")
	claims := jwt.MapClaims{
		"userid": userid,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Define a custom type for context keys
type contextKey string

const userIDKey contextKey = "userid"

// Get userid from the jwt token
func GetIdFromContext(ctx context.Context) (string, error) {
	id, ok := ctx.Value(userIDKey).(string)
	if !ok {
		return "", fmt.Errorf("user id not found in the context")
	}
	return id, nil
}

// Authentication function to be used by the router as middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Forbidden", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix from the token string
		tokenString := authHeader[len("Bearer "):]
		if tokenString == authHeader {
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
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if !token.Valid {
			log.Println("Token is invalid")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Extract userid from claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Println("Invalid token claims")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		userID, ok := claims["userid"].(string)
		if !ok {
			log.Println("UserID missing in token claims")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Add userid to context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Establish the database connection
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

// Request Logger
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf(" Method: %s\t URL: %s\t Time: %s\n RemoteAddr: %s\t UserAgent: %s\n", r.Method, r.URL, start.Format(time.RFC1123), r.RemoteAddr, r.UserAgent())
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v\n", time.Since(start))
	})
}
