package fetcher


import (
	"fmt"
	"strconv"
	"net/http"
	"math/rand"
	"encoding/json"
)

type API interface {
	Fetch(arg map[string]string) []map[string]string
}


type APIs []API


func validDigit(dgt string) bool {
	return len(dgt) == 1 && dgt[0] - '0' < 10
}


func validFloat(flt string) bool {
	_, err := strconv.ParseFloat(flt, 64)
	return err == nil
}


func (a APIs) Fetch(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Not Found", 404)
		return
	}
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
	fmt.Println("GET /roll:")
	for i := 0; i < 4; i++ {
		fmt.Printf("  Attempt %d:\n", i+1)
		api := a[rand.Intn(len(a))]
		ret := api.Fetch(arg)
		if len(ret) == 0 {
			continue
		}
		act := ret[rand.Intn(len(ret))]
		if jsn, err := json.Marshal(act); err == nil {
			res.Write(jsn)
			return
		}
	}
	http.Error(res, "Internal Server Error", 500)
}
