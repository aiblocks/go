package test

import (
	"io"
	"testing"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/aiblocks/go/services/millennium/internal/db2/schema"
	"github.com/aiblocks/go/services/millennium/internal/ledger"
	"github.com/aiblocks/go/services/millennium/internal/operationfeestats"
	"github.com/aiblocks/go/services/millennium/internal/test/scenarios"
	"github.com/aiblocks/go/support/db"
	"github.com/aiblocks/go/support/render/hal"
)

// CoreSession returns a db.Session instance pointing at the aiblocks core test database
func (t *T) CoreSession() *db.Session {
	return &db.Session{
		DB:  t.CoreDB,
		Ctx: t.Ctx,
	}
}

// Finish finishes the test, logging any accumulated millennium logs to the logs
// output
func (t *T) Finish() {
	RestoreLogger()
	// Reset cached ledger state
	ledger.SetState(ledger.State{})
	operationfeestats.ResetState()

	if t.LogBuffer.Len() > 0 {
		t.T.Log("\n" + t.LogBuffer.String())
	}
}

// MillenniumSession returns a db.Session instance pointing at the millennium test
// database
func (t *T) MillenniumSession() *db.Session {
	return &db.Session{
		DB:  t.MillenniumDB,
		Ctx: t.Ctx,
	}
}

func (t *T) loadScenario(scenarioName string, includeMillennium bool) {
	aiblocksCorePath := scenarioName + "-core.sql"

	scenarios.Load(AiBlocksCoreDatabaseURL(), aiblocksCorePath)

	if includeMillennium {
		millenniumPath := scenarioName + "-millennium.sql"
		scenarios.Load(DatabaseURL(), millenniumPath)
	}
}

// Scenario loads the named sql scenario into the database
func (t *T) Scenario(name string) *T {
	clearMillenniumDB(t.T, t.MillenniumDB)
	t.loadScenario(name, true)
	t.UpdateLedgerState()
	return t
}

// ScenarioWithoutMillennium loads the named sql scenario into the database
func (t *T) ScenarioWithoutMillennium(name string) *T {
	t.loadScenario(name, false)
	ResetMillenniumDB(t.T, t.MillenniumDB)
	t.UpdateLedgerState()
	return t
}

// ResetMillenniumDB sets up a new millennium database with empty tables
func ResetMillenniumDB(t *testing.T, db *sqlx.DB) {
	clearMillenniumDB(t, db)
	_, err := schema.Migrate(db.DB, schema.MigrateUp, 0)
	if err != nil {
		t.Fatalf("could not run migrations up on test db: %v", err)
	}
}

func clearMillenniumDB(t *testing.T, db *sqlx.DB) {
	_, err := schema.Migrate(db.DB, schema.MigrateDown, 0)
	if err != nil {
		t.Fatalf("could not run migrations down on test db: %v", err)
	}
}

// UnmarshalPage populates dest with the records contained in the json-encoded page in r
func (t *T) UnmarshalPage(r io.Reader, dest interface{}) hal.Links {
	var env struct {
		Embedded struct {
			Records json.RawMessage `json:"records"`
		} `json:"_embedded"`
		Links struct {
			Self hal.Link `json:"self"`
			Next hal.Link `json:"next"`
			Prev hal.Link `json:"prev"`
		} `json:"_links"`
	}

	err := json.NewDecoder(r).Decode(&env)
	t.Require.NoError(err, "failed to decode page")

	err = json.Unmarshal(env.Embedded.Records, dest)
	t.Require.NoError(err, "failed to decode records")

	return env.Links
}

// UnmarshalNext extracts and returns the next link
func (t *T) UnmarshalNext(r io.Reader) string {
	var env struct {
		Links struct {
			Next struct {
				Href string `json:"href"`
			} `json:"next"`
		} `json:"_links"`
	}

	err := json.NewDecoder(r).Decode(&env)
	t.Require.NoError(err, "failed to decode page")
	return env.Links.Next.Href
}

// UnmarshalExtras extracts and returns extras content
func (t *T) UnmarshalExtras(r io.Reader) map[string]string {
	var resp struct {
		Extras map[string]string `json:"extras"`
	}

	err := json.NewDecoder(r).Decode(&resp)
	t.Require.NoError(err, "failed to decode page")

	return resp.Extras
}

// UpdateLedgerState updates the cached ledger state (or panicing on failure).
func (t *T) UpdateLedgerState() {
	var next ledger.State

	err := t.CoreSession().GetRaw(&next, `
		SELECT
			COALESCE(MAX(ledgerseq), 0) as core_latest
		FROM ledgerheaders
	`)

	if err != nil {
		panic(err)
	}

	err = t.MillenniumSession().GetRaw(&next, `
			SELECT
				COALESCE(MIN(sequence), 0) as history_elder,
				COALESCE(MAX(sequence), 0) as history_latest
			FROM history_ledgers
		`)

	if err != nil {
		panic(err)
	}

	ledger.SetState(next)
}
