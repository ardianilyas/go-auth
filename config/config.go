package config

import (
    "os"
    "log"
    "strconv"

    "github.com/joho/godotenv"
)

var (
    DB_DSN      string
    JWT_SECRET  string
    ACCESS_EXP  int
    REFRESH_EXP int
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    DB_DSN = os.Getenv("DB_DSN")
    JWT_SECRET = os.Getenv("JWT_SECRET")

    ACCESS_EXP, _ = strconv.Atoi(os.Getenv("ACCESS_EXP"))
    if ACCESS_EXP == 0 {
        ACCESS_EXP = 15
    }

    REFRESH_EXP, _ = strconv.Atoi(os.Getenv("REFRESH_EXP"))
    if REFRESH_EXP == 0 {
        REFRESH_EXP = 7
    }
}
