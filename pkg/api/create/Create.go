package create

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"personnel-api/pkg/api/read"
	"personnel-api/pkg/svc"

	"google.golang.org/api/sheets/v4"
)

/*
POST
Body: {"spreadsheetID": "YOUR_SPREAD_SHEET_ID", "sheetName": "SHEET_NAME", "rows":[ ["3", "test1", "test1@gmail.com"], ["4", "test2", "test2@gmail.com"] ]}
*/

/*
POST
Body: {"title": "Spreadsheet Title"}
*/
func CreateSpreadsheet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req struct {
		Title string `json:"title"`
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	title := req.Title
	if title == "" {
		title = "New Spreadsheet"
	}

	spreadsheetID, err := CreateSpreadsheetHelper(title)
	if err != nil {
		http.Error(w, "Cannot create new spreadsheet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Title         string `json:"title"`
	}{
		SpreadsheetID: spreadsheetID,
		Title:         title,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateSpreadsheetHelper(title string) (string, error) {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return "", err
	}

	spreadsheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: title,
		},
		Sheets: []*sheets.Sheet{
			{
				Properties: &sheets.SheetProperties{
					Title: "Sheet1",
				},
			},
		},
	}

	createdSpreadsheet, err := service.Spreadsheets.Create(spreadsheet).Do()
	if err != nil {
		return "", err
	}

	return createdSpreadsheet.SpreadsheetId, nil
}

// check for valid length of input not included (each data in rows has to match what is in the sheet)
func CreateData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req struct {
		SpreadsheetID string          `json:"spreadsheetID"`
		SheetName     string          `json:"sheetName"`
		Rows          [][]interface{} `json:"rows"`
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
		http.Error(w, "spreadsheetID field is required", http.StatusBadRequest)
		return
	}

	dataRange, _, _ := read.GetSheetDataHelper(spreadsheetID, sheetName)
	dataRange = sheetName + "!" + dataRange

	rows := req.Rows
	if len(rows) == 0 {
		http.Error(w, "rows data field is required", http.StatusBadRequest)
		return
	}

	err = CreateDataHelper(spreadsheetID, dataRange, rows)
	if err != nil {
		http.Error(w, "Cannot create new rows in sheet", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Insert successfully!")
}

func CreateDataHelper(spreadsheetID string, dataRange string, rows [][]interface{}) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	valueRange := &sheets.ValueRange{
		Values: rows,
	}

	_, err = service.Spreadsheets.Values.Append(spreadsheetID, dataRange, valueRange).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return err
	}

	return nil
}

/*
POST
Body: {"spreadsheetID": "YOUR_SPREAD_SHEET_ID", "sheetName": "NEW_SHEET_NAME"}
*/
func CreateSheet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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
		SheetName     string `json:"sheetName"`
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

	err = CreateSheetHelper(spreadsheetID, sheetName)
	if err != nil {
		http.Error(w, "Cannot create new sheet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		SpreadsheetID string `json:"spreadsheetID"`
		SheetName     string `json:"sheetName"`
		Message       string `json:"message"`
	}{
		SpreadsheetID: spreadsheetID,
		SheetName:     sheetName,
		Message:       "Sheet created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateSheetHelper(spreadsheetID string, sheetName string) error {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return err
	}

	req := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddSheet: &sheets.AddSheetRequest{
					Properties: &sheets.SheetProperties{
						Title: sheetName,
					},
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
