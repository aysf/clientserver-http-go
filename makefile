server-mux:
	@echo "start server mux"
	go run cmd/serverWithMux/*.go

client-mux:
	@echo "start client mux"
	go run cmd/clientWithMux/*.go

server-mux-var:
	@echo "start client mux"
	go run cmd/serverWithMuxAndVar/*.go

client-mux-var:
	@echo "start client mux"
	go run cmd/clientWithMuxAndVar/*.go

server-mux-post:
	@echo "start client mux"
	go run cmd/serverWithMuxWithPost/*.go

client-mux-post:
	@echo "start client mux"
	go run cmd/clientWithMuxWithPost/*.go

server-secure:
	@echo "start client secure"
	go run cmd/serverSecure/*.go

client-secure:
	@echo "start client secure"
	go run cmd/clientSecure/*.go

client-insecure:
	@echo "start client secure"
	go run cmd/clientInsecure/*.go