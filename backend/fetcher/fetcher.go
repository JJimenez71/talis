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
	api := a[rand.Intn(len(a))]
	par := api.QryRul.Parameters(arg)


	out := api.Hos + "/" + api.Pat + "?"
	for key, val := range par {
		out += key + "=" + val + "&"
	}
	io.WriteString(res, out[:len(out)-1] + "\n")
}
