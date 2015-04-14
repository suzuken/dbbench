# benchmark

spec

* MacBook Pro, Late 2013
* 2.6 GHz Intel Core i7
* 16GB 1600 MHz DDR3
* sqlite3: 3.7.13
* bolt: based on https://github.com/boltdb/bolt/commit/3b449559cf34cbcc74460b59041a4399d3226e5a

# 2015/04/14

	-> % sw_vers
	ProductName:    Mac OS X
	ProductVersion: 10.9.5
	BuildVersion:   13F34

	-> % make bench
	go test -bench .
	testing: warning: no tests to run
	PASS
	BenchmarkBoltInsert       100000             93363 ns/op
	BenchmarkBoltBatchInsert          100000             91444 ns/op
	BenchmarkBoltInsertManualTx       100000             91008 ns/op
	BenchmarkBoltCoalescerInsert        3000            485062 ns/op
	BenchmarkSQLiteReplaceInsertWithTx         50000             37704 ns/op
	BenchmarkSQLiteReplaceInsert        2000            760006 ns/op
	BenchmarkSQLiteInsertWithTx       100000             29173 ns/op
	BenchmarkSQLiteInsert       2000            750359 ns/op
	BenchmarkSQLiteReadPrimaryKey     100000             22894 ns/op
	BenchmarkSQLiteReadKeyWithoutIndex           300           6194131 ns/op
	BenchmarkSQLiteReadKeyWithIndex   100000             23651 ns/op
	BenchmarkSQLiteReadRangeWithoutIndex       50000             26279 ns/op
	BenchmarkSQLiteReadRangeWithIndex          50000             26614 ns/op
	ok      github.com/suzuken/dbbench      190.210s
