package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("simulated error while reading request body")
}

func TestCreateData(t *testing.T) {
	// test for correct respond
	requestBody := []byte(`{"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", "sheetName": "Sheet1", "rows":[ ["3", "test1", "test1@gmail.com"], ["4", "test2", "test2@gmail.com"] ]}`)
	req := httptest.NewRequest(http.MethodPost, "/CreateData", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	CreateData(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	expectedResponse := "Insert successfully!"
	actualResponse, _ := io.ReadAll(res.Body)
	if string(actualResponse) != expectedResponse {
		t.Errorf("Expected response body %q but got %q", expectedResponse, string(actualResponse))
	}

	// test for error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodPost, "/CreateData", &errorReader{})

	CreateData(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/CreateData", bytes.NewReader([]byte(``)))
	CreateData(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/CreateData", bytes.NewReader([]byte(`{}`)))
	CreateData(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/CreateData", bytes.NewReader([]byte(`{"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"}`)))
	CreateData(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/CreateData", bytes.NewReader([]byte(`{"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", "sheetName": "Sheet1"}`)))
	CreateData(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/CreateData", bytes.NewReader([]byte(`{"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", "sheetName": "Sheet", "rows":[ ["3", "4", "test1", "test1@gmail.com"] ]}`)))
	CreateData(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}
}

func TestCreateSpreadsheet(t *testing.T) {
	// Mock request body
	reqBody := struct {
		Title string `json:"title"`
	}{
		Title: "Test Spreadsheet",
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "/CreateSheet", bytes.NewBuffer(bodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateSpreadsheet)

	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Title         string `json:"title"`
	}

	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if response.SpreadsheetID == "" {
		t.Errorf("Expected non-empty spreadsheetID")
	}

	if response.Title != "Test Spreadsheet" {
		t.Errorf("Expected title %v, got %v", "Test Spreadsheet", response.Title)
	}
}

func TestCreateSheet(t *testing.T) {
	// test for correct respond
	requestBody := []byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "NewTestSheet"
	}`)
	req := httptest.NewRequest(http.MethodPost, "/CreateSheet", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	CreateSheet(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	var response struct {
		SpreadsheetID string `json:"spreadsheetID"`
		SheetName     string `json:"sheetName"`
		Message       string `json:"message"`
	}

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if response.Message != "Sheet created successfully" {
		t.Errorf("Expected message %q but got %q", "Sheet created successfully", response.Message)
	}

	// test for error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodPost, "/CreateSheet", &errorReader{})

	CreateSheet(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/CreateSheet", bytes.NewReader([]byte(``)))
	CreateSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/CreateSheet", bytes.NewReader([]byte(`{}`)))
	CreateSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/CreateSheet", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"}`)))
	CreateSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/CreateSheet", bytes.NewReader([]byte(`{
		"sheetName": "NewTestSheet"}`)))
	CreateSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	// Test wrong method
	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPut, "/CreateSheet", bytes.NewReader(requestBody))
	CreateSheet(res_err, req_err)

	if res_err.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d but got %d", http.StatusMethodNotAllowed, res_err.Code)
	}
}
