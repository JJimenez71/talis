package fetcher


import (
	"fmt"
	"strconv"
	"net/http"
	"math/rand"
	"encoding/json"
)


type Activity struct {
	Name		string
	Address		string
	Hours		[]string
	Phone		string
	Website		string
	Rating		string
	Price		string
	Description	string
}


type Activities []Activity


type API interface {
	Fetch(arg map[string]string) Activities
}


type APIs []API


func validDigit(dgt string) bool {
	return len(dgt) == 1 && dgt[0] - '0' < 10
}


func validFloat(flt string) bool {
	_, err := strconv.ParseFloat(flt, 64)
	return err == nil
}

func validParameters(arg map[string]string) string {
	if !validDigit(arg["distance"]) {
		return "Bad Request: query parameter \"distance\"."
	}
	if !validDigit(arg["expense"]) {
		return "Bad Request: query parameter \"expense\"."
	}
	if !validFloat(arg["latitude"]) {
		return "Bad Request: query parameter \"latitude\"."
	}
	if !validFloat(arg["longitude"]) {
		return "Bad Request: query parameter \"longitude\"."
	}
	return ""
}

func (a APIs) Fetch(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Not Found", 404)
		return
	}
	qry := req.URL.Query()
	arg := map[string]string{}
	arg["distance"] = qry.Get("distance")
	arg["expense"] = qry.Get("expense")
	arg["latitude"] = qry.Get("latitude")
	arg["longitude"] = qry.Get("longitude")
	if err := validParameters(arg); err != "" {
		http.Error(res, err, 400)
		return
	}
	fmt.Println("GET /roll:")
	for i := 0; i < 4; i++ {
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
