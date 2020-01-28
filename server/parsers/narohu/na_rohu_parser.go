package narohu

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/vahanito/lunch-app-go/model"
	"net/http"
	"regexp"
)

const naRohuUrl = "https://restauracie.sme.sk/restauracia/narohu_9360-zilina_2737/denne-menu"

func ParseNaRohu() model.Menu {
	resp, _ := http.Get(naRohuUrl)
	defer resp.Body.Close()

	document, _ := goquery.NewDocumentFromReader(resp.Body)
	menuElement := document.Find(".dnesne_menu")

	menu := model.Menu{}
	regex, _ := regexp.Compile("[ ]{2,}|[\r\n]+")
	menu.Date = regex.ReplaceAllString(menuElement.Find("h2").Text(), " ")
	menu.Restaurant = "Na rohu"
	menu.Url = naRohuUrl

	menuItemDivs := menuElement.Find(".jedlo_polozka")

	if menuItemDivs.Size() == 1 {
		menu.Info = regex.ReplaceAllString(menuItemDivs.Text(), " ")
	}

	menuItemDivs.Each(func(i int, selection *goquery.Selection) {
		processMenuItemElement(selection, i, &menu, regex)
	})

	return menu
}

func processMenuItemElement(selection *goquery.Selection, index int, menu *model.Menu, regex *regexp.Regexp) {
	elementText := selection.Text()
	if index < 2 {
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
