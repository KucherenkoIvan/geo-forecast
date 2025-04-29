package controllers

import (
	"fmt"
	"net/http"
)

func AppInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Module: [ geo-forecast ]; Version: [ 0.0.1_alpha.1 ]")
}
