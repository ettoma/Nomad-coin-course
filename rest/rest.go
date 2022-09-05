package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ettoretoma/Nomad-coin-course/blockchain"
	"github.com/ettoretoma/Nomad-coin-course/utils"
)

var PORT string

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", PORT, u)
	return []byte(url), nil
}

type URLDocumentation struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDocumentation{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "GET",
			Description: "See all the blocks",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
		{
			URL:         URL("/blocks/{id}"),
			Method:      "GET",
			Description: "See a block",
		},
	}

	rw.Header().Add("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(data)

	//! it can be also be constructed like this:
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
}

type AddBlockBody struct {
	Message string
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := blockchain.GetBlockchain().AllBlocks()
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(data)

	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte(fmt.Sprintf("well done jedi, your data is: %s", addBlockBody.Message)))

	}

}

func Start(port int) {
	handler := http.NewServeMux()
	PORT = fmt.Sprintf(":%d", port)
	fmt.Println("Listening on http://localhost" + PORT)
	handler.HandleFunc("/blocks", blocks)
	handler.HandleFunc("/", documentation)
	log.Fatal(http.ListenAndServe(PORT, handler))
}
