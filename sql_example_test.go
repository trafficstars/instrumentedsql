package instrumentedsql_test

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mattn/go-sqlite3"

	"github.com/trafficstars/instrumentedsql"
	"github.com/trafficstars/instrumentedsql/opentracing"
)

// WrapDriverOpentracing demonstrates how to call wrapDriver and register a new driver.
// This example uses MySQL and opentracing to illustrate this
func ExampleWrapDriver_opentracing() {
	logger := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	})

	sql.Register("instrumented-mysql", instrumentedsql.WrapDriver(mysql.MySQLDriver{}, instrumentedsql.WithTracer(opentracing.NewTracer()), instrumentedsql.WithLogger(logger)))
	db, err := sql.Open("instrumented-mysql", "connString")

	// Proceed to handle connection errors and use the database as usual
	_, _ = db, err
}

// WrapDriverJustLogging demonstrates how to call wrapDriver and register a new driver.
// This example uses sqlite, but does not trace, but merely logs all calls
func ExampleWrapDriver_justLogging() {
	logger := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
		log.Printf("%s %v", msg, keyvals)
	})

	sql.Register("instrumented-sqlite", instrumentedsql.WrapDriver(&sqlite3.SQLiteDriver{}, instrumentedsql.WithLogger(logger)))
	db, err := sql.Open("instrumented-sqlite", "connString")

	// Proceed to handle connection errors and use the database as usual
	_, _ = db, err
}
