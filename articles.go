package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
)

var (
	articlesURL = ""
	green       = color.New(color.FgGreen).SprintFunc()
	red         = color.New(color.FgRed).SprintFunc()
)

func publishArticleByID(ids []string) {

	params := map[string]string{}

	articlesURL = urlPrefix + *environmentFlag + urlServiceAPI + urlSuffix + "articles"

	items := len(ids)

	fmt.Printf("Processing: %s items\n", green(items))

	for index, id := range ids {
		URL := articlesURL + "/" + id
		response, err := makePetition(http.MethodGet, URL, nil, tokenFlag, params)
		if err != nil {
			log.Fatalln(red(err))
		}

		attributes := response["attributes"].(map[string]interface{})

		defaultAttributes := map[string]interface{}{
			"when": "now",
		}

		if *keepPublishDate {
			defaultAttributes["when"] = attributes["publishedAt"].(string)
			defaultAttributes["isRepublish"] = true
		}

		dataPublish := map[string]interface{}{
			"data": map[string]interface{}{
				"type":       "flats",
				"attributes": defaultAttributes,
			},
		}

		dataPublishCasted, err := json.Marshal(dataPublish)
		if err != nil {
			log.Fatalln(red(err))
		}

		fmt.Printf("Publishing article (%s of %s) with id: %s", green(index+1), green(items), green(id))

		publishURL := articlesURL + "/" + id + "/publish"
		_, err = makePetition(http.MethodPost, publishURL, dataPublishCasted, tokenFlag, nil)
		if err != nil {
			log.Fatalln(err)
		}

		if attributes["category"] == nil {
			fmt.Printf(". This article doesn't have %s\n", red("category"))
			continue
		}

		category := attributes["category"].(string)
		seo := attributes["seo"].(map[string]interface{})

		if seo["slug"] == nil {
			fmt.Printf(". This article doesn't have %s\n", red("slug"))
			continue
		}
		slug := seo["slug"].(string)
		urlArticlePublished := urlPrefix + *environmentFlag + urlServiceHTML + urlSuffix + category + "/" + slug + ".html"
		fmt.Printf(", in: %s\n", green(urlArticlePublished))
	}

	os.Exit(0)
}

func publishArticles() {
	articlesURL = urlPrefix + *environmentFlag + urlServiceAPI + urlSuffix + "articles"

	params := map[string]string{
		"status": "STATUS_PUBLISHED",
		"limit":  *limitFlag,
		"page":   *pageFlag,
	}

	response, err := makePetition(http.MethodGet, articlesURL, nil, tokenFlag, params)
	if err != nil {
		log.Fatalln(red(err))
	}

	paginate := response["metadata"].(map[string]interface{})["paginate"].(map[string]interface{})

	actual, total, _ := getPagination(paginate)

	for actual <= total {
		params := map[string]string{
			"status": "STATUS_PUBLISHED",
			"limit":  *limitFlag,
			"page":   strconv.FormatInt(actual, 10),
		}

		response, err := makePetition(http.MethodGet, articlesURL, nil, tokenFlag, params)
		if err != nil {
			log.Fatalln(red(err))
		}

		paginate := response["metadata"].(map[string]interface{})["paginate"].(map[string]interface{})

		pageInt, pageCountInt, totalCountInt := getPagination(paginate)

		page := strconv.FormatInt(pageInt, 10)
		pageCount := strconv.FormatInt(pageCountInt, 10)
		totalCount := strconv.FormatInt(totalCountInt, 10)

		fmt.Printf("Processing: Page %s of %s with %s total items\n", green(page), green(pageCount), green(totalCount))

		data := response["data"].([]interface{})
		handleArticles(data, len(data))

		actual++
	}
}

func getPagination(paginate map[string]interface{}) (page, pageCount, totalCount int64) {
	page64 := paginate["page"].(float64)
	pageCount64 := paginate["pageCount"].(float64)
	totalCount64 := paginate["totalCount"].(float64)

	page = int64(page64)
	pageCount = int64(pageCount64)
	totalCount = int64(totalCount64)

	return
}

func handleArticles(data []interface{}, total int) {
	for index, articleRaw := range data {
		article := articleRaw.(map[string]interface{})

		attributes := article["attributes"].(map[string]interface{})

		articleID := article["id"].(string)

		articlesURLPublish := articlesURL + "/" + articleID + "/publish"

		defaultAttributes := map[string]interface{}{
			"when": "now",
		}

		if *keepPublishDate {
			defaultAttributes["when"] = attributes["publishedAt"].(string)
			defaultAttributes["isRepublish"] = true
		}

		dataPublish := map[string]interface{}{
			"data": map[string]interface{}{
				"type":       "flats",
				"attributes": defaultAttributes,
			},
		}

		dataPublishCasted, err := json.Marshal(dataPublish)
		if err != nil {
			log.Fatalln(red(err))
		}

		fmt.Printf("Publishing article (%s of %s) with id: %s", green(index+1), green(total), green(articleID))

		_, err = makePetition(http.MethodPost, articlesURLPublish, dataPublishCasted, tokenFlag, nil)
		if err != nil {
			log.Fatalln(err)
		}

		if attributes["category"] == nil {
			fmt.Printf(". This article doesn't have %s\n", red("category"))
			continue
		}

		category := attributes["category"].(string)
		seo := attributes["seo"].(map[string]interface{})

		if seo["slug"] == nil {
			fmt.Printf(". This article doesn't have %s\n", red("slug"))
			continue
		}

		slug := seo["slug"].(string)

		urlArticlePublished := urlPrefix + *environmentFlag + urlServiceHTML + urlSuffix + category + "/" + slug + ".html"

		fmt.Printf(", in: %s\n", green(urlArticlePublished))
	}
}
