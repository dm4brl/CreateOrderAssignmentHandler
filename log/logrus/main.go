package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    // ... (other imports)

    "github.com/go-redis/redis/v8"
    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
    "github.com/spf13/viper"
    "golang.org/x/net/context"
)

var (
    redisClient *redis.Client
    db          *sql.DB
)

// ...

func main() {
    // Initialize the log file
    logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer logFile.Close()

    // Set the log output to the log file
    log.SetOutput(logFile)

    // Load configuration using Viper
    viper.SetConfigName("config")  // Load a configuration file named config.yaml or config.json, etc.
    viper.AddConfigPath("config/") // Path to the directory where your config file is located
    viper.ReadInConfig()

    // ...

    // Create a PostgreSQLDatabaseService instance
    databaseService := NewPostgreSQLDatabaseService(db)

    r := mux.NewRouter()

    // Define a route for updating courier information
    r.HandleFunc("/update_courier_info", UpdateCourierInfo(databaseService)).Methods("POST")

    http.Handle("/", r)

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// UpdateCourierInfo is the handler for updating courier information.
func UpdateCourierInfo(dbService *PostgreSQLDatabaseService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Log request details
        log.Printf("Received request: %s %s", r.Method, r.URL.Path)

        // Parse the JSON body containing courier information
        var updatedCourier Courier
        err := json.NewDecoder(r.Body).Decode(&updatedCourier)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            log.Printf("Failed to parse JSON: %v", err)
            return
        }

        // ...

        // Log the successful update
        log.Printf("Courier information updated: ID=%d, Name=%s", updatedCourier.ID, updatedCourier.Name)

        // Return a response, e.g., confirming the update
        response := map[string]interface{}{
            "message":      "Courier information updated successfully",
            "updated_data": updatedCourier,
        }

        w.Header().Set("Content-Type", "application/json)
        json.NewEncoder(w).Encode(response)
    }
}
