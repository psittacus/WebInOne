package data

import "fmt"
import "strings"

var (
	columnDefinitions = map[ArticleType]string{
		Id:      "INTEGER PRIMARY KEY",
		Author:  "TEXT NOT NULL",
		Title:   "TEXT NOT NULL",
		Content: "TEXT",
		InDraft: "INTEGER DEFAULT 1 NOT NULL",
		Date:    "DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL",
		Public:  "INTEGER DEFAULT 1"}
)

// procedures {{{

func p_CreateTableIfNotExist() string {
	tableDefinition := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tableName)
	pre := ""
	for _, typ := range sortedArticleTypes {
		tableDefinition += fmt.Sprintf("%s%s %s", pre, typ, columnDefinitions[typ])
		pre = ", "
	}

	tableDefinition += ");"
	return tableDefinition
}

func p_InsertArticle(toInsert *article) string {
	return fmt.Sprintf("INSERT INTO %s (%s), VALUES (%s);", tableName, strings.Join(getTypesArticle(), ", "), strings.Join((toInsert).getValuesArticle(), ", "))
}

func p_ArticleBy(typ ArticleType, val string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE %s = %s;", tableName, typ, val)
}

func p_CountBy(typ ArticleType, val string) string {
	return fmt.Sprintf("SELECT COUNT(*) as count FROM %s WHERE %s = %s;", tableName, typ, val)
}

func p_CheckTriggerExists(triggerName string) string {
	return fmt.Sprintf("SELECT Count(*) as count FROM sqlite_master WHERE type='trigger' AND name='%s'", triggerName)
}

func p_TriggerUpdateDateOnChange(triggerName string) string {
	return fmt.Sprintf("CREATE TRIGGER %[1]s BEFORE UPDATE ON %[2]s FOR EACH ROW WHEN NEW.date <= OLD.date BEGIN UPDATE %[2]s SET date=CURRENT_TIMESTAMP where id=old.id; END;", triggerName, tableName)
}

// vim: fdn=1 fdm=marker:
