test:
	go test -v -cover -coverprofile=coverage.out ./...

cover:
	go tool cover -html=coverage.out

bench: bench_singly bench_doubly bench_vector

bench_singly:
	go test -gcflags="-l -N" -v ./collections/sequence_containers/singly/... -bench . -run ^$$

bench_doubly:
	go test -gcflags="-l -N" -v ./collections/sequence_containers/doubly/... -bench . -run ^$$

bench_vector:
	go test -v ./collections/sequence_containers/vector/... -bench . -run ^$$
