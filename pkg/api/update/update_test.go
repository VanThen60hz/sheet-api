package update

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

func TestUpdateDataRow(t *testing.T) {
	// test for correct respond
	requestBody := []byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "Sheet1",
		"rows": [
			[
				"3",
				"test3",
				"test3@gmail.com"
			],
			[
				"4",
				"test4",
				"test4@gmail.com"
			]
		],
		"range": [
			"4",
			"5"
		]
	}`)
	req := httptest.NewRequest(http.MethodPost, "/UpdateDataRow", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	UpdateDataRow(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	expectedResponse := "Update successfully!"
	actualResponse, _ := io.ReadAll(res.Body)
	if string(actualResponse) != expectedResponse {
		t.Errorf("Expected response body %q but got %q", expectedResponse, string(actualResponse))
	}

	// test for error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodPost, "/UpdateDataRow", &errorReader{})

	UpdateDataRow(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataRow", bytes.NewReader([]byte(``)))
	UpdateDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataRow", bytes.NewReader([]byte(`{}`)))
	UpdateDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataRow", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"}`)))
	UpdateDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataRow", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", "sheetName": "Sheet1"}`)))
	UpdateDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataRow", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", 
		"sheetName": "Sheet1",
		"rows": [
			[
				"3",
				"test3",
				"test3@gmail.com"
			],
			[
				"4",
				"test4",
				"test4@gmail.com"
			]
		]}`)))
	UpdateDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataRow", bytes.NewReader([]byte(`{
		"spreadsheetID": "error", 
		"sheetName": "Sheet1",
		"rows": [
			[
				"3",
				"test3",
				"test3@gmail.com"
			],
			[
				"4",
				"test4",
				"test4@gmail.com"
			]
		],
		"range": [
			"4",
			"5"
		]}`)))
	UpdateDataRow(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}
}

func TestUpdateDataCell(t *testing.T) {
	// test for correct respond
	requestBody := []byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetName": "Sheet1",
		"cells": [
			"test5",
			"test5@gmail.com"
		],
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
	req := httptest.NewRequest(http.MethodPost, "/UpdateDataCell", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	UpdateDataCell(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	expectedResponse := "Update successfully!"
	actualResponse, _ := io.ReadAll(res.Body)
	if string(actualResponse) != expectedResponse {
		t.Errorf("Expected response body %q but got %q", expectedResponse, string(actualResponse))
	}

	// test for error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodPost, "/UpdateDataCell", &errorReader{})

	UpdateDataCell(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataCell", bytes.NewReader([]byte(``)))
	UpdateDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataCell", bytes.NewReader([]byte(`{}`)))
	UpdateDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataCell", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"}`)))
	UpdateDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataCell", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", "sheetName": "Sheet1"}`)))
	UpdateDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataCell", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", 
		"sheetName": "Sheet1",
		"cells": [
			"test5",
			"test5@gmail.com"
		]}`)))
	UpdateDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateDataCell", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w", 
		"sheetName": "Sheet1",
		"cells": [
			"test5",
			"test5@gmail.com"
		],
		"range": [
			[
				"4",
				"-1"
			],
			[
				"5",
				"2"
			]
		]}`)))
	UpdateDataCell(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}
}

func TestUpdateSpreadsheet(t *testing.T) {
	// Mock request body
	reqBody := struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Title         string `json:"title"`
	}{
		SpreadsheetID: "test-spreadsheet-id",
		Title:         "Updated Spreadsheet Title",
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("PUT", "/UpdateSheet", bytes.NewBuffer(bodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateSpreadsheet)

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
		Message       string `json:"message"`
	}

	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if response.SpreadsheetID != "test-spreadsheet-id" {
		t.Errorf("Expected spreadsheetID %v, got %v", "test-spreadsheet-id", response.SpreadsheetID)
	}

	if response.Title != "Updated Spreadsheet Title" {
		t.Errorf("Expected title %v, got %v", "Updated Spreadsheet Title", response.Title)
	}

	if response.Message != "Spreadsheet title updated successfully" {
		t.Errorf("Expected message %v, got %v", "Spreadsheet title updated successfully", response.Message)
	}
}

func TestUpdateSheet(t *testing.T) {
	// test for correct respond
	requestBody := []byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetID": 123456,
		"newSheetName": "UpdatedSheet"
	}`)
	req := httptest.NewRequest(http.MethodPut, "/UpdateSheet", bytes.NewReader(requestBody))
	res := httptest.NewRecorder()

	UpdateSheet(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.Code)
	}

	var response struct {
		SpreadsheetID string `json:"spreadsheetID"`
		SheetID       int64  `json:"sheetID"`
		NewSheetName  string `json:"newSheetName"`
		Message       string `json:"message"`
	}

	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	if response.Message != "Sheet updated successfully" {
		t.Errorf("Expected message %q but got %q", "Sheet updated successfully", response.Message)
	}

	// test for error handling
	res_err := httptest.NewRecorder()
	req_err := httptest.NewRequest(http.MethodPut, "/UpdateSheet", &errorReader{})

	UpdateSheet(res_err, req_err)

	if res_err.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPut, "/UpdateSheet", bytes.NewReader([]byte(``)))
	UpdateSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPut, "/UpdateSheet", bytes.NewReader([]byte(`{}`)))
	UpdateSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPut, "/UpdateSheet", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w"}`)))
	UpdateSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPut, "/UpdateSheet", bytes.NewReader([]byte(`{
		"spreadsheetID": "13e2IAKNmZuj1asSc8qDKWhGC1OYRDxOEpUxEDPQpw8w",
		"sheetID": 123456}`)))
	UpdateSheet(res_err, req_err)

	if res_err.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res_err.Code)
	}

	// Test wrong method
	res_err = httptest.NewRecorder()
	req_err = httptest.NewRequest(http.MethodPost, "/UpdateSheet", bytes.NewReader(requestBody))
	UpdateSheet(res_err, req_err)

	if res_err.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d but got %d", http.StatusMethodNotAllowed, res_err.Code)
	}
}
