package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getName() string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error to load env")
	}
	return os.Getenv("Name")

}

func IsProduction() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load env")
	}
	return os.Getenv("Production")
}

func MongoDBURI() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load env")
	}
	return os.Getenv("MONGODBURI")
}

func TokenLifeSpan() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load env")
	}
	return os.Getenv("TOKEN_HOUR_LIFESPAN")

}
