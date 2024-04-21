package vendors

import (
	"backend/getter"
)


type Foo struct {
	Get getter.Request
	Ber string
}


func FooAPI(cfg map[string]string) Foo {
	return Foo {
		Get: getter.Request {
			Hos: "foo.com",
			Pat: "bar",
			Map: getter.Mapping{
				{Hed: "Authorization", Val: "Bearer"},
				{Arg: "token", Hed: "Authorization", App: " "},
				{Hed: "Content-Type", Val: "Application/json"},
				{Qry: "search", Val: "Pubs And Bars"},
				{Arg: "latitude", Qry: "coords"},
				{Arg: "longitude", Qry: "coords", App: ","},
				{Qry: "range", Val: "320"},
				{Arg: "distance", Equ: "0", Qry: "range", Val: "5"},
				{Arg: "distance", Equ: "1", Qry: "range", Val: "20"},
				{Arg: "distance", Equ: "2", Qry: "range", Val: "80"},
				{Qry: "minCost", Val: "$$$$"},
				{Qry: "maxCost", Val: "$$$$"},
				{Arg: "expense", Equ: "0", Qry: "minCost", Val: "$"},
				{Arg: "expense", Equ: "0", Qry: "maxCost", Val: "$$"},
				{Arg: "expense", Equ: "1", Qry: "minCost", Val: "$$"},
				{Arg: "expense", Equ: "1", Qry: "maxCost", Val: "$$$"},
				{Arg: "expense", Equ: "2", Qry: "minCost", Val: "$$$"},
				{Arg: "expense", Equ: "2", Qry: "maxCost", Val: "$$$$"},
			},
		},
		Ber: cfg["FOO_API_KEY"],
        }
}


func (f Foo) Fetch(arg map[string]string) []map[string]string {
	arg["token"] = f.Ber
	f.Get.Fetch(arg)
	return []map[string]string{
	}
}

