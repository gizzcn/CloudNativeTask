package main

import (
	"github.com/tav/golly/log"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/source", source)
	err := http.ListenAndServe(":8083", nil)
	if err != nil{
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request){

	ver := os.Getenv("VERSION")
	w.Header().Add("VERSION", ver)
	for key, value := range r.Header{
		w.Header().Add(key, value[0])
	}
	io.WriteString(w, "200")
	io.WriteString(w, " Source: http://101.200.230.148:8083/source")
}

func source(w http.ResponseWriter, r *http.Request){
	f, err := ioutil.ReadFile("main.go")
	if err != nil {
		io.WriteString(w, "read fail: " + err.Error())
	}
	io.WriteString(w, string(f))
}