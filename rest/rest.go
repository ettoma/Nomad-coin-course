package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ettoretoma/Nomad-coin-course/blockchain"
	"github.com/ettoretoma/Nomad-coin-course/utils"
	"github.com/gorilla/mux"
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

type BalanceByAddress struct {
	Address string
	Balance int
}

type addTxPayload struct {
	To     string
	Amount int
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDocumentation{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See documentation",
		},
		{
			URL:         URL("/stats"),
			Method:      "GET",
			Description: "See status of the blockchain",
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
		{
			URL:         URL("/balance/{address}"),
			Method:      "GET",
			Description: "Get TxOuts per address",
		},
	}

	rw.Header().Add("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(data)

	//! it can be also be constructed like this:
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := blockchain.Blockchain().Blocks()
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(data)

	case "POST":
		blockchain.Blockchain().AddBlock()
		rw.WriteHeader(http.StatusCreated)

	}

}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)

	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {

		json.NewEncoder(rw).Encode(block)
	}

}

func stats(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(blockchain.Blockchain())

}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")
	switch total {
	case "true":
		amount := blockchain.Blockchain().BalanceByAddress(address)
		json.NewEncoder(rw).Encode(BalanceByAddress{
			address, amount,
		})
	default:
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain().UTxOutsByAddress(address)))
	}

}

func mempool(rw http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Mempool.Txs))
}

func transactions(rw http.ResponseWriter, r *http.Request) {
	var payload addTxPayload
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	err := blockchain.Mempool.AddTx(payload.To, payload.Amount)
	if err != nil {
		json.NewEncoder(rw).Encode(errorResponse{"not enough funds"})
	}
	rw.WriteHeader(http.StatusCreated)
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(port int) {
	router := mux.NewRouter()
	PORT = fmt.Sprintf(":%d", port)
	fmt.Println("Listening on http://localhost" + PORT)

	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	router.HandleFunc("/stats", stats).Methods("GET")
	router.HandleFunc("/balance/{address}", balance).Methods("GET")
	router.HandleFunc("/mempool", mempool).Methods("GET")
	router.HandleFunc("/transaction", transactions).Methods("POST")
	router.HandleFunc("/", documentation).Methods("GET")
	log.Fatal(http.ListenAndServe(PORT, router))
}
