package read

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"personnel-api/pkg/svc"
	"strconv"
	"strings"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"
)

// GET
// Body: {"spreadsheetID": "YOUR_SPREAD_SHEET_ID"}
func GetAll(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req map[string]string
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	spreadsheetID, ok := req["spreadsheetID"]
	if !ok || spreadsheetID == "" {
		http.Error(w, "spreadsheetID field is required", http.StatusBadRequest)
		return
	}

	allData, err := GetAllHelper(spreadsheetID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve data from all sheets: %v", err), http.StatusInternalServerError)
		return
	}

	dataJSON, err := json.Marshal(allData)
	if err != nil {
		http.Error(w, "Failed to convert data to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}

func GetAllHelper(spreadsheetID string) ([]interface{}, error) {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return nil, err
	}
	spreadsheet, err := service.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve spreadsheet: %v", err)
	}
	var allData []interface{}

	for _, sheet := range spreadsheet.Sheets {
		_, sheetData, err := GetSheetDataHelper(spreadsheetID, sheet.Properties.Title)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve sheet data: %v", err)
		}

		allData = append(allData, sheetData)
	}
	return allData, nil
}

// GET
// Body: {"spreadsheetID": "YOUR_SPREAD_SHEET_ID", "sheetName": "SHEET_NAME"}
func GetSheetData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req map[string]string
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	spreadsheetID, ok := req["spreadsheetID"]
	if !ok || spreadsheetID == "" {
		http.Error(w, "spreadsheetID field is required", http.StatusBadRequest)
		return
	}

	sheetName, ok := req["sheetName"]
	if !ok || sheetName == "" {
		http.Error(w, "sheetName field is required", http.StatusBadRequest)
	}

	_, data, err := GetSheetDataHelper(spreadsheetID, sheetName)
	if err != nil {
		http.Error(w, "failed to retrieve data from sheet", http.StatusBadRequest)
		return
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to convert data to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}

func GetSheetDataHelper(spreadsheetID string, sheetName string) (string, []interface{}, error) {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return "", nil, err
	}

	spreadsheet, err := service.Spreadsheets.Values.Get(spreadsheetID, sheetName).Do()
	if err != nil {
		return "", nil, fmt.Errorf("failed to retrieve spreadsheet: %v", err)
	}

	var allData []interface{}

	if len(spreadsheet.Values) == 0 {
		return "", allData, nil
	}

	startRow := 0
	startColumn := 0

	if len(spreadsheet.Values) > 0 {
		for i, row := range spreadsheet.Values {
			if len(row) > 0 {
				startRow = i
				break
			}
		}

		if startRow > 0 {
			for j, value := range spreadsheet.Values[startRow] {
				if value != nil && value != "" {
					startColumn = j
					break
				}
			}
		}
	}

	data := spreadsheet.Values[startRow:]
	for i, row := range data {
		data[i] = row[startColumn:]
	}
	allData = append(allData, data)

	dataRange := ColumnIndexToLetter(startColumn) + ":" + ColumnIndexToLetter(startColumn+len(data[0])-1)
	return dataRange, allData, nil
}

func ColumnIndexToLetter(index int) string {
	var result string
	for {
		result = string(rune('A'+(index%26))) + result
		if index /= 26; index <= 0 {
			break
		}
	}
	return result
}

// GET
// Body: {"spreadsheetID": "YOUR_SPREAD_SHEET_ID", "sheetName": "SHEET_NAME", "columnName": "COLUMN_NAME"}
func GetByColumn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req map[string]string
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	spreadsheetID, ok := req["spreadsheetID"]
	if !ok || spreadsheetID == "" {
		http.Error(w, "spreadsheetID field is required", http.StatusBadRequest)
		return
	}

	sheetName, ok := req["sheetName"]
	if !ok || sheetName == "" {
		http.Error(w, "sheetName field is required", http.StatusBadRequest)
	}

	columnName, ok := req["columnName"]
	if !ok || columnName == "" {
		http.Error(w, "sheetName field is required", http.StatusBadRequest)
	}

	_, data, err := GetByColumnHelper(spreadsheetID, sheetName, columnName)
	if err != nil {
		http.Error(w, "cannot extract column data", http.StatusInternalServerError)
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to convert data to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}

func GetByColumnHelper(spreadsheetID string, sheetName string, columnName string) (int, []interface{}, error) {
	_, sheetData, err := GetSheetDataHelper(spreadsheetID, sheetName)
	if err != nil {
		return -1, nil, fmt.Errorf("failed to retrieve spreadsheet data: %v", err)
	}

	columnIdx := -1

	for i, name := range sheetData[0].([][]interface{})[0] {
		if name == columnName {
			columnIdx = i
			break
		}
	}

	if columnIdx == -1 {
		return -1, nil, fmt.Errorf("no column found with that name: %v", err)
	}

	var allData []interface{}
	for _, row := range sheetData {
		for _, item := range row.([][]interface{}) {
			if len(item) == 0 || len(item) <= columnIdx {
				continue
			}
			allData = append(allData, item[columnIdx])
		}
	}

	return columnIdx, allData, nil
}

// GET
// Body: {"spreadsheetID": "YOUR_SPREAD_SHEET_ID", "sheetName": "SHEET_NAME", "columnName": "COLUMN_NAME", "operator": "OP", "value": "VALUE"}
func GetByFilter(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var req map[string]string
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	spreadsheetID, ok := req["spreadsheetID"]
	if !ok || spreadsheetID == "" {
		http.Error(w, "spreadsheetID field is required", http.StatusBadRequest)
		return
	}

	sheetName, ok := req["sheetName"]
	if !ok || sheetName == "" {
		http.Error(w, "sheetName field is required", http.StatusBadRequest)
	}

	columnName, ok := req["columnName"]
	if !ok || columnName == "" {
		http.Error(w, "sheetName field is required", http.StatusBadRequest)
	}

	operator, ok := req["operator"]
	if !ok || operator == "" {
		http.Error(w, "operator field is required", http.StatusBadRequest)
	}

	value, ok := req["value"]
	if !ok || value == "" {
		http.Error(w, "value field is required", http.StatusBadRequest)
	}

	filteredData, err := GetByFilterHelper(spreadsheetID, sheetName, columnName, operator, value)
	if err != nil {
		http.Error(w, "cannot filter column data", http.StatusInternalServerError)
	}

	dataJSON, err := json.Marshal(filteredData)
	if err != nil {
		http.Error(w, "Failed to convert data to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
}

func GetByFilterHelper(spreadsheetID string, sheetName string, columnName string, operator string, value string) ([]interface{}, error) {
	_, sheetData, err := GetSheetDataHelper(spreadsheetID, sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve sheet data: %v", err)
	}

	columnIdx, columnData, err := GetByColumnHelper(spreadsheetID, sheetName, columnName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve column data: %v", err)
	}

	dataType := 0
	valueType := 0

	var filteredData []interface{}
	subList, ok := sheetData[0].([][]interface{})
	if ok {
		columnTitle := subList[0]
		filteredData = append(filteredData, columnTitle)

	}

	if len(columnData) <= 1 {
		return filteredData, nil
	} else {
		dataType, err = CheckStringType(columnData[1].(string))
		if err != nil {
			return nil, fmt.Errorf("unknown type: %v", err)
		}
		valueType, err = CheckStringType(value)
		if err != nil {
			return nil, fmt.Errorf("unknown type: %v", err)
		}
		if dataType != 3 && valueType != dataType {
			return nil, fmt.Errorf("incorrect value type: %v", err)
		}
	}

	for _, cell := range sheetData[0].([][]interface{})[1:] {
		if len(cell) == 0 || len(cell) <= columnIdx {
			continue
		}
		if dataType == 1 {
			valueInt, _ := strconv.Atoi(value)
			cellInt, _ := strconv.Atoi(cell[columnIdx].(string))

			switch operator {
			case "=":
				if cellInt == valueInt {
					filteredData = append(filteredData, cell)
				}
			case ">":
				if cellInt > valueInt {
					filteredData = append(filteredData, cell)
				}
			case "<":
				if cellInt < valueInt {
					filteredData = append(filteredData, cell)
				}
			default:
				return nil, fmt.Errorf("unsupported operator: %s", operator)
			}
		} else if dataType == 2 {
			valueFloat, _ := strconv.ParseFloat(value, 64)
			cellFloat, _ := strconv.ParseFloat(cell[columnIdx].(string), 64)
			switch operator {
			case "=":
				if cellFloat == valueFloat {
					filteredData = append(filteredData, cell)
				}
			case ">":
				if cellFloat > valueFloat {
					filteredData = append(filteredData, cell)
				}
			case "<":
				if cellFloat < valueFloat {
					filteredData = append(filteredData, cell)
				}
			default:
				return nil, fmt.Errorf("unsupported operator: %s", operator)
			}
		} else {
			switch operator {
			case "contain":
				if strings.Contains(cell[columnIdx].(string), value) {
					filteredData = append(filteredData, cell)
				}
			default:
				return nil, fmt.Errorf("unsupported operator: %s", operator)
			}
		}
	}

	return filteredData, nil
}

func CheckStringType(value string) (int, error) {
	if _, err := strconv.Atoi(value); err == nil {
		return 1, nil
	}

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return 2, nil
	}

	return 3, nil
}

/*
GET
Query params: spreadsheetID=YOUR_SPREAD_SHEET_ID
*/
func GetSheets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	spreadsheetID := r.URL.Query().Get("spreadsheetID")
	if spreadsheetID == "" {
		http.Error(w, "spreadsheetID parameter is required", http.StatusBadRequest)
		return
	}

	sheets, err := GetSheetsHelper(spreadsheetID)
	if err != nil {
		http.Error(w, "Cannot get sheets info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		SpreadsheetID string `json:"spreadsheetID"`
		Sheets        []struct {
			SheetID    int64  `json:"sheetID"`
			SheetName  string `json:"sheetName"`
			SheetIndex int64  `json:"sheetIndex"`
		} `json:"sheets"`
	}{
		SpreadsheetID: spreadsheetID,
		Sheets: make([]struct {
			SheetID    int64  `json:"sheetID"`
			SheetName  string `json:"sheetName"`
			SheetIndex int64  `json:"sheetIndex"`
		}, 0),
	}

	for _, sheet := range sheets {
		response.Sheets = append(response.Sheets, struct {
			SheetID    int64  `json:"sheetID"`
			SheetName  string `json:"sheetName"`
			SheetIndex int64  `json:"sheetIndex"`
		}{
			SheetID:    sheet.Properties.SheetId,
			SheetName:  sheet.Properties.Title,
			SheetIndex: sheet.Properties.Index,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetSheetsHelper(spreadsheetID string) ([]*sheets.Sheet, error) {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return nil, err
	}

	spreadsheet, err := service.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return nil, err
	}

	return spreadsheet.Sheets, nil
}

// GET
// No body required
// Returns a list of all spreadsheets accessible to the user
func ListAllSpreadsheets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	spreadsheets, err := ListAllSpreadsheetsHelper()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to list spreadsheets: %v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		Spreadsheets []struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			CreatedTime  string `json:"createdTime"`
			ModifiedTime string `json:"modifiedTime"`
		} `json:"spreadsheets"`
	}{
		Spreadsheets: make([]struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			CreatedTime  string `json:"createdTime"`
			ModifiedTime string `json:"modifiedTime"`
		}, 0),
	}

	for _, spreadsheet := range spreadsheets {
		response.Spreadsheets = append(response.Spreadsheets, struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			CreatedTime  string `json:"createdTime"`
			ModifiedTime string `json:"modifiedTime"`
		}{
			ID:           spreadsheet.Id,
			Name:         spreadsheet.Name,
			CreatedTime:  spreadsheet.CreatedTime,
			ModifiedTime: spreadsheet.ModifiedTime,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListAllSpreadsheetsHelper() ([]*drive.File, error) {
	service, err := svc.SetupGoogleDriveService()
	if err != nil {
		return nil, fmt.Errorf("failed to setup Google Drive service: %v", err)
	}

	// Query to find all Google Sheets files
	query := "mimeType='application/vnd.google-apps.spreadsheet'"

	// List all spreadsheets
	results, err := service.Files.List().
		Q(query).
		Fields("files(id, name, createdTime, modifiedTime)").
		OrderBy("modifiedTime desc").
		Do()
	if err != nil {
		return nil, fmt.Errorf("failed to list spreadsheets: %v", err)
	}

	return results.Files, nil
}

// GET
// Query params: spreadsheetID=YOUR_SPREAD_SHEET_ID
// Returns basic information about a specific spreadsheet
func GetSpreadsheetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	spreadsheetID := r.URL.Query().Get("spreadsheetID")
	if spreadsheetID == "" {
		http.Error(w, "spreadsheetID parameter is required", http.StatusBadRequest)
		return
	}

	spreadsheet, err := GetSpreadsheetByIdHelper(spreadsheetID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get spreadsheet info: %v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}{
		ID:    spreadsheet.SpreadsheetId,
		Title: spreadsheet.Properties.Title,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetSpreadsheetByIdHelper(spreadsheetID string) (*sheets.Spreadsheet, error) {
	service, err := svc.SetupGoogleSheetsService()
	if err != nil {
		return nil, fmt.Errorf("failed to setup Google Sheets service: %v", err)
	}

	spreadsheet, err := service.Spreadsheets.Get(spreadsheetID).Fields("spreadsheetId,properties(title)").Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get spreadsheet: %v", err)
	}

	return spreadsheet, nil
}
