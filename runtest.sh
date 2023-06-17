DB_NAME=dbtest \
DB_HOST=hosttest \
DB_PORT=porttest \
DB_USER=usertest \
DB_PASSWORD=passtest \
DB_SSL_MODE=disable \
APP_PORT=4000 \
gotest $(go list ./... | grep -v /ent | grep -v /staticdata | grep -v /extra | grep -v /database) -coverprofile=cover.out
echo "#######TOTAL######"
echo ""
go tool cover -func=cover.out
echo ""
