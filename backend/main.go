package main

import (
	"fmt"
	"jobsalign/routes"
	"jobsalign/tools"
	"net/http"

	"github.com/unidoc/unipdf/v3/common/license"
)

func init() {
	err := license.SetMeteredKey("4dbaca401330dd4c67d4164d8a49c606e8ce6b5060accd23df050308a80590ab")
	if err != nil {
		panic(err)
	}
}

func main() {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/test", tools.HandleWithCORS(routes.ResumeProcess))

	err := http.ListenAndServe(":8080", serverMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
