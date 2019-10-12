package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/psittacus/WebInOne/backend/data"
	"log"
	"net/http"
)

/*
func DateBeforeHandler(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	writer.WriteHeader(http.StatusOK)

	response := datasource.GetEverythingFromAuthor(vars["date"])

	fmt.Fprintf(writer, response)
}

func DateHandler(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	writer.WriteHeader(http.StatusOK)

	response := datasource.GetEverythingFromAuthor(vars["date"])

	fmt.Fprintf(writer, response)
}

func AuthorHandler(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	writer.WriteHeader(http.StatusOK)

	response := datasource.GetEverythingFromAuthor(vars["author"])

	fmt.Fprintf(writer, response)
}
*/

func PostIDHandler(writer http.ResponseWriter, req *http.Request) {
	// datasource := Datasource{}
	// vars := mux.Vars(req)
	writer.WriteHeader(http.StatusOK)
	// response := datasource.GetArticleWithId(vars["id"])

	fmt.Fprintf(writer, "<response nimpl>")
}

func main() {
	source, err := data.NewSqlite()
	if err != nil {
		// for now...
		log.Fatal(err)
	}
	// article
	_, err = source.GetArticleWhere(data.Id, "1")

	router := mux.NewRouter()
	router.HandleFunc("/post/{id:[0-9]+}", PostIDHandler).
		Methods("GET")

	/*	router.HandleFunc("/date/before/{date:[1-31].[1-12].[2019-2119]}", DateBeforeHandler).
			Methods("GET")
		router.HandleFunc("/date/{date:[1-31].[1-12].[2019-2119]", DateHandler).
			Methods("GET")
		router.HandleFunc("/author/{author}", AuthorHandler).
			Methods("GET")
	*/

	log.Fatal(http.ListenAndServe(":8080", router))
	// _, _ = data.NewSqlite()
}
