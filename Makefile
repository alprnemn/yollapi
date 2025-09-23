
# this command will compile project and create binary file named bin/social
build:
	@go build -o ./bin/yollapi ./cmd/api

# if compile operation is success, this command(run) runs the binary file on terminal
run: build
	@./bin/yollapi