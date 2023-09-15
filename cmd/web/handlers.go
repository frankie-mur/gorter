package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *application) urlFind(w http.ResponseWriter, r *http.Request) {
	url, err := app.urls.FindOriginalUrl("test.com")
	if err != nil {
		w.Write([]byte("An error occurred"))
		return
	}
	http.Redirect(w, r, *url, http.StatusFound)
	//w.Write([]byte(fmt.Sprintf("Found a single document: %+v\n", url)))
}

func (app *application) urlCreate(w http.ResponseWriter, r *http.Request) {
	err := app.urls.CreateUrl("test.com", "https://www.google.com", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created url\n")
}
