package delete

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testSpreadsheetID = "1Y6NRaduDsw_Wxu0yEhomYWXeCjBteFlnovj7TiPAyM8"

type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("simulated error while reading request body")
}

func TestDeleteDataRow(t *testing.T) {
	// test for correct respond
	requestBody := []byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "Sheet1",
		"range": [
			"4",
			"5"
		]
	}`)
	req := httptest.NewRequest(http.MethodPost, "/DeleteDataRow", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	DeleteDataRow(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	expectedResponse := "Delete successfully!"
	actualResponse, _ := io.ReadAll(res.Body)
	if string(actualResponse) != expectedResponse {
		t.Errorf("Expected response body %q but got %q", expectedResponse, string(actualResponse))
	}

	// test for error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodPost, "/DeleteDataRow", &errorReader{})

	DeleteDataRow(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataRow", bytes.NewReader([]byte(``)))
	DeleteDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataRow", bytes.NewReader([]byte(`{}`)))
	DeleteDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataRow", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"
	}`)))
	DeleteDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataRow", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "Sheet1"
	}`)))
	DeleteDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataRow", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "Sheet1",
		"range": [
			"-4",
			"5"
		]
	}`)))
	DeleteDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}
}

func TestDeleteDataCell(t *testing.T) {
	// test for correct respond
	requestBody := []byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "Sheet1",
		"range": [
			[
				"4",
				"1"
			],
			[
				"5",
				"2"
			]
		]
	}`)
	req := httptest.NewRequest(http.MethodPost, "/DeleteDataCell", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	DeleteDataCell(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	expectedResponse := "Delete successfully!"
	actualResponse, _ := io.ReadAll(res.Body)
	if string(actualResponse) != expectedResponse {
		t.Errorf("Expected response body %q but got %q", expectedResponse, string(actualResponse))
	}

	// test for error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodPost, "/DeleteDataCell", &errorReader{})

	DeleteDataCell(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataCell", bytes.NewReader([]byte(``)))
	DeleteDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataCell", bytes.NewReader([]byte(`{}`)))
	DeleteDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataCell", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"
	}`)))
	DeleteDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataCell", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "Sheet1"
	}`)))
	DeleteDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteDataCell", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "Sheet1",
		"range": [
			[
				"4",
				"-1"
			],
			[
				"5",
				"2"
			]
		]
	}`)))
	DeleteDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}
}

func TestDeleteSpreadsheet(t *testing.T) {
	// Skip test in short mode (CI/CD environment)
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	// test for correct respond
	requestBody := []byte(fmt.Sprintf(`{
		"spreadsheetID": "%s"
	}`, testSpreadsheetID))
	req := httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	DeleteSpreadsheet(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	// Parse the response
	var response struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Message       string `json:"message"`
	}
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Errorf("Failed to parse response body: %v", err)
	}

	// Check response fields
	expectedResponse := struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Message       string `json:"message"`
	}{
		SpreadsheetID: testSpreadsheetID,
		Message:       "Spreadsheet deleted successfully",
	}

	if response != expectedResponse {
		t.Errorf("Response mismatch\nExpected: %+v\nGot: %+v", expectedResponse, response)
	}

	// test for error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodDelete, "/DeleteSheet", &errorReader{})

	DeleteSpreadsheet(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader([]byte(``)))
	DeleteSpreadsheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader([]byte(`{}`)))
	DeleteSpreadsheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader([]byte(`{
		"spreadsheetID": ""
	}`)))
	DeleteSpreadsheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader([]byte(`{invalid json}`)))
	DeleteSpreadsheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}
}

func TestDeleteSheet(t *testing.T) {
	// test for correct respond
	requestBody := []byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetID": 123456
	}`)
	req := httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	DeleteSheet(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	var response struct {
		SpreadsheetID string `json:"spreadsheetID"`
		SheetID       int64  `json:"sheetID"`
		Message       string `json:"message"`
	}

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if response.Message != "Sheet deleted successfully" {
		t.Errorf("Expected message %q but got %q", "Sheet deleted successfully", response.Message)
	}

	// test for error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodDelete, "/DeleteSheet", &errorReader{})

	DeleteSheet(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader([]byte(``)))
	DeleteSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader([]byte(`{}`)))
	DeleteSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"}`)))
	DeleteSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodDelete, "/DeleteSheet", bytes.NewReader([]byte(`{
		"sheetID": 123456}`)))
	DeleteSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	// Test wrong method
	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/DeleteSheet", bytes.NewReader(requestBody))
	DeleteSheet(res_err, req_err)

	if res_err.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d but got %d", http.StatusMethodNotAllowed, res_err.Code)
	}
}
