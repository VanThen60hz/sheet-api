package main

import (
	"log"
	"net/http"

	"personnel-api/pkg/api/create"
	"personnel-api/pkg/api/delete"
	"personnel-api/pkg/api/read"
	"personnel-api/pkg/api/update"
	"personnel-api/pkg/authorization"
	"personnel-api/pkg/middleware"

	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

var enforcer *casbin.Enforcer

func main() {
	model := "model.conf"
	policy := "policy.csv"

	adapter := fileadapter.NewAdapter(policy)
	if adapter == nil {
		log.Fatal(adapter)
	}

	var err error
	enforcer, err = casbin.NewEnforcer(model, adapter)
	if err != nil {
		log.Fatal(err)
	}

	// Register routes
	registerReadRoutes()
	registerCreateRoutes()
	registerUpdateRoutes()
	registerDeleteRoutes()
	registerAuthRoutes()

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerReadRoutes() {
	readRoutes := map[string]http.HandlerFunc{
		"/GetAll":              read.GetAll,
		"/GetSheetData":        read.GetSheetData,
		"/GetByColumn":         read.GetByColumn,
		"/GetByFilter":         read.GetByFilter,
		"/GetSheets":           read.GetSheets,
		"/ListAllSpreadsheets": read.ListAllSpreadsheets,
		"/GetSpreadsheetById":  read.GetSpreadsheetById,
	}

	for path, handler := range readRoutes {
		http.HandleFunc(path, middleware.EnableCORS(middleware.Authorize(enforcer)(handler)))
	}
}

func registerCreateRoutes() {
	createRoutes := map[string]http.HandlerFunc{
		"/CreateData":        create.CreateData,
		"/CreateSpreadsheet": create.CreateSpreadsheet,
		"/CreateSheet":       create.CreateSheet,
	}

	for path, handler := range createRoutes {
		http.HandleFunc(path, middleware.EnableCORS(middleware.Authorize(enforcer)(handler)))
	}
}

func registerUpdateRoutes() {
	updateRoutes := map[string]http.HandlerFunc{
		"/UpdateDataRow":     update.UpdateDataRow,
		"/UpdateDataCell":    update.UpdateDataCell,
		"/UpdateSpreadsheet": update.UpdateSpreadsheet,
		"/UpdateSheet":       update.UpdateSheet,
	}

	for path, handler := range updateRoutes {
		http.HandleFunc(path, middleware.EnableCORS(middleware.Authorize(enforcer)(handler)))
	}
}

func registerDeleteRoutes() {
	deleteRoutes := map[string]http.HandlerFunc{
		"/DeleteDataRow":     delete.DeleteDataRow,
		"/DeleteDataCell":    delete.DeleteDataCell,
		"/DeleteSpreadsheet": delete.DeleteSpreadsheet,
		"/DeleteSheet":       delete.DeleteSheet,
	}

	for path, handler := range deleteRoutes {
		http.HandleFunc(path, middleware.EnableCORS(middleware.Authorize(enforcer)(handler)))
	}
}

func registerAuthRoutes() {
	authRoutes := map[string]http.HandlerFunc{
		"/AddPolicy":    authorization.AddPolicy,
		"/RemovePolicy": authorization.RemovePolicy,
	}

	for path, handler := range authRoutes {
		http.HandleFunc(path, middleware.EnableCORS(middleware.Authorize(enforcer)(handler)))
	}
}
