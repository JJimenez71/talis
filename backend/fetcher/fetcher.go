package fetcher


import (
	"io"
	"strconv"
	"net/http"
	"math/rand"
	"backend/mapper"
)


type API struct {
	Hos string
	Pat string
	QryRul mapper.Rules
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
	arg := map[string]string{}
	qry := req.URL.Query()
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
	api := a[rand.Intn(len(a))]
	par := api.QryRul.Parameters(arg)



	out := api.Hos + "/" + api.Pat + "?"
	for key, val := range par {
		out += key + "=" + val + "&"
	}
	io.WriteString(res, out[:len(out)-1] + "\n")
}
