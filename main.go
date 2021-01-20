// main.go

package main

import ("os"
	"fmt")

func main() {
	fmt.Printf("Welcome to Cosmos !\n")
	app := App{}
	app.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	app.Run(":9100")
}
