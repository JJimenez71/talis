package main


import (
	"fmt"
	"io"
	"net/http"
	"time"
	"backend/config"
	"backend/fetcher"
	"backend/vendors"
)


func getRoot(res http.ResponseWriter, req *http.Request) {
	/*
	select {
	case <-req.Context().Done():
		fmt.Printf("Timeout: Get /\n")
	}
	*/
	if req.Method != "GET" {
		http.Error(res, "Not Found", 404)
		return
	}
	fmt.Println("GET /")
	io.WriteString(res, "Try out our /roll GET endpoint!\n")
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
	cfg := config.Load(".env")
	api := fetcher.APIs {
		//vendors.BarAPI(cfg),
		//vendors.FooAPI(cfg),
		vendors.YelpAPI(cfg),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", panicWrapper(getRoot))
	mux.HandleFunc("/roll", panicWrapper(api.Fetch))
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
		panic("cannot start configured server")
	}
}
