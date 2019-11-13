package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// used for source filename and sqlite article table name
	tableName             = "DataBlog"
	sqliteExtension       = ".sqlite3"
	triggerNameUpdateDate = "UpdateDate"
)

var (
	sqliteDatabasePath = "./" + tableName + sqliteExtension
)

type Source interface {
	GetArticleWhere(typ ArticleType, id string) ([]article, error)

	// GetArticle(articlePropertiesMatching article) ([]article, error)

	// constrains on data input exist through sqlite table definition
	InsertArticle(art *article) error

	// UpdateArticle(articlePropertiesMatching article) error
	// DeleteArticle(art article) error
}

type sqlite struct {
	path string
}

//my attempt of a simple insert into sql:

func InsertNewArticle(id int, author string, title string, content string, indraft bool, date string, public bool) (bool, error) {
	db, err := sql.Open("sqlite3", sqliteDatabasePath)
	if err != nil {
		return false, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("insert into " + tableName + "(id, author, title, content, indraft, date, public) values(?,?,?,?,?,?,?)")
	defer stmt.Close()

	_, err = stmt.Exec(id, author, title, content, indraft, date, public)
	if err != nil {
		return false, err
	}

	tx.Commit()
	return true, nil
}

// Tries to create a new sqlite datasource. If it already exists
func NewSqlite() (Source, error) {
	sqlDb := &sqlite{path: sqliteDatabasePath}
	err := sqlDb.execSqlite(p_CreateTableIfNotExist())
	if err != nil {
		return sqlDb, err
	}
	return sqlDb, sqlDb.createTriggerUpdateDateIfNotExist()
}

// Generic function to get an array of matching articles.
// Fires two sql statements. One to get the count of matching rows and the second one to write to the array.
func (s *sqlite) GetArticleWhere(typ ArticleType, val string) ([]article, error) {
	rows, err := s.querySqlite(p_ArticleBy(typ, val))
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	nMatchingArticles, err := s.getCountWhere(typ, val)
	if err != nil {
		return nil, err
	}

	articles := make([]article, 0, nMatchingArticles)
	for rows.Next() {
		articleArray := make([]string, getArticleLength())
		err = rows.Scan(interface{}(articleArray))
		if err != nil {
			return nil, err
		}

		articles = append(articles, *NewArticleWithArray(articleArray))
	}

	return articles, nil
}

func (s *sqlite) InsertArticle(toInsert *article) error {
	panic("p_InsertArticle returns values without \" which needs to be fixed for insert to work - i guess")
	return s.execSqlite(p_InsertArticle(toInsert))
}

func (s *sqlite) getCountWhere(typ ArticleType, val string) (count int, err error) {
	rows, err := s.querySqlite(p_CountBy(typ, val))
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count, err
}

func (s *sqlite) checkTriggerExists(triggerName string) (bool, error) {
	var count int
	rows, err := s.querySqlite(p_CheckTriggerExists(triggerName))
	defer rows.Close()
	if err != nil {
		return false, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return false, err
		}
	}

	return count > 0, nil
}

func (s *sqlite) createTriggerUpdateDateIfNotExist() error {
	exist, err := s.checkTriggerExists(triggerNameUpdateDate)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	return s.execSqlite(p_TriggerUpdateDateOnChange(triggerNameUpdateDate))
}

// generic {{{

func (s *sqlite) execSqlite(statement string) error {
	db, err := sql.Open("sqlite3", s.path)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(statement)
	return err
}

func (s *sqlite) querySqlite(statement string) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", s.path)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db.Query(statement)
}

// }}}

// vim: set fdn=1 fdm=indent:
