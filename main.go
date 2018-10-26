package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
	currentAPI      string
	publishIdsFlag  *string
)

func main() {
	v := flag.Bool("v", false, "Print the version of the program")
	version := flag.Bool("version", false, "Print the version of the program")

	environmentFlag = flag.String("environment", "dev", "Environment to make the petition {dev, staging, prod}")
	limitFlag = flag.String("limit", "50", "Limit of items in the response")
	pageFlag = flag.String("page", "1", "Number of the page where start")
	tokenFlag = flag.String("token", "", "Token needed for make the petition")
	keepPublishDate = flag.Bool("keep-publish-date", false, "Flag to keep publish date")
	publishIdsFlag = flag.String("publish-ids", "", "Publish by article id")

	configEnvs := map[string]string{
		"dev":     "dev.api",
		"prod":    "api-v2",
		"staging": "staging.api",
	}

	flag.Parse()

	currentAPI = configEnvs[*environmentFlag]

	if *v || *version {
		fmt.Printf("go-publish version %s\n", appVersion)
		os.Exit(0)
	}

	if *publishIdsFlag != "" {
		articlesID := strings.Split(*publishIdsFlag, ",")
		if len(articlesID) > 0 {
			publishArticleByID(articlesID)
		}
	}

	if *tokenFlag == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	publishArticles()
}
