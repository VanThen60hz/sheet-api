package delete

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
DELETE
Body: {
		"spreadsheetID": "YOUR_SPREAD_SHEET_ID",
		"sheetName": "SHEET_NAME",
		"range": [3, 4]
	  }
*/

func DeleteDataRow(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req struct {
		SpreadsheetID string        `json:"spreadsheetID"`
		SheetName     string        `json:"sheetName"`
		Range         []interface{} `json:"range"`
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

	dataRange := req.Range
	if len(dataRange) == 0 {
		http.Error(w, "range field is required", http.StatusBadRequest)
		return
	}

	err = DeleteDataRowHelper(spreadsheetID, sheetName, dataRange)
	if err != nil {
		http.Error(w, "Cannot delete the rows requested", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Delete successfully!")
}

func DeleteDataRowHelper(spreadsheetID string, sheetName string, dataRange []interface{}) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	columnRange, _, _ := read.GetSheetDataHelper(spreadsheetID, sheetName)
	arr := strings.Split(columnRange, ":")
	request := &sheets.ClearValuesRequest{}

	for i := range dataRange {
		rowNum := dataRange[i].(string)
		rowRange := sheetName + "!" + arr[0] + rowNum + ":" + arr[1] + rowNum

		_, err := service.Spreadsheets.Values.Clear(spreadsheetID, rowRange, request).Do()
		if err != nil {
			return err
		}
	}

	return nil
}

/*
DELETE

	Body: {
				"spreadsheetID": "YOUR_SPREAD_SHEET_ID",
				"sheetName": "SHEET_NAME",
				"range": [["4", "1"], ["5", "2"]]
			  }
*/
func DeleteDataCell(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req struct {
		SpreadsheetID string          `json:"spreadsheetID"`
		SheetName     string          `json:"sheetName"`
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

	dataRange := req.Range
	if len(dataRange) == 0 {
		http.Error(w, "range field is required", http.StatusBadRequest)
		return
	}

	err = DeleteDataCellHelper(spreadsheetID, sheetName, dataRange)
	if err != nil {
		http.Error(w, "Cannot delete the rows requested", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Delete successfully!")
}

func DeleteDataCellHelper(spreadsheetID string, sheetName string, dataRange [][]interface{}) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	request := &sheets.ClearValuesRequest{}

	for _, pos := range dataRange {
		row := pos[0].(string)
		col_int, err := strconv.Atoi(pos[1].(string))
		if err != nil {
			return err
		}
		col := read.ColumnIndexToLetter(col_int)
		rowRange := sheetName + "!" + col + row + ":" + col + row

		_, err = service.Spreadsheets.Values.Clear(spreadsheetID, rowRange, request).Do()
		if err != nil {
			return err
		}
	}
	return nil
}

/*
DELETE
Body: {"spreadsheetID": "YOUR_SPREAD_SHEET_ID"}
*/
func DeleteSpreadsheet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req struct {
		SpreadsheetID string `json:"spreadsheetID"`
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

	err = DeleteSpreadsheetHelper(spreadsheetID)
	if err != nil {
		http.Error(w, "Cannot delete spreadsheet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Message       string `json:"message"`
	}{
		SpreadsheetID: spreadsheetID,
		Message:       "Spreadsheet deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteSpreadsheetHelper(spreadsheetID string) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	_, err = service.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return err
	}

	driveService, err := svc.SetupGoogleDriveService()
	if err != nil {
		return fmt.Errorf("unable to create drive service: %v", err)
	}

	err = driveService.Files.Delete(spreadsheetID).Do()
	if err != nil {
		return fmt.Errorf("failed to delete spreadsheet: %v", err)
	}

	return nil
}

/*
DELETE
Body: {"spreadsheetID": "YOUR_SPREAD_SHEET_ID", "sheetID": 123456}
*/
func DeleteSheet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	err = DeleteSheetHelper(spreadsheetID, req.SheetID)
	if err != nil {
		http.Error(w, "Cannot delete sheet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		SpreadsheetID string `json:"spreadsheetID"`
		SheetID       int64  `json:"sheetID"`
		Message       string `json:"message"`
	}{
		SpreadsheetID: spreadsheetID,
		SheetID:       req.SheetID,
		Message:       "Sheet deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteSheetHelper(spreadsheetID string, sheetID int64) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	req := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				DeleteSheet: &sheets.DeleteSheetRequest{
					SheetId: sheetID,
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
