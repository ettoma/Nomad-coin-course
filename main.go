package main

import (
	// "github.com/ettoretoma/Nomad-coin-course/explorer"
	"github.com/ettoretoma/Nomad-coin-course/rest"
)

func main() {
	// go explorer.Start(3000)
	rest.Start(8080)
}
