package main

import (
	"log"
	"net/http"

	"github.com/fencedin/csv2json"
)

var (
	columns = []string{"Name", "Date", "Title"}
)

func csv2JsonServer(w http.ResponseWriter, req *http.Request) {

	log.Println(req)

	jsonData, err := csv2json.Convert(req.Body, columns)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, "Could not convert csv to json", 503)
	}
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/csv2json", csv2JsonServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
