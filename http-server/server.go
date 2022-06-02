package main

import (
	"io"
	"net/http"
	"net/http/httputil"
)

func echo(w http.ResponseWriter, req *http.Request) {
	dumpRequest(w, "echo", req)
}

func RunServer() {
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":8090", nil)
}

func dumpRequest(writer io.Writer, header string, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	writer.Write([]byte("\n" + header + ": \n"))
	writer.Write(data)
	return nil
}
