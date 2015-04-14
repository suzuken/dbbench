package dbbench_test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/suzuken/dummy"
	"log"
	"os"
	"testing"
)

func BenchmarkSQLiteReplaceInsertWithTx(b *testing.B) {
	os.Remove("./sqlite_bench.db")
	d := dummy.NewGenerator()

	db, err := sql.Open("sqlite3", "./sqlite_bench.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id string not null primary key, a string, s integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert or replace into foo(id, a, s) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = stmt.Exec(d.String(100), d.String(100), d.Int(13))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func BenchmarkSQLiteReplaceInsert(b *testing.B) {
	os.Remove("./sqlite_bench.db")
	d := dummy.NewGenerator()

	db, err := sql.Open("sqlite3", "./sqlite_bench.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id string not null primary key, a string, s integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	stmt, err := db.Prepare("insert or replace into foo(id, a, s) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = stmt.Exec(d.String(100), d.String(100), d.Int(13))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkSQLiteInsertWithTx(b *testing.B) {
	os.Remove("./sqlite_bench.db")
	d := dummy.NewGenerator()

	db, err := sql.Open("sqlite3", "./sqlite_bench.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id string not null primary key, a string, s integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into foo(id, a, s) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = stmt.Exec(d.String(100), d.String(100), d.Int(13))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func BenchmarkSQLiteInsert(b *testing.B) {
	os.Remove("./sqlite_bench.db")
	d := dummy.NewGenerator()

	db, err := sql.Open("sqlite3", "./sqlite_bench.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id string not null primary key, a string, s integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	stmt, err := db.Prepare("insert into foo(id, a, s) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = stmt.Exec(d.String(100), d.String(100), d.Int(13))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// fetch by pkey, scan without index, with index, etc..
func BenchmarkSQLiteReadPrimaryKey(b *testing.B) {
	os.Remove("./sqlite_bench.db")
	d := dummy.NewGenerator()

	db, err := sql.Open("sqlite3", "./sqlite_bench.db")
	if err != nil {
		log.Fatalf("cannot create db object: %s", err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id string not null primary key, a string, s integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("[ERR] cannot start transaction: %s", err)
	}
	stmt, err := db.Prepare("insert into foo(id, a, s) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < 10000; i++ {
		_, err = stmt.Exec(d.String(100), d.String(100), d.Int(13))
		if err != nil {
			log.Fatalf("failing insertion: %s", err)
		}
	}
	tx.Commit()

	b.ResetTimer()
	sl, err := db.Prepare("select * from foo where id=?")
	if err != nil {
		log.Fatalf("failing prepare query: %s", err)
	}
	defer sl.Close()

	for i := 0; i < b.N; i++ {
		var id string
		var a string
		var s int
		_ = sl.QueryRow(d.String(100)).Scan(&id, &a, &s)
	}
}

func BenchmarkSQLiteReadKeyWithoutIndex(b *testing.B) {
	os.Remove("./sqlite_bench.db")
	d := dummy.NewGenerator()

	db, err := sql.Open("sqlite3", "./sqlite_bench.db")
	if err != nil {
		log.Fatalf("cannot create db object: %s", err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id string not null primary key, a string, s integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("[ERR] cannot start transaction: %s", err)
	}
	stmt, err := db.Prepare("insert into foo(id, a, s) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < 10000; i++ {
		_, err = stmt.Exec(d.String(100), d.String(100), d.Int(13))
		if err != nil {
			log.Fatalf("failing insertion: %s", err)
		}
	}
	tx.Commit()

	b.ResetTimer()
	sl, err := db.Prepare("select * from foo where a=?")
	if err != nil {
		log.Fatalf("failing prepare query: %s", err)
	}
	defer sl.Close()

	for i := 0; i < b.N; i++ {
		var id string
		var a string
		var s int
		_ = sl.QueryRow(d.String(100)).Scan(&id, &a, &s)
	}
}

func BenchmarkSQLiteReadKeyWithIndex(b *testing.B) {
	os.Remove("./sqlite_bench.db")
	d := dummy.NewGenerator()

	db, err := sql.Open("sqlite3", "./sqlite_bench.db")
	if err != nil {
		log.Fatalf("cannot create db object: %s", err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id string not null primary key, a string, s integer);
	create index if not exists test_idx on foo(a);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("[ERR] cannot start transaction: %s", err)
	}
	stmt, err := db.Prepare("insert into foo(id, a, s) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < 10000; i++ {
		_, err = stmt.Exec(d.String(100), d.String(100), d.Int(13))
		if err != nil {
			log.Fatalf("failing insertion: %s", err)
		}
	}
	tx.Commit()

	b.ResetTimer()
	sl, err := db.Prepare("select * from foo where a=?")
	if err != nil {
		log.Fatalf("failing prepare query: %s", err)
	}
	defer sl.Close()

	for i := 0; i < b.N; i++ {
		var id string
		var a string
		var s int
		_ = sl.QueryRow(d.String(100)).Scan(&id, &a, &s)
	}
}

func BenchmarkSQLiteReadRangeWithoutIndex(b *testing.B) {
	os.Remove("./sqlite_bench.db")
	d := dummy.NewGenerator()

	db, err := sql.Open("sqlite3", "./sqlite_bench.db")
	if err != nil {
		log.Fatalf("cannot create db object: %s", err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id string not null primary key, a string, s integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("[ERR] cannot start transaction: %s", err)
	}
	stmt, err := db.Prepare("insert into foo(id, a, s) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < 10000; i++ {
		_, err = stmt.Exec(d.String(100), d.String(100), d.Int(13))
		if err != nil {
			log.Fatalf("failing insertion: %s", err)
		}
	}
	tx.Commit()

	b.ResetTimer()
	sl, err := db.Prepare("select * from foo where s > ?")
	if err != nil {
		log.Fatalf("failing prepare query: %s", err)
	}
	defer sl.Close()

	for i := 0; i < b.N; i++ {
		var id string
		var a string
		var s int
		_ = sl.QueryRow(d.Int(13)).Scan(&id, &a, &s)
	}
}

func BenchmarkSQLiteReadRangeWithIndex(b *testing.B) {
	os.Remove("./sqlite_bench.db")
	d := dummy.NewGenerator()

	db, err := sql.Open("sqlite3", "./sqlite_bench.db")
	if err != nil {
		log.Fatalf("cannot create db object: %s", err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id string not null primary key, a string, s integer);
	create index if not exists test_idx on foo(s);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("[ERR] cannot start transaction: %s", err)
	}
	stmt, err := db.Prepare("insert into foo(id, a, s) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < 10000; i++ {
		_, err = stmt.Exec(d.String(100), d.String(100), d.Int(13))
		if err != nil {
			log.Fatalf("failing insertion: %s", err)
		}
	}
	tx.Commit()

	b.ResetTimer()
	sl, err := db.Prepare("select * from foo where s > ?")
	if err != nil {
		log.Fatalf("failing prepare query: %s", err)
	}
	defer sl.Close()

	for i := 0; i < b.N; i++ {
		var id string
		var a string
		var s int
		_ = sl.QueryRow(d.Int(13)).Scan(&id, &a, &s)
	}
}
