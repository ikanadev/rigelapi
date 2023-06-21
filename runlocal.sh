# set DB_SSL_MODE to  verify-full when connecting to neon db
DB_HOST=localhost DB_NAME=enttest DB_PORT=5432 DB_USER=taylor DB_PASSWORD=postgres DB_SSL_MODE=disable APP_PORT=4000 go run main.go
