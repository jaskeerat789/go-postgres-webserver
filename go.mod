module github.com/jaskeerat789/go-postgres-webserver

go 1.16

require (
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-hclog v0.16.0
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/jackc/pgx/v4 v4.11.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/text v0.3.6 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.9
)

replace github.com/jaskeerat789/go-postgres-webserver => ./
