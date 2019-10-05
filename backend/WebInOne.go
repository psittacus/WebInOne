package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/psittacus/WebInOne/backend/datasource"
	"log"
	"net/http"
)

func PostIDHandler(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	writer.WriteHeader(http.StatusOK)

	response := datasource.GetEverythingFromID(vars["id"])

	fmt.Fprintf(writer, response)
}

func main() {
	//	fmt.Println(datasource.GetArticleWithID("ID"))

	router := mux.NewRouter()
	router.HandleFunc("/post/{id:[0-9]+}", PostIDHandler).
		Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
