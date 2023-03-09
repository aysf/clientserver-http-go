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