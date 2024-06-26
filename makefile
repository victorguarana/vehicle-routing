test:
	ginkgo run -r

lint:
	golangci-lint run

count-lines:
	find . -name '*.go' | xargs wc -l