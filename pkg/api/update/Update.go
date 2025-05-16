package update

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"personnel-api/pkg/api/read"
	"personnel-api/pkg/svc"

	"google.golang.org/api/sheets/v4"
)

/*
PUT
Body: {
		"spreadsheetID": "YOUR_SPREAD_SHEET_ID",
		"sheetName": "SHEET_NAME",
		"rows":[ ["3", "test1", "test1@gmail.com"], ["4", "test2", "test2@gmail.com"]],
		"range": [row1, row2, ...]
	  }
*/
// check for valid length of input not included (rows and range has to match length)
func UpdateDataRow(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req struct {
		SpreadsheetID string          `json:"spreadsheetID"`
		SheetName     string          `json:"sheetName"`
		Rows          [][]interface{} `json:"rows"`
		Range         []interface{}   `json:"range"`
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	spreadsheetID := req.SpreadsheetID
	if spreadsheetID == "" {
		http.Error(w, "spreadsheetID field is required", http.StatusBadRequest)
		return
	}

	sheetName := req.SheetName
	if sheetName == "" {
		http.Error(w, "sheetName field is required", http.StatusBadRequest)
		return
	}

	rows := req.Rows
	if len(rows) == 0 {
		http.Error(w, "rows data field is required", http.StatusBadRequest)
		return
	}

	dataRange := req.Range
	if len(dataRange) == 0 {
		http.Error(w, "range field is required", http.StatusBadRequest)
		return
	}

	err = UpdateDataRowHelper(spreadsheetID, sheetName, dataRange, rows)
	if err != nil {
		http.Error(w, "Cannot update the rows requested", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Update successfully!")
}

func UpdateDataRowHelper(spreadsheetID string, sheetName string, dataRange []interface{}, rows [][]interface{}) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	columnRange, _, err := read.GetSheetDataHelper(spreadsheetID, sheetName)
	if err != nil {
		return err
	}
	arr := strings.Split(columnRange, ":")

	for i, row := range rows {
		valueRange := &sheets.ValueRange{
			Values: [][]interface{}{row},
		}

		rowNum := dataRange[i].(string)
		rowRange := sheetName + "!" + arr[0] + rowNum + ":" + arr[1] + rowNum

		_, err := service.Spreadsheets.Values.Update(spreadsheetID, rowRange, valueRange).ValueInputOption("USER_ENTERED").Do()
		if err != nil {
			return err
		}
	}

	return nil
}

/*
PUT

	Body: {
			"spreadsheetID": "YOUR_SPREAD_SHEET_ID",
			"sheetName": "SHEET_NAME",
			"cells":["test5", "test5@gmail.com"],
			"range": [["4", "1"], ["5", "2"]]
		  }
*/
//No type check yet
func UpdateDataCell(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req struct {
		SpreadsheetID string          `json:"spreadsheetID"`
		SheetName     string          `json:"sheetName"`
		Cells         []interface{}   `json:"cells"`
		Range         [][]interface{} `json:"range"`
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	spreadsheetID := req.SpreadsheetID
	if spreadsheetID == "" {
		http.Error(w, "spreadsheetID field is required", http.StatusBadRequest)
		return
	}

	sheetName := req.SheetName
	if sheetName == "" {
		http.Error(w, "sheetName field is required", http.StatusBadRequest)
		return
	}

	cells := req.Cells
	if len(cells) == 0 {
		http.Error(w, "cells data field is required", http.StatusBadRequest)
		return
	}

	dataRange := req.Range
	if len(dataRange) == 0 {
		http.Error(w, "range field is required", http.StatusBadRequest)
		return
	}

	err = UpdateDataCellHelper(spreadsheetID, sheetName, cells, dataRange)
	if err != nil {
		http.Error(w, "Cannot update the cells requested", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Update successfully!")
}

func UpdateDataCellHelper(spreadsheetID string, sheetName string, cells []interface{}, dataRange [][]interface{}) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	for i, pos := range dataRange {
		row := pos[0].(string)
		col_int, err := strconv.Atoi(pos[1].(string))
		if err != nil {
			return err
		}
		col := read.ColumnIndexToLetter(col_int)
		rowRange := sheetName + "!" + col + row + ":" + col + row

		cell_data := cells[i]
		valueRange := &sheets.ValueRange{
			Values: [][]interface{}{{cell_data}},
		}

		_, err = service.Spreadsheets.Values.Update(spreadsheetID, rowRange, valueRange).ValueInputOption("USER_ENTERED").Do()
		if err != nil {
			return err
		}
	}

	return nil
}

/*
PUT
Body: {"spreadsheetID": "your-spreadsheet-id", "title": "New Title"}
*/
func UpdateSpreadsheet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Title         string `json:"title"`
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	if req.SpreadsheetID == "" || req.Title == "" {
		http.Error(w, "Missing spreadsheetID or title", http.StatusBadRequest)
		return
	}

	err = UpdateSpreadsheetHelper(req.SpreadsheetID, req.Title)
	if err != nil {
		http.Error(w, "Failed to update spreadsheet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Title         string `json:"title"`
		Message       string `json:"message"`
	}{
		SpreadsheetID: req.SpreadsheetID,
		Title:         req.Title,
		Message:       "Spreadsheet title updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateSpreadsheetHelper(spreadsheetID, title string) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	requests := []*sheets.Request{
		{
			UpdateSpreadsheetProperties: &sheets.UpdateSpreadsheetPropertiesRequest{
				Properties: &sheets.SpreadsheetProperties{
					Title: title,
				},
				Fields: "title",
			},
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	}

	_, err = service.Spreadsheets.BatchUpdate(spreadsheetID, batchUpdateRequest).Do()
	return err
}

/*
PUT
Body: {"spreadsheetID": "YOUR_SPREAD_SHEET_ID", "sheetID": 123456, "newSheetName": "NEW_SHEET_NAME"}
*/
func UpdateSheet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req struct {
		SpreadsheetID string `json:"spreadsheetID"`
		SheetID       int64  `json:"sheetID"`
		NewSheetName  string `json:"newSheetName"`
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	spreadsheetID := req.SpreadsheetID
	if spreadsheetID == "" {
		http.Error(w, "spreadsheetID field is required", http.StatusBadRequest)
		return
	}

	if req.SheetID == 0 {
		http.Error(w, "sheetID field is required", http.StatusBadRequest)
		return
	}

	if req.NewSheetName == "" {
		http.Error(w, "newSheetName field is required", http.StatusBadRequest)
		return
	}

	err = UpdateSheetHelper(spreadsheetID, req.SheetID, req.NewSheetName)
	if err != nil {
		http.Error(w, "Cannot update sheet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		SpreadsheetID string `json:"spreadsheetID"`
		SheetID       int64  `json:"sheetID"`
		NewSheetName  string `json:"newSheetName"`
		Message       string `json:"message"`
	}{
		SpreadsheetID: spreadsheetID,
		SheetID:       req.SheetID,
		NewSheetName:  req.NewSheetName,
		Message:       "Sheet updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateSheetHelper(spreadsheetID string, sheetID int64, newSheetName string) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	req := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				UpdateSheetProperties: &sheets.UpdateSheetPropertiesRequest{
					Properties: &sheets.SheetProperties{
						SheetId: sheetID,
						Title:   newSheetName,
					},
					Fields: "title",
				},
			},
		},
	}

	_, err = service.Spreadsheets.BatchUpdate(spreadsheetID, req).Do()
	if err != nil {
		return err
	}

	return nil
}
