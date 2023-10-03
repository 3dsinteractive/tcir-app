#!/bin/bash

# test_api is the main function to run test
function test_api() {
    # If found go in default path, it will use go from default path
    GO=/usr/local/go/bin/go
    if [ -f "$GO" ]; then
        /usr/local/go/bin/go mod init automationworkshop/main
        /usr/local/go/bin/go mod tidy
        /usr/local/go/bin/go get
        /usr/local/go/bin/go mod vendor
        /usr/local/go/bin/go test . --run TestMainFailTestSuite
    else 
        go mod init automationworkshop/main
        go mod tidy
        go get
        go mod vendor
        go test . --run TestMainFailTestSuite
    fi
}

# Run main test process
test_api