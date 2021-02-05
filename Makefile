test:
	go test -v -cover -coverprofile=coverage.out ./...

cover:
	go tool cover -html=coverage.out

bench: bench_singly bench_doubly bench_vector

bench_initial:
	go test -v  -bench . -run ^$$ ./collections/sequence_containers/initial/...

bench_singly:
	go test -v  -bench . -run ^$$ ./collections/sequence_containers/singly/...

bench_doubly:
	go test -v  -bench . -run ^$$ ./collections/sequence_containers/doubly/...

bench_vector:
	go test -v  -bench . -run ^$$ ./collections/sequence_containers/vector/...

lint:
	golangci-lint run
