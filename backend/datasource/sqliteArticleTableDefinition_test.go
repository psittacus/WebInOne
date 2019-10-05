package datasource

import (
	"database/sql"
	"testing"
)

func TestCreateDatabase(t *testing.T) {
	_, _ = datasource.setupDatasource()

}
