package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (app *application) urlFind(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "*")
	url, err := app.urls.FindOriginalUrl(shortUrl)
	if err != nil {
		w.Write([]byte("An error occurred"))
		return
	}
	http.Redirect(w, r, *url, http.StatusFound)
	//w.Write([]byte(fmt.Sprintf("Found a single document: %+v\n", url)))
}

type urlPost struct {
	OriginalURL string `json:"originalUrl"`
	ShortUrl    string `json:"shortUrl"`
}

func (p *urlPost) Bind(r *http.Request) error {
	//Validate that fields are present
	if p.OriginalURL == "" || p.ShortUrl == "" {
		return errors.New("invalid data")
	}

	return nil
}

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func (app *application) urlCreate(w http.ResponseWriter, r *http.Request) {
	data := &urlPost{}
	err := render.Bind(r, data)
	if err != nil {
		//Send back 400
		render.Render(w, r, ErrInvalidRequest(err))

	}
	err = app.urls.CreateUrl(data.ShortUrl, data.OriginalURL)
	if err != nil {
		log.Fatal(err)
	}
}
