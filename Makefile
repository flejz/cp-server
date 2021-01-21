dev:
	reflex -r "\.go" -s -- bash -c "PORT=2000 SALT=dev go run init/service.go"

dev-server:
	reflex -r "\.go" -s -- bash -c "SQLITE_PATH=~/cp-server.db PORT=2000 SALT=dev go run init/server.go"
