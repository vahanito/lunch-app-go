package delpatio

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/vahanito/lunch-app-go/model"
	"net/http"
	"regexp"
	"strings"
)

const delPatioUrl = "https://restauracie.sme.sk/restauracia/del-patio_10778-zilina_2737/denne-menu"

func ParseDelPatio() model.Menu {
	resp, _ := http.Get(delPatioUrl)
	defer resp.Body.Close()

	document, _ := goquery.NewDocumentFromReader(resp.Body)
	menuElement := document.Find(".dnesne_menu")

	menu := model.Menu{}
	regex, _ := regexp.Compile("[ ]{2,}|[\r\n]+")
	menu.Date = regex.ReplaceAllString(menuElement.Find("h2").Text(), " ")
	menu.Restaurant = "Del patio"
	menu.Url = delPatioUrl

	menuItemDivs := menuElement.Find(".jedlo_polozka")

	if menuItemDivs.Size() == 1 {
		menu.Info = regex.ReplaceAllString(menuItemDivs.Text(), " ")
	}

	isSoup := false
	menuItemDivs.Each(func(i int, selection *goquery.Selection) {
		processMenuItemElement(selection, &isSoup, &menu, regex)
	})

	return menu
}

func processMenuItemElement(selection *goquery.Selection, isSoup *bool, menu *model.Menu, regex *regexp.Regexp) {
	elementText := selection.Text()
	if strings.Contains(elementText, "Polievka") {
		*isSoup = true
	} else if strings.Contains(elementText, "Hlavné jedlo") || strings.Contains(elementText, "Špeciál") {
		*isSoup = false
	} else if *isSoup {
		fillSoupInfo(menu, regex, elementText)
	} else {
		menu.Menus = append(menu.Menus, model.MenuItem{Name: regex.ReplaceAllString(elementText, " ")})
	}
}

func fillSoupInfo(menu *model.Menu, regex *regexp.Regexp, elementText string) {
	if len(menu.Soup.Name) == 0 {
		menu.Soup.Name += regex.ReplaceAllString(elementText, " ")
	} else {
		menu.Soup.Name += " / " + regex.ReplaceAllString(elementText, " ")
	}
}
