package main

import (
	"net/http"
	"testing"

	"github.com/3dsinteractive/testify/suite"
	"github.com/stretchr/testify/assert"
)

type MainFailTestSuite struct {
	suite.Suite
}

func TestMainFailTestSuite(t *testing.T) {
	suite.Run(t, new(MainFailTestSuite))
}

func (ts *MainFailTestSuite) TestOnPostClient() {
	is := assert.New(ts.T())

	var ctx *ContextMock
	var cfg *ConfigMock
	var rnd *RandomMock
	var err error
	var status map[string]interface{}

	// Case: random start with 0
	ctx = NewContextMock()
	status = map[string]interface{}{
		"status": "success",
	}
	ctx.On("Response", http.StatusOK, status)
	cfg = NewConfigMock()
	rnd = NewRandomMock()
	rnd.On("Random").Return("0000") // start with 0
	err = onPostClient(ctx, cfg, rnd)
	is.NoError(err)
	ctx.AssertExpectations(ts.T())

}
