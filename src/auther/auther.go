package auther

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"jandre.dev/auther/src/hasher"
)

func GetPassword() {
	envFile, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal(err)
	}

	authKey := envFile["AUTH_KEY"]

	fmt.Printf("Working on handling authKey for hashing algorithm: %v\n", authKey)

	hasher.Hasher("Something")
}
