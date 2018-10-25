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
	keepPublishDate *bool
	startDateFlag   *string
	endDateFlag     *string
	typePostFlag    *string
	statusPostFlag  *string
)

func main() {
	v := flag.Bool("v", false, "Print the version of the program")
	version := flag.Bool("version", false, "Print the version of the program")

	environmentFlag = flag.String("environment", "dev", "Environment to make the petition {dev, staging}")
	limitFlag = flag.String("limit", "50", "Limit of items in the response")
	pageFlag = flag.String("page", "1", "Number of the page where start")
	tokenFlag = flag.String("token", "", "Token needed for make the petition")
	keepPublishDate = flag.Bool("keep-publish-date", false, "Flag to keep publish date")
	startDateFlag = flag.String("start-date", "2018-01-01", "Year to bring Article, Default: 2018-01-01")
	endDateFlag = flag.String("end-date", "2018-10-31", "Month to bring Articles. Default: 2018-12-31")
	typePostFlag = flag.String("type-post", "VIDEO", "Article type to be searched. Default: video")
	statusPostFlag = flag.String("status-post", "STATUS_PUBLISHED", "Article status to be searched. Default: published")

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
