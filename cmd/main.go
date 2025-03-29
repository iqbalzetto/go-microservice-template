package main

import (
	"database/sql"
	"fmt"
	"go-microservice-template/pkg/config"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-microservice-template/internal/app"
	router "go-microservice-template/internal/app/api/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Global DB connection pool
var db *sql.DB
var minioClient *minio.Client

func initDB() (*sql.DB, error) {
	// Load environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Validate required environment variables
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}

	// PostgreSQL DSN (Data Source Name)
	// Format: postgres://user:password@host:port/dbname?options
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		url.QueryEscape(dbUser), url.QueryEscape(dbPass), dbHost, dbPort, dbName)

	// Add additional parameters
	val := url.Values{}
	val.Add("sslmode", "disable") // Change as needed: "require", "verify-full", etc.
	val.Add("timezone", "Asia/Jakarta")
	dsn = fmt.Sprintf("%s?%s", dsn, val.Encode())

	// Open the database connection
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB connection: %v", err)
	}

	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(5)
	dbConn.SetConnMaxLifetime(60 * time.Minute)

	// Test the connection
	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		return nil, fmt.Errorf("failed to ping DB: %v", err)
	}

	log.Println("✅ Successfully connected to the PostgreSQL database")

	return dbConn, nil
}

func initMinio() (*minio.Client, error) {
	minioKey := os.Getenv("MINIO_KEY")
	minioSecret := os.Getenv("MINIO_SECRET")
	minioUrl := os.Getenv("MINIO_URL")
	var err error
	minioClient, err = minio.New(minioUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(minioKey, minioSecret, ""),
		Secure: false, // Set to true if using HTTPS
	})
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}
	fmt.Println("✅ MinIO client initialized")
	return minioClient, nil
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Graceful shutdown function
func handleShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("\nShutting down gracefully...")

		if db != nil {
			fmt.Println("Closing PostgreSQL connection...")
			db.Close()
		}

		fmt.Println("Cleanup complete. Exiting.")
		os.Exit(0)
	}()
}

func main() {
	//SET UP TRACER
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "staging" || appEnv == "production" {
		tracer.Start(
			tracer.WithEnv(appEnv),
			tracer.WithServiceName(config.GetModuleName()),
		)
		defer tracer.Stop()
	}

	//SET UP DATABASE CONNECTION
	var err error
	db, err = initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Ensure the DB connection is closed on shutdown

	//SET UP MINIO CLIENT
	minioClient, err = initMinio()
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}

	//SET UP HANDLER
	handlers := app.InitUserDomainHandler(db, minioClient)

	//SET UP ROUTER
	e := echo.New()
	router.InitRoutes(e, handlers)

	// Handle graceful shutdown
	handleShutdown()

	e.Logger.Fatal(e.Start(":4001"))
}
