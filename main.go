package main

import(
	"fmt"
	"log"
	"os"
	"github.com/adfolks/liwa-nw/vm"
	"github.com/cloudflare/cloudflare-go"
	
)
//this varible can remain uninitilazied 
var localVariable string

func main()  {
	//this varible must be used
	var localVariable string
	_ , err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
		if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Welcome to ooty!!")
	vm.ReturnHai()

}