package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func isWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	//colorWhite := "\033[37m"
	if isWindows() {
		colorRed = ""
		colorGreen = ""
		colorReset = ""
		//colorWhite = ""
	}

	tempDir, err := os.Getwd()
	if err != nil {
		fmt.Printf(string(colorRed))
		fmt.Println(err)
		fmt.Printf(string(colorReset))
	}
	var port = flag.String("port", "9000", "TCP port")
	var dir = flag.String("dir", tempDir, "The Directory to serve")
	var help = flag.Bool("h", false, "Help Menu")
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	// create file server handler
	fs := http.FileServer(http.Dir(*dir))
	fmt.Println("Serving Dir:" + string(colorGreen) + *dir + string(colorReset))
	fmt.Println("The Http server is running on port TCP/" + string(colorGreen) + *port)
	// start HTTP server with `fs` as the default handler
	log.Fatal(http.ListenAndServe(":"+*port, fs))

}
