package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// uploadHandler handles the file upload request
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// check if the request method is POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request method")
		return
	}
	// parse the multipart form data from the request body
	err := r.ParseMultipartForm(10 << 20) // 10 MB max size
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to parse form data:", err)
		return
	}
	// get the file header from the form data
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No file found in form data:", err)
		return
	}
	defer file.Close() // close the file handle when done

	// create a new file with the same name as the uploaded file
	dst, err := os.Create(header.Filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to create file:", err)
		return
	}
	defer dst.Close() // close the file handle when done

	// copy the contents of the uploaded file to the new file
	n, err := io.Copy(dst, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to copy file:", err)
		return
	}

	// write a success message to the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %s uploaded successfully with size %d bytes\n", header.Filename, n)
}

func main() {
	// create a new HTTP server mux
	mux := http.NewServeMux()

	// register the upload handler for the "/upload" path
	mux.HandleFunc("/upload", uploadHandler)

	// start listening on port 8080
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
