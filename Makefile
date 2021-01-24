dev: dev-sqlite

dev-mem:
	reflex -r "\.go" -s -- bash -c "DB_TYPE=mem MEM_NAME=test PORT=2000 SALT=dev go run cmd/cp-tcp-server.go"

dev-sqlite:
	reflex -r "\.go" -s -- bash -c "DB_TYPE=sqlite SQLITE_PATH=cp-server.db PORT=2000 SALT=dev go run cmd/cp-tcp-server.go"

dev-cli-get:
	go run cmd/cp-cli.go get -u vai -p vai

dev-cli-set:
	go run cmd/cp-cli.go set bleiba -u vai -p vai

debug: debug-mem

debug-mem:
	DB_TYPE=mem MEM_NAME=test PORT=2000 SALT=dev dlv debug cmd/cp-tcp-server.go

debug-sqlite:
	DB_TYPE=sqlite SQLITE_PATH=cp-server.db PORT=2000 SALT=dev dlv debug cmd/cp-tcp-server.go

dev-server:
	reflex -r "\.go" -s -- bash -c "SQLITE_PATH=~/cp-server.db PORT=2000 SALT=dev go run init/server.go"
