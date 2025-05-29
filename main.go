/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/monakhovm/gcloud-version/cmd"
)

func main() {
	// // Якщо додаток запущено як CLI-команда
	// if len(os.Args) > 1 && os.Args[1] != "serve" {
	// 	cmd.Execute()
	// 	return
	// }

	// Web-сервер
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", cmd.AppVersion())
	})

	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
