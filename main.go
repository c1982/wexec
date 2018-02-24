package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	port := os.Args[1]

	http.HandleFunc("/do", executeHandler)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("serve error: ", err)
	}
}

func executeHandler(w http.ResponseWriter, r *http.Request) {

	var output []byte
	var err error

	script := r.FormValue("file")

	if output, err = exec.Command(script).Output(); err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(output))
}
