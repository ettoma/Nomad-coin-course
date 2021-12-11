package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/ettoretoma/Nomad-coin-course/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on port http://localhost%s", port)
	http.ListenAndServe(port, nil)
	log.Fatal(http.ListenAndServe(port, nil))
}
