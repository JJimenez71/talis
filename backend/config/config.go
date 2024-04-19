package config


import (
	"io/ioutil"
	"strings"
)


func Load(pat string) map[string]string {
	fil := map[string]string{}
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
			panic("invalid key-value: \""+l+"\"")
		}
		key := strings.TrimSpace(l[:idx])
		val := strings.TrimSpace(l[idx+1:])
		if key == "" {
			panic("invalid key-value: missing key: \""+l+"\"")
		}
		if val == "" {
			panic("invalid key-value: missing value: \""+l+"\"")
		}
		fil[key] = val
	}
	return fil
}
