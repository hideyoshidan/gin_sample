# 1. 構築 & 動作確認
### ① Build
```
$ docker compose build --no-cache
$ docker compose up -d
```

### ② container　に入る
```
$ docker compose exec go_app bash 
root@0b27509921ce:/var/www/go_bbs#
```

### ③ migrateコマンドが使えるか確認
```
# migrate --help
Usage: migrate OPTIONS COMMAND [arg...]
       migrate [ -version | -help ]

Options:
  -source          Location of the migrations (driver://url)
  -path            Shorthand for -source=file://path
  -database        Run migrations against this database (driver://url)
  -prefetch N      Number of migrations to load in advance before executing (default 10)
  -lock-timeout N  Allow N seconds to acquire database lock (default 15)
  -verbose         Print verbose logging
  -version         Print version
  -help            Print usage

Commands:
  create [-ext E] [-dir D] [-seq] [-digits N] [-format] NAME
			   Create a set of timestamped up/down migrations titled NAME, in directory D with extension E.
			   Use -seq option to generate sequential up/down migrations with N digits.
			   Use -format option to specify a Go time format string. Note: migrations with the same time cause "duplicate migration version" error. 
  goto V       Migrate to version V
  up [N]       Apply all or N up migrations
  down [N]     Apply all or N down migrations
  drop         Drop everything inside database
  force V      Set version V but don't run migration (ignores dirty state)
  version      Print current migration version

Source drivers: github-ee, gitlab, go-bindata, godoc-vfs, gcs, file, s3, github
Database drivers: sqlserver, cockroachdb, firebirdsql, mysql, postgres, spanner, crdb-postgres, mongodb, redshift, stub, cassandra, clickhouse, cockroach, firebird, postgresql
```

### ④ go modules install
```
# go mod tidy
go: downloading github.com/stretchr/testify v1.4.0
go: downloading github.com/go-sql-driver/mysql v1.5.0
go: downloading github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5
go: downloading github.com/jinzhu/now v1.0.1
go: downloading golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd
go: downloading github.com/denisenkom/go-mssqldb v0.0.0-20191124224453-732737034ffd
go: downloading github.com/lib/pq v1.1.1
go: downloading github.com/mattn/go-sqlite3 v2.0.1+incompatible
go: downloading github.com/go-playground/assert/v2 v2.0.1
go: downloading gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe
```

### ⑤ migrationを実行(Makefileを使用)
```
# make migrate-up
migrate -database 'mysql://go_user:go@tcp(go_db:3306)/go_sample?charset=utf8&parseTime=True&loc=Local' -path 'database/migrations' up
1/u create_sample_table (48.039873ms)
```

### ⑥ terminalのタブをもう一つ開き、mysqlコンテナに入り、適当なデータを登録
```
$ docker-compose exec go_db bash -c "mysql go_sample -ugo_user -pgo"
mysql: [Warning] Using a password on the command line interface can be insecure.
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 6
Server version: 5.7.35 MySQL Community Server (GPL)

Copyright (c) 2000, 2021, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```

### ⑦ samples tableがあるか確認
```
mysql> show tables;
+---------------------+
| Tables_in_go_sample |
+---------------------+
| samples             |
| schema_migrations   |
+---------------------+
2 rows in set (0.00 sec)
```

### ⑧ samples tableに適当なデータをinsert
```
mysql> insert into samples (text) value('AAAAAAAAAAAAAAA');
Query OK, 1 row affected (0.00 sec)
```

### ⑨ goのコンテナに戻り、実行(Makefileを使用)
```
# make run
go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /sample/:id               --> github.com/hideyoshidan/go_bbs/infrastructure.(*Routing).setRouting.func1 (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```
### ⑩ ブラウザで「http://localhost:8080/sample/1」にアクセスし、以下の値が出ることを確認
```
{"message":"success","data":{"id":1,"text":"AAAAAAAAAAAAAAA"}}
```



# 2. 参考サイト
## Docker
    - 第一弾 : https://bit.ly/2UuBqNX	
    - 第二弾 : https://bit.ly/3kzsnpC

## Go
    - https://bit.ly/2UuSeEu	

## GIN
    - https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a
    - https://bsblog.casareal.co.jp/archives/4822
    - https://qiita.com/Asuforce/items/0bde8cabb30ac094fcb4

## GO module管理について
    - https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode

## GINを動かすまでの手順
```
$ go mod tidy
$ go run main
```

## GOPATHについて
    - https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/01.2.html

## Sliceについて
    - https://tamago-engineer.hateblo.jp/entry/2018/01/05/224402
    - https://bit.ly/3zkPxEj
    - https://code-graffiti.com/make-of-slice-in-golang/
    - https://www.techscore.com/tech/Go/Lang/Basic14/#2

## interfaceについて
    - https://qiita.com/rtok/items/46eadbf7b0b7a1b0eb08

## golang-migrate
    - https://dev.classmethod.jp/articles/db-migrate-with-golang-migrate/
