package server

import (
	"encoding/json"
	"fmt"
	"github.com/aabiji/read/epub"
	"github.com/gorilla/mux"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
    j, _ := json.Marshal(map[string]string{
        "Server error": fmt.Sprintf("%s", err.Error()),
    })
    w.Write(j)
}

func getBookInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := fmt.Sprintf("epub/tests/%s.epub", vars["name"])
	w.Header().Set("Content-Type", "application/json")

	book, err := epub.New(path)
    if err != nil {
        handleError(w, err)
        return
    }

	json.NewEncoder(w).Encode(book)
}