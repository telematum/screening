package main

import (
	"net/http"
)

func main() {
	setupJsonApi()
	http.ListenAndServe(":80", nil)
}
