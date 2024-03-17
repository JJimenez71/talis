package main


import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"backend/mapper"
)


func getRoot(res http.ResponseWriter, req *http.Request) {
	/*
	select {
	case <-req.Context().Done():
		fmt.Printf("Timeout: Get /\n")
	}
	*/
	io.WriteString(res, "Try out our /roll GET endpoint!\n")
	fmt.Printf("Get /\n")
}


func validDigit(dgt string) bool {
	return len(dgt) == 1 && dgt[0] - '0' < 10
}


func validFloat(flt string) bool {
	_, err := strconv.ParseFloat(flt, 64)
	return err == nil
}


func getRoll(res http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()
	arg := map[string]string{}
	arg["distance"] = qry.Get("distance")
	if !validDigit(arg["distance"]) {
		http.Error(res, "Bad Request: query parameter \"distance\".", 400)
		return
	}
	arg["expense"] = qry.Get("expense")
	if !validDigit(arg["expense"]) {
		http.Error(res, "Bad Request: query parameter \"expense\".", 400)
		return
	}
	arg["latitude"] = qry.Get("latitude")
	if !validFloat(arg["latitude"]) {
		http.Error(res, "Bad Request: query parameter \"latitude\".", 400)
		return
	}
	arg["longitude"] = qry.Get("longitude")
	if !validFloat(arg["longitude"]) {
		http.Error(res, "Bad Request: query parameter \"longitude\".", 400)
		return
	}
	rul := []mapper.Rule{
		// "apiKey" hardcode value
		{"", "", "apiKey", "A1B2C3D4E5F6", ""},
		// "coords" mapping
		{"latitude", "", "coords", "", ","},
		{"longitude", "", "coords", "", ","},
		// "range" default value
		{"", "", "range", "100", ""},
		// "range" mapping
		{"distance", "0", "range", "5", ""},
		{"distance", "1", "range", "10", ""},
		{"distance", "2", "range", "20", ""},
		// "minCost"/"maxCost" default value
		{"", "", "minCost", "$", ""},
		{"", "", "maxCost", "$$$", ""},
		// "minCost"/"maxCost" mapping
		{"expense", "0", "minCost", "$", ""},
		{"expense", "0", "maxCost", "$$", ""},
		{"expense", "1", "minCost", "$$", ""},
		{"expense", "1", "maxCost", "$$$", ""},
		{"expense", "2", "minCost", "$$$", ""},
		{"expense", "2", "maxCost", "$$$$", ""},
	}
	kvp := mapper.FormQuery(rul, arg)
	for key, val := range kvp {
		io.WriteString(res, "&" + key + "=" + val)
	}
	io.WriteString(res, "\n")
	//panic("endpoint not implemented")
}


func methodWrapper(fun func(http.ResponseWriter, *http.Request), met string) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		if met != req.Method {
			http.Error(res, "Not Found", 404)
			return
		}
		fun(res, req)
	}
}


func panicWrapper(fun func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(res, "Internal Server Error", 500)
				fmt.Printf("Panic: %v\n", err)
			}
		}()
		fun(res, req)
	}
}


func configureServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", panicWrapper(methodWrapper(getRoot, "GET")))
	mux.HandleFunc("/roll", panicWrapper(methodWrapper(getRoll, "GET")))
	han := http.TimeoutHandler(mux, 5 * time.Second, "Gateway Timeout\n")
	svr := &http.Server{
		ReadTimeout:	5 * time.Second,
		WriteTimeout:	5 * time.Second,
		Addr:		":8080",
		Handler:	han, //mux
	}
	return svr
}


func main() {
        svr := configureServer()
	if err := svr.ListenAndServe(); err != nil {
		fmt.Printf("Server error two: %s\n", err)
	}
}
