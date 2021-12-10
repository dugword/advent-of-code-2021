.PHONY: \
	all \
	clean \
	test \
	view-coverage

all:
	for dir in cmd/*; do \
		outfile=$$(echo $$dir | perl -pe 's<^cmd/><>'); \
		go run "./$$dir" -input ./$$dir/input > ./solutions/$${outfile}_part_1; \
		go run "./$$dir" -part-2 -input ./$$dir/input > ./solutions/$${outfile}_part_2; \
	done

clean:
	rm solutions/*

test:
	go test -coverprofile=coverage.out ./cmd/...

view-coverage: test
	go tool cover -html=coverage.out
