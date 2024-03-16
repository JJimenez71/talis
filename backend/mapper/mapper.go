package mapper


import (
)

type Rule struct {
	ArgKey string
	ArgVal string
	QryKey string
	QryVal string
	App    string
}

func validRule(rul Rule) bool {
	// { "" , "" , param , hardcoded , "" }
	// set param value to hardcoded
	if rul.ArgKey == "" && rul.ArgVal == "" &&
		rul.QryKey != "" && rul.QryVal != "" &&rul.App == "" {
		return true
	}
	// { "" , "" , param , hardcoded , separator }
	// append hardcoded to param value with separator between
	if rul.ArgKey == "" && rul.ArgVal == "" &&
		rul.QryKey != "" && rul.QryVal != "" &&rul.App != "" {
		return true
	}
	// { arg, "" , param , "" , ""}
	// set param value to arg value
	if rul.ArgKey != "" && rul.ArgVal == "" &&
		rul.QryKey != "" && rul.QryVal == "" &&rul.App == "" {
		return true
	}
	// { arg, "" , param , "", separator }
	// append arg value to param value with separator between
	if rul.ArgKey != "" && rul.ArgVal == "" &&
		rul.QryKey != "" && rul.QryVal == "" &&rul.App != "" {
		return true
	}
	// { arg, target , param , hardcoded, "" }
	// set param value to hardcoded if arg value matches target value
	if rul.ArgKey != "" && rul.ArgVal != "" &&
		rul.QryKey != "" && rul.QryVal != "" &&rul.App == "" {
		return true
	}
	// { arg, target , param , hardcoded, separator }
	// append hardcoded to param value with separator between if
	// arg value matches target value
	if rul.ArgKey != "" && rul.ArgVal != "" &&
		rul.QryKey != "" && rul.QryVal != "" &&rul.App != "" {
		return true
	}
	return false
}

func FormQuery(rul []Rule, arg map[string]string) map[string]string {
	qry := map[string]string{}
	for _, r := range rul {
		if !validRule(r) {
			panic("invalid mapper Rule")
		}
		if r.ArgVal != "" && arg[r.ArgKey] != r.ArgVal {
			continue
		}
                val := r.QryVal
		if val == "" {
			val = arg[r.ArgKey]
		}
		if _, ok := qry[r.QryKey]; ok && r.App != "" {
			val = qry[r.QryKey] + r.App + val
		}
		qry[r.QryKey] = val
	}
	return qry
}
