package configs

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Config struct holds the configuration data
type Config struct {
    ServerAddress string
}

// ProvideConfig returns the configuration settings
func ProvideConfig() *Config {
    return &Config{
        ServerAddress: ":8080",
    }
}

// ProvideDB sets up the database connection using GORM
func ProvideDB() (*gorm.DB, error) {
    dsn := "host=localhost user=postgres password=mysecretpassword dbname=mydb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to the database: %v", err)
        return nil, err
    }
    return db, nil
}
