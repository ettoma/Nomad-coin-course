package main

import (
	"github.com/ettoretoma/Nomad-coin-course/db"
	"github.com/ettoretoma/Nomad-coin-course/rest"
)

// import (
// "github.com/ettoretoma/Nomad-coin-course/explorer"
// "github.com/ettoretoma/Nomad-coin-course/rest"
// )

// func usage() {
// 	fmt.Println("Welcome to the coin")
// 	fmt.Println("Please use one of the following commands: ")
// 	fmt.Println("explorer: start the HTML explorer")
// 	fmt.Println("rest: start the REST API")

// 	os.Exit(0)
// }

func main() {
	defer db.Close()
	// go explorer.Start(3000)
	// blockchain.Blockchain().AddBlock("new block")
	rest.Start(8000)
}
