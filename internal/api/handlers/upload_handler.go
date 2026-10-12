// Package handlers implements the HTTP handlers for the API endpoints.
package handlers

import (
	"net/http"

	"github.com/kbg05/pixstash/internal/database"
	"github.com/kbg05/pixstash/internal/storage"
)

type UploadHandler struct{
	q *database.Queries
	store storage.BlobStore
}

func (h *UploadHandler) handle(w http.ResponseWriter, r *http.Request){
	// Implementation of the upload handlers
	r.Body=http.MaxBytesReader(w, r.Body, 100<<20) // 100 MB limit
	err:=r.ParseMultipartForm(100 << 20)
	if err!=nil{
		http.Error (w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, header, err:=r.FormFile("file")
	if err!=nil{
		http.Error(w, "failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()
	blobInfo, err:=h.store.Put(r.Context(), file, header.Filename)
	if err!=nil{
		http.Error(w, "failed to store file", http.StatusInternalServerError)
		return
	}

}



