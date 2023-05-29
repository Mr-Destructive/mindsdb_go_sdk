package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/connectors"
	"github.com/mr-destructive/mindsdb_go_sdk/mindsdb/rest_api"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	_ = mindsdb.LoadEnvFromFile(".env")
	email := os.Getenv("email")
	password := os.Getenv("password")

	api, _ := connectors.Login(email, password)
	fmt.Printf("%s\n", api.Url)
	body, _ := rest_api.GetDatabases(api)
	fmt.Printf("%s", body)
	body, _ = rest_api.DeleteDatabase(api, "test_1235")
	//body, _ := rest_api.CreateDatabase(api, map[string]string{"name": "test_123", "engine": "mindsdb"})
	fmt.Printf("%s", body)

}
