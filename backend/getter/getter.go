package getter


// Below are valid uses of the Rule struct:
// ----------------------------------------
// { "" , "" , param_name , hardcoded_value , "" }
//     set param value to hardcoded
// { "" , "" , param_name , hardcoded_value , separator_value }
//     append hardcoded to param value with separator between
// { arg_name, "" , param_name , "" , ""}
//     set param value to arg value
// { arg_name, "" , param_name , "", separator_value }
//     append arg value to param value with separator between
// { arg_name, target_value , param_name , hardcoded_value, "" }
//     set param value to hardcoded if arg value matches target value
// { arg_name, target_value , param_name , hardcoded_value, separator_value }
//     append hardcoded to param value with separator between if
//     arg value matches target value
type Rule struct {
	ArgKey string
	ArgVal string
	QryKey string
	QryVal string
	App    string
}


type Rules []Rule


type Request struct {
	Hos string
	Pat string
	Rul Rules
}


func (r Rules) parameters(arg map[string]string) map[string]string {
	qry := map[string]string{}
	for _, i := range r {
		if i.QryKey == "" {
			panic("invalid rule: missing param key: "+
			"{"+i.ArgKey+","+i.ArgVal+","+i.QryKey+","+i.QryVal+","+i.App+"}")
		}
		if (i.ArgKey == "" && i.QryVal == "")  {
			panic("invalid rule: missing param value: "+
			"{"+i.ArgKey+","+i.ArgVal+","+i.QryKey+","+i.QryVal+","+i.App+"}")
		}
		if i.ArgVal != "" && arg[i.ArgKey] != i.ArgVal {
			continue
		}
		val := i.QryVal
		if val == "" {
			val = arg[i.ArgKey]
		}
		if _, ok := qry[i.QryKey]; ok && i.App != "" {
			val = qry[i.QryKey] + i.App + val
		}
		qry[i.QryKey] = val
	}
	return qry
}


func (r Request) Fetch(arg map[string]string) string {
	req := r.Hos + "/" + r.Pat + "?"
	for key, val := range r.Rul.parameters(arg) {
                req += key + "=" + val + "&"
        }
	return req[:len(req)-1] + "\n"
}



