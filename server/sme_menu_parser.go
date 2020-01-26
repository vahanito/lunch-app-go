package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://restauracie.sme.sk/restauracia/del-patio_10778-zilina_2737/denne-menu")
	//bytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	//fmt.Println("HTML:\n\n", string(bytes))
	document, _ := goquery.NewDocumentFromReader(resp.Body)
	fmt.Println(document.Find(".emptyContentMessageContainer").Text())
}
