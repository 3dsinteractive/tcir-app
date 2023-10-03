test:
	go test . --run TestMainTestSuite
test-fail:
	go test . --run TestMainFailTestSuite
tidy:
	go mod tidy
build:
	go build .