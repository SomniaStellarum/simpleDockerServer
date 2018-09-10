package main

import (
	"net/http"
	"github.com/SomniaStellarum/StellarUtilities/slog"
)

func main() {
	slog.SetDebug()
	http.Handle("/static/", slog.DebugPrintURL(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/html")))))
	http.ListenAndServe(":8080", nil)
}