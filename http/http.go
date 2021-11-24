package handler

import (
	"fmt"
	"net/http"
	"ImageRandom/control"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	image := control.Random()
	fmt.Fprintf(w, "" + image)
}