package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file\n%v", err)
	}

	fmt.Println("#### .env data ####")
	fmt.Println("TOKEN       :", os.Getenv("Token"))
	fmt.Println("MYSQL_URL   :", os.Getenv("MYSQL_URL"))
	fmt.Println("PAGE_ROW    :", os.Getenv("PAGE_ROW"))
	fmt.Println("ROLE_FORMAT :", os.Getenv("ROLE_FORMAT"))
	fmt.Println("## .env data end ##")

}
