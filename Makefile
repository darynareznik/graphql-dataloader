.PHONY: gen
gen:
	go get -tool github.com/99designs/gqlgen
	go tool gqlgen generate

.PHONY: deps
deps:
	go mod tidy