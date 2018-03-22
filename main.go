package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func LogRequest(request *http.Request) {
	fmt.Printf("%s %s %q\n", time.Now(), request.Method, request.URL.Path)
}

func CreateWalletEndpoint(response http.ResponseWriter, request *http.Request) {
	LogRequest(request)
	response.Header().Set("Content-Type", "application/json")
	queryParams := request.URL.Query()
	var wallet Wallet
	json.NewDecoder(request.Body).Decode(&wallet)
	if queryParams.Get("key") == "" {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "A password is required to create a wallet" }`))
		return
	}
	wallet.Create(queryParams.Get("key"))
	json.NewEncoder(response).Encode(wallet)
}

func DestroyWalletEndpoint(response http.ResponseWriter, request *http.Request) {
	LogRequest(request)
	response.Header().Set("Content-Type", "application/json")
	var wallet Wallet
	err := wallet.Destroy()
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(wallet)
}

func DumpWalletEndpoint(response http.ResponseWriter, request *http.Request) {
	LogRequest(request)
	response.Header().Set("Content-Type", "application/json")
	queryParams := request.URL.Query()
	if queryParams.Get("key") == "" {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "A password is required to decrypt a wallet" }`))
		return
	}
	var wallet Wallet
	err := wallet.Dump(queryParams.Get("key"))
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(wallet)
}

func ImportCoinEndpoint(response http.ResponseWriter, request *http.Request) {
	LogRequest(request)
	response.Header().Set("Content-Type", "application/json")
	queryParams := request.URL.Query()
	var coin Coin
	json.NewDecoder(request.Body).Decode(&coin)
	err := coin.Import(network[strings.ToLower(coin.Symbol)])
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	var wallet Wallet
	err = wallet.DecryptFile(queryParams.Get("key"))
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "The password is not correct" }`))
		return
	}
	err = wallet.Import(coin, queryParams.Get("key"))
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(coin)
}

func CreateCoinEndpoint(response http.ResponseWriter, request *http.Request) {
	LogRequest(request)
	response.Header().Set("Content-Type", "application/json")
	var coin Coin
	json.NewDecoder(request.Body).Decode(&coin)
	queryParams := request.URL.Query()
	if queryParams.Get("key") == "" {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "A password is required to use a wallet" }`))
		return
	}
	coin.Generate(network[strings.ToLower(coin.Symbol)])
	var wallet Wallet
	err := wallet.DecryptFile(queryParams.Get("key"))
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "The password is not correct" }`))
		return
	}
	err = wallet.Import(coin, queryParams.Get("key"))
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(coin)
}

func AuthenticateWalletEndpoint(response http.ResponseWriter, request *http.Request) {
	LogRequest(request)
	response.Header().Set("Content-Type", "application/json")
	queryParams := request.URL.Query()
	if queryParams.Get("key") == "" {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "A password is required to decrypt a wallet" }`))
		return
	}
	var wallet Wallet
	authenticated := wallet.Authenticate(queryParams.Get("key"))
	response.Write([]byte(`{ "authenticated": ` + strconv.FormatBool(authenticated) + ` }`))
}

func GetAddressesEndpoint(response http.ResponseWriter, request *http.Request) {
	LogRequest(request)
	response.Header().Set("Content-Type", "application/json")
	queryParams := request.URL.Query()
	if queryParams.Get("key") == "" {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "A password is required to decrypt a wallet" }`))
		return
	}
	var wallet Wallet
	err := wallet.GetAddresses(queryParams.Get("key"))
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(wallet)
}

func ExportWalletEndpoint(response http.ResponseWriter, request *http.Request) {
	LogRequest(request)
	response.Header().Set("Content-Disposition", "attachment; filename=wallet.dat")
	file, err := ioutil.ReadFile("wallet.dat")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	response.Write(file)
}

func CreateTransactionEndpoint(response http.ResponseWriter, request *http.Request) {
	LogRequest(request)
	response.Header().Set("Content-Type", "application/json")
	var transaction Transaction
	json.NewDecoder(request.Body).Decode(&transaction)
	queryParams := request.URL.Query()
	if queryParams.Get("key") == "" {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "A password is required to decrypt a wallet" }`))
		return
	}
	var wallet Wallet
	err := wallet.Dump(queryParams.Get("key"))
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	for _, coin := range wallet.Coins {
		if coin.UncompressedAddress == transaction.SourceAddress || coin.CompressedAddress == transaction.SourceAddress {
			tx, err := CreateTransaction(network[strings.ToLower(coin.Symbol)], coin.WIF, transaction.DestinationAddress, transaction.Amount, transaction.TxId)
			if err != nil {
				response.WriteHeader(500)
				response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
				return
			}
			json.NewEncoder(response).Encode(tx)
			return
		}
	}
	response.WriteHeader(500)
	response.Write([]byte(`{ "message": "The source address does not exist in the wallet" }`))
}

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/wallet", CreateWalletEndpoint).Methods("POST")
	api.HandleFunc("/wallet", DestroyWalletEndpoint).Methods("DELETE")
	api.HandleFunc("/wallet", DumpWalletEndpoint).Methods("GET")
	api.HandleFunc("/addresses", GetAddressesEndpoint).Methods("GET")
	api.HandleFunc("/authenticate", AuthenticateWalletEndpoint).Methods("POST")
	api.HandleFunc("/import-coin", ImportCoinEndpoint).Methods("POST")
	api.HandleFunc("/create-coin", CreateCoinEndpoint).Methods("POST")
	api.HandleFunc("/backup", ExportWalletEndpoint).Methods("GET")
	api.HandleFunc("/transaction", CreateTransactionEndpoint).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("ui/dist").HTTPBox()))
	fmt.Println("Listening at :12345...")
	log.Fatal(http.ListenAndServe(":12345", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
