dev: dev-mem

dev-mem:
	reflex -r "\.go" -s -- bash -c "DB_TYPE=mem MEM_NAME=test PORT=2000 SALT=dev go run cmd/tcpserver/main.go"

dev-sqlite:
	reflex -r "\.go" -s -- bash -c "DB_TYPE=sqlite PORT=2000 SALT=dev go run cmd/tcpserver/main.go"

dev-cli-get:
	go run cmd/cp-cli.go get -u test -p test

dev-cli-set:
	go run cmd/cp-cli.go set bleiba -u test -p test

debug: debug-mem

debug-mem:
	DB_TYPE=mem MEM_NAME=test PORT=2000 SALT=dev dlv debug cmd/cp-tcp-server.go

debug-sqlite:
	DB_TYPE=sqlite SQLITE_PATH=cp-server.db PORT=2000 SALT=dev dlv debug cmd/cp-tcp-server.go
