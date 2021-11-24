package handler

import (
	"ImageRandom/control"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	image := control.Random()
	fmt.Fprintf(w, "" + image)
}