package config


import (
	"io/ioutil"
	"strings"
)


type File map[string]string


func Load(pat string) File {
	fil := File{}
	txt, err := ioutil.ReadFile(pat)
	if err != nil {
		panic("cannot read: \""+pat+"\"")
	}
	lns := strings.Split(string(txt), "\n")
	for _, l := range lns {
		if l == "" {
			continue
		}
		idx := strings.Index(l, "=")
		if idx == -1 {
			panic("invalid key value line: \""+l+"\"")
		}
		key := strings.TrimSpace(l[:idx])
		val := strings.TrimSpace(l[idx+1:])
		if key == "" {
			panic("invalid key value line: missing key: \""+l+"\"")
		}
		if val == "" {
			panic("invalid key value line: missing value: \""+l+"\"")
		}
		fil[key] = val
	}
	return fil
}


func (f File) Value(key string) string {
	if val, ok := f[key]; ok {
		return val
	}
	return ""
}
