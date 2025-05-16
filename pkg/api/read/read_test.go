package read

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("simulated error while reading request body")
}

func TestGetAll(t *testing.T) {
	// test read ok
	requestBody := []byte(`{"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"}`)
	req := httptest.NewRequest(http.MethodGet, "/GetAll", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	GetAll(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	// test error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodGet, "/GetAll", &errorReader{})

	GetAll(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetAll", bytes.NewReader([]byte(``)))

	GetAll(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetAll", bytes.NewReader([]byte(`{}`)))

	GetAll(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetAll", bytes.NewReader([]byte(`{
		"spreadsheetID": "error"
		}`)))

	GetAll(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}
}

func TestGetSheetData(t *testing.T) {
	// test read ok
	requestBody := []byte(`{"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", "sheetName": "Sheet1"}`)
	req := httptest.NewRequest(http.MethodGet, "/GetSheetData", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	GetSheetData(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	// test error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodGet, "/GetSheetData", &errorReader{})

	GetSheetData(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetSheetData", bytes.NewReader([]byte(``)))

	GetSheetData(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetSheetData", bytes.NewReader([]byte(`{}`)))

	GetSheetData(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetSheetData", bytes.NewReader([]byte(`{"spreadsheetID": "error"}`)))

	GetSheetData(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}
}

func TestGetByColumn(t *testing.T) {
	// test read ok
	requestBody := []byte(`{"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
	"sheetName": "Sheet1",
	"columnName": "Email"}`)
	req := httptest.NewRequest(http.MethodGet, "/GetByColumn", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	GetByColumn(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	// test error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodGet, "/GetByColumn", &errorReader{})

	GetByColumn(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByColumn", bytes.NewReader([]byte(``)))

	GetByColumn(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByColumn", bytes.NewReader([]byte(`{}`)))

	GetByColumn(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByColumn", bytes.NewReader([]byte(`{"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"}`)))

	GetByColumn(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByColumn", bytes.NewReader([]byte(`{"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", "sheetName": "Sheet1"}`)))

	GetByColumn(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByColumn", bytes.NewReader([]byte(`{"spreadsheetID": "Error", "sheetName": "Sheet1", "columnName": "Email"}`)))

	GetByColumn(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}
}

func TestGetByFilter(t *testing.T) {
	// test read ok
	requestBody1 := []byte(`{    "spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
    "sheetName": "Sheet1",
    "columnName": "ID",
    "operator": ">",
    "value": "1"}`)

	req := httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader(requestBody1))
	res := httptest.NewRecorder()

	GetByFilter(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	requestBody2 := []byte(`{    "spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
    "sheetName": "Sheet1",
    "columnName": "ID",
    "operator": "=",
    "value": "1"}`)

	req = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader(requestBody2))
	res = httptest.NewRecorder()

	GetByFilter(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	requestBody3 := []byte(`{    "spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
    "sheetName": "Sheet1",
    "columnName": "ID",
    "operator": "<",
    "value": "3"}`)

	req = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader(requestBody3))
	res = httptest.NewRecorder()

	GetByFilter(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	requestBody4 := []byte(`{    
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "Sheet1",
		"columnName": "Email",
		"operator": "contain",
		"value": ".com"
	}`)

	req = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader(requestBody4))
	res = httptest.NewRecorder()

	GetByFilter(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	requestBody5 := []byte(`{    
	"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
    "sheetName": "Sheet2",
    "columnName": "Score",
    "operator": "<",
    "value": "10.0"}`)

	req = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader(requestBody5))
	res = httptest.NewRecorder()

	GetByFilter(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	requestBody6 := []byte(`{    "spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
    "sheetName": "Sheet2",
    "columnName": "Score",
    "operator": "=",
    "value": "9.8"}`)

	req = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader(requestBody6))
	res = httptest.NewRecorder()

	GetByFilter(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	requestBody7 := []byte(`{    "spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
    "sheetName": "Sheet2",
    "columnName": "Score",
    "operator": ">",
    "value": "9.5"}`)

	req = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader(requestBody7))
	res = httptest.NewRecorder()

	GetByFilter(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	// test error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodGet, "/GetByFilter", &errorReader{})

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(``)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(`{}`)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"}`)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", 
		"sheetName": "Sheet1"}`)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", 
		"sheetName": "Sheet1", 
		"columnName": "ID"}`)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", 
		"sheetName": "Sheet1", 
		"columnName": "ID", 
		"operator": ">"}`)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(`{
		"spreadsheetID": "error", 
		"sheetName": "Sheet1", 
		"columnName": "ID", 
		"operator": ">", 
		"value": "1"}`)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", 
		"sheetName": "Sheet1", 
		"columnName": "ID", 
		"operator": "+", 
		"value": "1"}`)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", 
		"sheetName": "Sheet1", 
		"columnName": "Email", 
		"operator": "+", 
		"value": ".com"}`)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetByFilter", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", 
		"sheetName": "Sheet2", 
		"columnName": "Score", 
		"operator": "+", 
		"value": "9.8"}`)))

	GetByFilter(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}
}

func TestGetSheets(t *testing.T) {
	// test for correct respond
	req := httptest.NewRequest(http.MethodGet, "/GetSheets?spreadsheetID=13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", nil)
	res := httptest.NewRecorder()

	GetSheets(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	var response struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Sheets        []struct {
			SheetID    int64  `json:"sheetID"`
			SheetName  string `json:"sheetName"`
			SheetIndex int64  `json:"sheetIndex"`
		} `json:"sheets"`
	}

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if response.SpreadsheetID == "" {
		t.Error("Expected non-empty spreadsheetID")
	}

	if len(response.Sheets) == 0 {
		t.Error("Expected at least one sheet in response")
	}

	// test for error handling - missing spreadsheetID
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodGet, "/GetSheets", nil)
	GetSheets(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	// test for error handling - invalid spreadsheetID
	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodGet, "/GetSheets?spreadsheetID=invalid", nil)
	GetSheets(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	// Test wrong method
	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/GetSheets?spreadsheetID=13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", nil)
	GetSheets(res_err, req_err)

	if res_err.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d but got %d", http.StatusMethodNotAllowed, res_err.Code)
	}
}
