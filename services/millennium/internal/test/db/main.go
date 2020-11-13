// Package db provides helpers to connect to test databases.  It has no
// internal dependencies on millennium and so should be able to be imported by
// any millennium package.
package db

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	// pq enables postgres support
	_ "github.com/lib/pq"
	db "github.com/aiblocks/go/support/db/dbtest"
)

var (
	coreDB     *sqlx.DB
	coreUrl    *string
	millenniumDB  *sqlx.DB
	millenniumUrl *string
)

// Millennium returns a connection to the millennium test database
func Millennium(t *testing.T) *sqlx.DB {
	if millenniumDB != nil {
		return millenniumDB
	}
	postgres := db.Postgres(t)
	millenniumUrl = &postgres.DSN
	millenniumDB = postgres.Open()

	return millenniumDB
}

// MillenniumURL returns the database connection the url any test
// use when connecting to the history/millennium database
func MillenniumURL() string {
	if millenniumUrl == nil {
		log.Panic(fmt.Errorf("Millennium not initialized"))
	}
	return *millenniumUrl
}

// AiBlocksCore returns a connection to the aiblocks core test database
func AiBlocksCore(t *testing.T) *sqlx.DB {
	if coreDB != nil {
		return coreDB
	}
	postgres := db.Postgres(t)
	coreUrl = &postgres.DSN
	coreDB = postgres.Open()
	return coreDB
}

// AiBlocksCoreURL returns the database connection the url any test
// use when connecting to the aiblocks-core database
func AiBlocksCoreURL() string {
	if coreUrl == nil {
		log.Panic(fmt.Errorf("AiBlocksCore not initialized"))
	}
	return *coreUrl
}
