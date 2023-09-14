package main

import (
	"fmt"
	"net/http"
)

func (app *application) urlCreate(w http.ResponseWriter, r *http.Request) {
	url, err := app.urls.FindById("6502959c94d6a9aec3a91dce")
	if err != nil {
		w.Write([]byte("An error occurred"))
	}
	w.Write([]byte(fmt.Sprintf("Found a single document: %+v\n", url)))
}
