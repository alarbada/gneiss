default:
	go build -o gneiss.bin ./cmd

run-example:
	go run ./examples/

build-cmd:
	go build -o gneiss.bin ./cmd/gneiss

watch-go:
	reflex -d none -s -r '.*\.go' -- make -s build-cmd
