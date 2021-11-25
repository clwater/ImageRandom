package handler

import (
	"ImageRandom/control"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	image := control.Random()
	//fmt.Fprintf(w, "" + image)

	http.Redirect(w, req, image, http.StatusTemporaryRedirect)
}