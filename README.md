# 1. 構築手順
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
root@0b27509921ce:/var/www/go_bbs# migrate --help
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
