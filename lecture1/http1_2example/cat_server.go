package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	port := os.Args[1]
	protocol := os.Args[2]

	const prefix = "TODO:YOUR_PREFIX"
	const delay = 0
	const pixelWidth = 800
	const size = 32

	fs := http.StripPrefix("/image/", http.FileServer(http.Dir(prefix)))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(rw, fmt.Sprintf("<body style='width:%dpx'>", pixelWidth)+
			GenerateHTML(pixelWidth, size)+"</body>")
	})

	http.HandleFunc("/image/", func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fs.ServeHTTP(rw, r)
	})

	if protocol == "http2" {
		fmt.Println("HTTP2 Server started at port " + port)
		log.Fatal(http.ListenAndServeTLS(":"+port, "server.crt", "server.key", nil))
	} else if protocol == "http1" {
		fmt.Println("HTTP1.1 Server started at port " + port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	} else {
		fmt.Print("specify protocol")
		os.Exit(2)
	}

}

func GenerateHTML(widthPx int, sizeCells int) string {
	wh := widthPx / sizeCells

	var sb strings.Builder

	for i := 1; i < sizeCells+1; i++ {
		for j := 1; j < sizeCells+1; j++ {
			sb.WriteString(
				fmt.Sprintf("<img src=\"/image/Cat_%02d_%02d.png\" width=%d height=%d>", i, j, wh, wh))
		}
	}

	return sb.String()
}
