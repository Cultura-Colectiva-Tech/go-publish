package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	urlPrefix      = "https://"
	urlServiceAPI  = ".api"
	urlServiceHTML = ".html"
	urlSuffix      = ".culturacolectiva.com/"
)

var (
	environmentFlag *string
	limitFlag       *string
	pageFlag        *string
	tokenFlag       *string
)

func main() {
	v := flag.Bool("v", false, "Print the version of the program")
	version := flag.Bool("version", false, "Print the version of the program")

	environmentFlag = flag.String("environment", "dev", "Environment to make the petition {dev, staging}")
	limitFlag = flag.String("limit", "50", "Limit of items in the response")
	pageFlag = flag.String("page", "1", "Number of the page where start")
	tokenFlag = flag.String("token", "", "Token needed for make the petition")

	flag.Parse()

	if *v || *version {
		fmt.Printf("go-publish version %s\n", appVersion)
		os.Exit(0)
	}

	if *tokenFlag == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	publishArticles()
}
