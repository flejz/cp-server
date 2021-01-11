dev:
	reflex -r "\.go" -s -- sh -c "SALT=dev go run init/service.go"
