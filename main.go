package main

import(
	"fmt"
	"log"
	"os"
	"github.com/cloudflare/cloudflare-go"
)
func main()  {
	_ , err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
		if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Welcome to ooty!!")

}