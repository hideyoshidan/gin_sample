MIGRATIONDIR = 'database/migrations'
MYSQL_URL = 'mysql://go_user:go@tcp(go_db:3306)/go_sample  ?charset=utf8&parseTime=True&loc=Local'

#===============================================
#  build 関連
#=============================================== 
# go build
run:
	go run main.go

#===============================================
# migration関連
#=============================================== 
# migrationfileを作成
# 例) $ make migrate-create "SAMPLE=ファイル名"
migrate-create:
	migrate create -ext sql -dir ${MIGRATIONDIR} -seq ${SAMPLE}

# migration実行
# 例) $ make migrate-up
migrate-up:
	migrate -database ${MYSQL_URL} -path ${MIGRATIONDIR} up

# migration down
# 例) $ make migrate-down
migrate-down:
	migrate -database ${MYSQL_URL} -path ${MIGRATIONDIR} down

# migration force
# 例) $ make migrate-force "VERSION=1"
migrate-force:
	migrate -database ${MYSQL_URL} -path ${MIGRATIONDIR} force ${VERSION} 