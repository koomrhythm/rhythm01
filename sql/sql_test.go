package sql

import (
	"database/sql"
	"testing"
)

type stub struct {
	query        string
	args         any
	lastInsertID int64
	rowsAffected int64
}

func (m *stub) LastInsertId() (int64, error) {
	return m.lastInsertID, nil
}

func (m *stub) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func (s *stub) Exec(query string, args ...any) (sql.Result, error) {
	s.query = query
	s.args = args
	return s, nil
}

func TestExecQuery(t *testing.T) {
	s := &stub{rowsAffected: 1}

	query := "SELECT * FROM TEST"
	r, err := execQuery(s, query, nil)
	if err != nil {
		t.Fatal(err)
	}

	if r != 1 {
		t.Fatalf("Expected 1, got %d", r)
	}

	if s.query != query {
		t.Fatalf("Expected %s, got %s", query, s.query)
	}
}
