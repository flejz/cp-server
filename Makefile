dev:
	reflex -r "\.go" -s -- bash -c "PORT=2000 SALT=dev go run init/service.go"
