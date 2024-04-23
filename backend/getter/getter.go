package getter


import (
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
)


type Rule struct {
	App string
	Arg string
	Equ string
	Hed string
	Qry string
	Val string
}


type Mapping []Rule


type Request struct {
	Hos string
	Map Mapping
	Pat string
}


func (m Mapping) headers(arg map[string]string) map[string]string {
	hed := map[string]string{}
	for _, r := range m {
		if r.Hed == "" {
			continue
		}
		if r.Equ != "" && arg[r.Arg] != r.Equ {
			continue
		}
		val := r.Val
		if val == "" {
			val = arg[r.Arg]
		}
		if _, ok := hed[r.Hed]; ok && r.App != "" {
			val = hed[r.Hed] + r.App + val
		}
		hed[r.Hed] = val
	}
	return hed
}


func (m Mapping) query(arg map[string]string) map[string]string {
	qry := map[string]string{}
	for _, r := range m {
		if r.Qry == "" {
			continue
		}
		if r.Equ != "" && arg[r.Arg] != r.Equ {
			continue
		}
		val := r.Val
		if val == "" {
			val = arg[r.Arg]
		}
		if _, ok := qry[r.Qry]; ok && r.App != "" {
			val = qry[r.Qry] + r.App + val
		}
		qry[r.Qry] = val
	}
	return qry
}



func (r Request) Fetch(arg map[string]string) []byte {
	des := r.Hos + "/" + r.Pat
	fmt.Println("    " + des)
	qry := url.Values{}
	fmt.Println("      query")
	for k, v := range r.Map.query(arg) {
		fmt.Printf("        %s=%s\n", k, v)
		qry.Add(k, v)
	}
	if str := qry.Encode(); len(str) > 0 {
		des += "?" + str
	}
	req, errReq := http.NewRequest("GET", des, nil)
	if errReq != nil {
		return nil
	}
	fmt.Println("      headers")
	for k, v := range r.Map.headers(arg) {
		fmt.Printf("        %s: %s\n", k, v)
		req.Header.Set(k, v)
	}
	res, errRes := http.DefaultClient.Do(req)
    	if errRes != nil {
        	return nil
    	}
	defer res.Body.Close()
	ret, errRet := ioutil.ReadAll(res.Body)
        if errRet != nil {
		return nil
	}
	return ret
}
