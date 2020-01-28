package drhunger

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/vahanito/lunch-app-go/model"
	parsers "github.com/vahanito/lunch-app-go/parsers/utils"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const drHungerUrl = "https://www.drhunger.sk"

func ParseDrHunger() model.Menu {
	resp, _ := http.Get(drHungerUrl)
	defer resp.Body.Close()

	document, _ := goquery.NewDocumentFromReader(resp.Body)
	menuElement := document.Find(".mk-accordion-single").FilterFunction(func(i int, selection *goquery.Selection) bool {
		return strings.Contains(selection.Text(), parsers.TranslateDay(time.Now().Format("Monday")))
	})

	menu := model.Menu{}
	regex, _ := regexp.Compile("[ ]{2,}|[\r\n]+")
	menu.Restaurant = "Dr. Hunger"
	menu.Url = drHungerUrl

	menuItemTds := menuElement.Find("td")

	menuItemTds.Each(func(i int, selection *goquery.Selection) {
		processMenuItemElement(selection, i, &menu, regex)
	})

	return menu
}

func processMenuItemElement(selection *goquery.Selection, index int, menu *model.Menu, regex *regexp.Regexp) {
	elementText := selection.Text()
	if index < 1 {
		menu.Soup.Name += regex.ReplaceAllString(elementText, " ")
	} else {
		menu.Menus = append(menu.Menus, model.MenuItem{Name: regex.ReplaceAllString(elementText, " ")})
	}
}
