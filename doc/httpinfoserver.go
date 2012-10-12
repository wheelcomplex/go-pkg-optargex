//
// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/ and
// http://creativecommons.org/publicdomain/zero/1.0/legalcode
//
//thanks https://github.com/jteeuwen/go-pkg-optarg
//
//https://github.com/wheelcomplex/go-pkg-optargex/blob/master/doc/httpinfoserver.go
//
//create by <Wheelcomplex Yin> wheelcomplex@gmail.com
//
//Show out http request information to client
//
package main

import (
	"fmt"
	"github.com/wheelcomplex/go-pkg-optargex"
	"net/http"
	"strings"
	"time"
)

var listenAddr = "0.0.0.0:8088"
var showhelp = false
var showversion = false

const AdminEmail = "wheelcomplex@gmail.com"

func handler(w http.ResponseWriter, r *http.Request) {

	var showurl string

	fmt.Fprintf(w, optargex.VersionString()+"\n")
	fmt.Fprintf(w, "\nAdministrator: "+AdminEmail+"\n")
	fmt.Fprintf(w, "\nDate: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Fprintf(w, "\nlocal connection: %v\n", r.LocalAddr())
	fmt.Fprintf(w, "\nremote connection: %s\n", r.RemoteAddr)

	switch {
	case r.RequestURI == "":
		showurl = "/"
	case r.RequestURI[:1] != "/":
		showurl = "/" + r.RequestURI
	default:
		showurl = r.RequestURI
	}

	fmt.Fprintf(w, "\nrequest URI: %s http://%s%s %s\n\n", r.Method, r.Host, showurl, r.Proto)

	for key, value := range r.Header {
		fmt.Fprintf(w, "request Header: %s : %s\n", key, value)
	}

}

func main() {
	optargex.SetVersion("HTTP Info Server v1.0.0")

	optargex.Add("v", "version", "show version.", false)
	optargex.Add("h", "help", "show this help.", false)
	optargex.Add("l", "listen", "Listen Address(format: ip-address:port)", "0.0.0.0:8088")
	// Parse os.Args
	for opt := range optargex.Parse() {
		//fmt.Printf("checking: %v, %v\n", opt.ShortName, opt.Name)
		switch opt.ShortName {
		case "v":
			showversion = opt.Bool()
		case "h":
			showhelp = opt.Bool()
		case "l":
			listenAddr = opt.String()
		}
		switch opt.Name {
		case "version":
			showversion = opt.Bool()
		case "help":
			showhelp = opt.Bool()
		case "listen":
			listenAddr = opt.String()
		}
	}
	if showversion {
		optargex.Version()
		return
	}
	if showhelp {
		optargex.Usage()
		return
	}
	if strings.Contains(listenAddr, ".") {
		if !strings.Contains(listenAddr, ":") {
			listenAddr = listenAddr + ":8088"
		}

	} else {
		if !strings.Contains(listenAddr, ":") {
			listenAddr = ":" + listenAddr
		}
	}
	fmt.Printf(optargex.VersionString())
	fmt.Printf("Listening at %s ...\n", listenAddr)

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		fmt.Printf("ERROR: ListenAndServe: %v\n", err)
	} else {
		fmt.Printf("Listening at %s exited.\n", listenAddr)
	}
}
