package parser

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func GetForm(id string, body io.Reader) map[string]string {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	var m = make(map[string]string)
	doc.Find(id + " input").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		m[name] = value
	})
	return m
}

func GetCSRFToken(body io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	token, _ := doc.Find("meta[name=csrf-token]").Attr("content")
	return token
}
