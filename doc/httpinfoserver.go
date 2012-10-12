//
//Show out http request information to client
//
package main

import (
	"fmt"
	"github.com/wheelcomplex/go-pkg-optargex"
	"net/http"
	"time"
)

var listenAddr = "0.0.0.0:8088"
var showhelp = false

const AdminEmail = "sa@xiaomi.com"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, optargex.Version()+"\n")
	fmt.Fprintf(w, "\nAdministrator: "+AdminEmail+"\n")
	fmt.Fprintf(w, "\nDate: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Fprintf(w, "\nlocal connection: %v\n", r.LocalAddr())
	fmt.Fprintf(w, "\nremote connection: %s\n", r.RemoteAddr)

	fmt.Fprintf(w, "\nrequest URI: %s http://%s%s %s\n\n", r.Method, r.Host, r.URL.Path[1:], r.Proto)

	for key, value := range r.Header {
		fmt.Fprintf(w, "request Header: %s : %s\n", key, value)
	}

}

func main() {
	optargex.Add("h", "help", "show this help.", false)
	optargex.Add("l", "listen", "Listen Address(format: ip-address:port)", "0.0.0.0:8088")
	// Parse os.Args
	for opt := range optargex.Parse() {
		//fmt.Printf("checking: %v, %v\n", opt.ShortName, opt.Name)
		switch opt.ShortName {
		case "h":
			showhelp = opt.Bool()
		case "l":
			listenAddr = opt.String()
		}
		switch opt.Name {
		case "help":
			showhelp = opt.Bool()
		case "listen":
			listenAddr = opt.String()
		}
	}
	if showhelp {
		optargex.Usage()
		return
	}

	fmt.Printf("HTTP Info Server v1.0.0 listening at %s", listenAddr)

	http.HandleFunc("/", handler)
	http.ListenAndServe(listenAddr, nil)
}
