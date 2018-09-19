/* main_test.go */

package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestUr(t *testing.T) {
	tests := []struct {
		req    events.APIGatewayProxyRequest
		expect string
		err    error
	}{
		{
			req:    events.APIGatewayProxyRequest{Body: `{"name": "Test"}`},
			expect: "The name is Test",
			err:    nil,
		},
	}

	for _, test := range tests {
		resp, err := Ur(test.req)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, resp.Body)
	}
}
