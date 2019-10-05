package main

import (
	"fmt"
	"github.com/psittacus/WebInOne/backend/datasource"
)

func main() {
	fmt.Println(datasource.GetArticleWithID("ID"))
}
