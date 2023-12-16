package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

func upload(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	r.ParseMultipartForm(10 << 20) // 10MB

	// Retrieve file from the form data
	file, handler, err := r.FormFile("MyFile")
	if err != nil {
		log.Println("Error retrieving the file")
		log.Println(err)
		return 
	}
	
	defer file.Close()
	log.Println("Uploaded File: %+v\n", handler.Filename)
	log.Println("File Size: %+v\n", handler.Size)
	log.Println("MIME Header: %+v\n", handler.Header)

	// Write the file to the Server
	tempFile, err := ioutil.TempFile("uploads", "upload-*.png")
	if err != nil {
		log.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8000", nil)
}

func main() {
	fmt.Println("Go File Upload Server Started at http://localhost:8000")
    setupRoutes()
}