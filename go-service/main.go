package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("--- Go Service ---")

	createHTTPServer()
}

type SearchBody struct {
	Search string `json:"search"`
}

func createHTTPServer() {
	http.HandleFunc("/user-search", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			return
		}

		decoder := json.NewDecoder(r.Body)

		var body SearchBody
		err := decoder.Decode(&body)
		if err != nil {
			log.Fatalln("Error reading body:", err)
		}

		fmt.Printf("Search request for: %v\n", body.Search)

		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.WriteHeader(http.StatusOK)

		res, err := json.Marshal(map[string]string{"search": body.Search})

		if err != nil {
			log.Fatalln("Error marshalling response:", err)
		}

		_, err = rw.Write(res)
		if err != nil {
			log.Fatalln("Error writing response:", err)
		}
	})

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
