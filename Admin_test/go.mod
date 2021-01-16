module awesomeProject3

go 1.14

require (
	github.com/go-sql-driver/mysql v1.4.0
	github.com/gorilla/mux v1.7.4
	github.com/jmoiron/sqlx v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.15.0
	google.golang.org/appengine v1.6.5 // indirect
)

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
