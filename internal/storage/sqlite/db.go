// Package sqlite is the home for a SQLite-backed implementation of whatever
// storage interface you decide to define.
//
// We've provided one helper — Open — that opens a SQLite file and applies
// the schema in schema.sql. Everything else (queries, types, transactions)
// is up to you.
//
// If you choose the in-memory path instead, you can delete this directory.
package sqlite

import (
	"database/sql"
	_ "embed"
	"fmt"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

// Open opens (or creates) a SQLite database at path and applies the schema.
// Use ":memory:" for an ephemeral in-process database, which is handy for
// tests.
func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}
	if _, err := db.Exec(schema); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("apply schema: %w", err)
	}
	return db, nil
}
