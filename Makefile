dev: dev-sqlite

dev-mem:
	reflex -r "\.go" -s -- bash -c "DB_TYPE=mem MEM_NAME=test PORT=2000 SALT=dev go run cmd/tcp_server.go"

dev-sqlite:
	reflex -r "\.go" -s -- bash -c "DB_TYPE=sqlite SQLITE_PATH=cp-server.db PORT=2000 SALT=dev go run cmd/tcp_server.go"

debug: debug-mem

debug-mem:
	DB_TYPE=mem MEM_NAME=test PORT=2000 SALT=dev dlv debug cmd/tcp_server.go

dev-server:
	reflex -r "\.go" -s -- bash -c "SQLITE_PATH=~/cp-server.db PORT=2000 SALT=dev go run init/server.go"
