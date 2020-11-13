package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDloPriceRequest(t *testing.T) {
	req, err := createDloPriceRequest()
	assert.NoError(t, err)
	assert.Equal(t, "GET", req.Method)
	assert.Equal(t, stelExURL, req.URL.String())
}

func TestParseAiBlocksExpertResponse(t *testing.T) {
	body := "hello"
	gotPrice, gotErr := parseAiBlocksExpertLatestPrice(body)
	assert.EqualError(t, gotErr, "mis-formed response from aiblocks expert")

	body = "hello,"
	gotPrice, gotErr = parseAiBlocksExpertLatestPrice(body)
	assert.EqualError(t, gotErr, "mis-formed price from aiblocks expert")

	body = "[[10001,hello]"
	gotPrice, gotErr = parseAiBlocksExpertLatestPrice(body)
	assert.Error(t, gotErr)

	body = "[[100001,5.00],[100002,6.00]]"
	wantPrice := 5.00
	gotPrice, gotErr = parseAiBlocksExpertLatestPrice(body)
	assert.NoError(t, gotErr)
	assert.Equal(t, wantPrice, gotPrice)
}
