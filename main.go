package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ettoretoma/Nomad-coin-course/utils"
)

const port string = ":4000"

type URLdescriptions struct {
	URL         string
	Method      string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLdescriptions{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See documentation",
		},
	}
	b, err := json.Marshal(data)
	utils.HandleError(err)
	fmt.Println(b)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listing on port http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
