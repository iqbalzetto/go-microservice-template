package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"user-service/pkg/config"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Global DB connection pool
var db *sql.DB

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

	// Test the connection
	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		return nil, fmt.Errorf("failed to ping DB: %v", err)
	}

	log.Println("✅ Successfully connected to the PostgreSQL database")
	return dbConn, nil
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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

	//SET UP SERVER
	e := echo.New()
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Users!")

	})
	e.Logger.Fatal(e.Start(":4001"))
}
