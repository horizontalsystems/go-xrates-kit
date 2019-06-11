package handler

import (
	"fmt"
	"testing"
	"time"
)

func TestReformatIPFSResponse(t *testing.T) {
	dCcy := "BTC"
	fCcy := "USD"

	epochSec := int64(1559214146)
	timeObj := time.Unix(epochSec, 0).UTC() // May 30, 2019 11:02 AM
	//------------------------------------------------------------------
	expectedResponse1 := "{\"BTC\":{\"currency\":\"USD\",\"rate\":\"12.0\",\"time\":1559214146}}"
	jsonData1 := "{\"00\" : \"10.0\",\"01\" : \"11.0\",\"02\" : \"12.0\"}"

	t.Run("1 Test", testReformatIPFSResponseFunc(expectedResponse1, jsonData1, dCcy, fCcy, &timeObj))

	//------------------------------------------------------------------
	expectedResponse2 := "{\"BTC\":{\"currency\":\"USD\",\"rate\":\"12.0\",\"time\":1559214146}}"
	jsonData2 := "\"12.0\""

	t.Run("2 Test", testReformatIPFSResponseFunc(expectedResponse2, jsonData2, dCcy, fCcy, &timeObj))
}

func testReformatIPFSResponseFunc(expectedResponse string, jsonData string,
	dCcy string, fCcy string, timeObj *time.Time) func(*testing.T) {
	return func(t *testing.T) {

		actual := reformatIPFSResponse(jsonData, dCcy, fCcy, timeObj)

		if actual != expectedResponse {
			t.Error(fmt.Sprintf("Expected Result is %v , but got %v", expectedResponse, actual))
		}
	}
}
