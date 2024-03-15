package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
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
	dis := qry.Get("distance")
	if !validDigit(dis) {
		http.Error(res, "Bad Request: query parameter \"distance\".", 400)
		return
	}
	exp := qry.Get("expense")
	if !validDigit(exp) {
		http.Error(res, "Bad Request: query parameter \"expense\".", 400)
		return
	}
	lat := qry.Get("latitude")
	if !validFloat(lat) {
		http.Error(res, "Bad Request: query parameter \"latitude\".", 400)
		return
	}
	lon := qry.Get("longitude")
	if !validFloat(lon) {
		http.Error(res, "Bad Request: query parameter \"longitude\".", 400)
		return
	}
	panic("endpoint not implemented")
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
