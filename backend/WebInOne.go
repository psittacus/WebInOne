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

	author := datasource.GetAuthorFromID(vars["id"])
	title := datasource.GetTitleFromID(vars["id"])
	article := datasource.GetArticleFromID(vars["id"])
	inDraft := datasource.GetInDraftFromID(vars["id"])
	date := datasource.GetDateFromID(vars["id"])
	public := datasource.GetPublicFromID(vars["id"])

	fmt.Fprintf(writer, "Argument: %v", vars["id"])
}

func main() {
	//	fmt.Println(datasource.GetArticleWithID("ID"))

	router := mux.NewRouter()
	router.HandleFunc("/post/{id:[0-9]+}", PostIDHandler).
		Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
