package sql

import "database/sql"

type Execer interface {
	Exec(query string, args ...any) (sql.Result, error)
}

func execQuery(db Execer, query string, args ...interface{}) (int64, error) {
	res, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return ra, nil
}
