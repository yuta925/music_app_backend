package config

import (
	"log"
	"os"
	"time"
)

var (
	env                       string
	sigKey                    string // JWTトークンの署名
	jst                       *time.Location
)

func init() {
	sigKey = os.Getenv("SIG_KEY")
	if sigKey == "" {
		log.Print("SIG_KEY environment variable is empty")
	}
	env = os.Getenv("ENV")
	if env == "" {
		log.Print("ENV environment variable is empty")
	}
	if j, err := time.LoadLocation("Asia/Tokyo"); err != nil {
		log.Print("Failed to load location")
	} else {
		jst = j
	}
}

func IsDevelopment() bool {
	return env == "development"
}

func SigKey() string {
	return sigKey
}


func JST() *time.Location {
	return jst
}
