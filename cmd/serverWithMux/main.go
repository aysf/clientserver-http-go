package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/one", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helo one!"))
		w.Write([]byte("\n------\n"))

		// getting param query from url
		url := r.URL.Query()
		p1 := url.Get("param1")
		w.Write([]byte(p1))
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

	})
	mux.Handle("/two", http.HandlerFunc(Two))

	fmt.Println("Server listening on port 9080...")
	l, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatal("error net listen:", err)
	}

	http.Serve(l, mux)

}

func Two(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is two!, hello two!"))
}
