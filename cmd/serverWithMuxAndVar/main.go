package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/one", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this is one")
	})
	mux.Handle("/cekquery", http.HandlerFunc(cekQuery))
	mux.Handle("/viewquery", http.HandlerFunc(viewquery))
	mux.Handle("/user/", http.HandlerFunc(user))

	fmt.Println("Server listening on port 9080...")
	l, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatal("error net listen:", err)
	}

	http.Serve(l, mux)

}

func user(w http.ResponseWriter, r *http.Request) {
	s := r.URL.RequestURI()
	log.Println("s", s)

	e := strings.Split(s, "/")
	log.Println("e:", e)

	q := strings.Split(e[2], "?")
	log.Println("q:", q)

	fmt.Println("this is user id:", q[0])
}

func viewquery(w http.ResponseWriter, r *http.Request) {

	b := r.URL.Query()

	fmt.Println(b)
	if b["param1"] != nil {
		p1 := b["param1"]
		fmt.Println(p1[0])
	}

}

func cekQuery(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is two!, hello two!"))

	w.Write([]byte("helo one!"))
	w.Write([]byte("\n------\n"))

	// getting param query from url
	url := r.URL.Query()
	p1 := url.Get("param1")
	w.Write([]byte("param1:" + p1))
	w.Write([]byte("\n------\n"))

	p2 := url.Get("param2")
	w.Write([]byte("param2:" + p2))
	fmt.Fprintf(w, "param type: %T", p2)
	w.Write([]byte("\n------\n"))

	host := r.URL.Host
	w.Write([]byte("host:" + host))
	w.Write([]byte("\n------\n"))

	p := r.URL.Port()
	w.Write([]byte("port:" + p))
	w.Write([]byte("\n------\n"))

	uri := r.URL.RequestURI()
	w.Write([]byte("uri:" + uri))
	w.Write([]byte("\n------\n"))

	hostname := r.URL.Hostname()
	w.Write([]byte("hostname:" + hostname))
	w.Write([]byte("\n------\n"))

	op := r.URL.Opaque
	w.Write([]byte("Opaque:" + op))
	w.Write([]byte("\n------\n"))

	rawPath := r.URL.RawPath
	w.Write([]byte("rawPath:" + rawPath))
	w.Write([]byte("\n------\n"))

	rawQuery := r.URL.RawQuery
	w.Write([]byte("rawQuery:" + rawQuery))
	w.Write([]byte("\n------\n"))

	rawFragmen := r.URL.RawFragment
	w.Write([]byte("rawFragmen:" + rawFragmen))
	w.Write([]byte("\n------\n"))

	remoteAddr := r.RemoteAddr
	w.Write([]byte("remoteAddr:" + remoteAddr))
	w.Write([]byte("\n------\n"))

	method := r.Method
	w.Write([]byte("method:" + method))
	w.Write([]byte("\n------\n"))

	proto := r.Proto
	w.Write([]byte("proto:" + proto))
	w.Write([]byte("\n------\n"))

	headerTest := r.Header.Get("test")
	w.Write([]byte("headerTest:" + headerTest))
	w.Write([]byte("\n------\n"))

	headerVal := r.Header.Values("test")
	for _, hv := range headerVal {
		w.Write([]byte("headerVal:" + hv))
	}
	w.Write([]byte("\n------\n"))

	fq := r.URL.ForceQuery
	if fq {
		w.Write([]byte("fq: true"))
	} else {
		w.Write([]byte("fq: false"))
	}
	w.Write([]byte("\n------\n"))

	fr := r.URL.Fragment
	w.Write([]byte("fr:" + fr))
	w.Write([]byte("\n------\n"))

}
