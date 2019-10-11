package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dataRootName       = "DataBlog"
	sqliteExtension    = "sqlite3"
	sqliteDatabasePath = "./" + dataRootName + sqliteExtension
)

type Source interface {
	Setup() error
	// GetArticleWithId(id string) Article
	GetArticleWithId(id string) string
}

type sqlite struct {
	path string
}

type Article struct {
	Id      int
	Author  string
	Title   string
	Content string
	InDraft int
	Date    string
	Public  int
}

func NewSqlite() Source {
	return &sqlite{path: sqliteDatabasePath}
}

func GAWI(id string) string {
	return GetArticleWithId(id)
}

func GetArticleWithId(id string) string {
	db, err := sql.Open("sqlite3", sqlite.path)
	if err != nil {
		return ""
	}
	defer db.Close()
	rows, err := db.Query(initArticleWithIdStatement(id))

	if err != nil {
		return ""
	}

	article := Article{}
	for rows.Next() {
		err2 := rows.Scan(&article.Id, &article.Author, &article.Title, &article.Content, &article.InDraft, &article.Date, &article.Public)
		if err2 != nil {
			return ""
		}
		fmt.Println(article.Id)
	}

	return article.Content
}

func (s *sqlite) Setup() error {
	db, err := sql.Open("sqlite3", s.path)
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare(sqliteArticleTableDefinition.getSqliteTableCreationStatement())
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	return err
}

func initArticleWithIdStatement(id string) string {
	return "SELECT * FROM DataBlog WHERE id = " + id
}
