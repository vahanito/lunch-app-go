package pivarium

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/vahanito/lunch-app-go/model"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const url = "https://restauracie.sme.sk/restauracia/plzenska-restauracia-pivarium_6977-zilina_2737/denne-menu"

func ParsePivarium() model.Menu {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	document, _ := goquery.NewDocumentFromReader(resp.Body)
	menuElement := document.Find(".dnesne_menu")

	menu := model.Menu{}
	regex, _ := regexp.Compile("[ ]{2,}|[\r\n]+")
	menu.Date = regex.ReplaceAllString(menuElement.Find("h2").Text(), " ")
	menu.Restaurant = "Pivarium"
	menu.Url = url

	menuItemDivs := menuElement.Find(".jedlo_polozka")

	if menuItemDivs.Size() == 1 {
		menu.Info = regex.ReplaceAllString(menuItemDivs.Text(), " ")
	}

	isSoup := false
	continueLoop := true
	menuItemDivs.Each(func(i int, selection *goquery.Selection) {
		if strings.Contains(selection.Text(), "Pri obj") {
			continueLoop = false
		}
		if continueLoop {
			processMenuItemElement(selection, &isSoup, &menu, regex)
		}
	})

	return menu
}

func processMenuItemElement(selection *goquery.Selection, isSoup *bool, menu *model.Menu, regex *regexp.Regexp) {
	elementText := selection.Text()
	formattedNow := time.Now().Format("02.01.2006")

	if strings.Contains(elementText, formattedNow) {
		*isSoup = true
		return
	} else if strings.Contains(elementText, "MENU") {
		*isSoup = false
	} else if *isSoup {
		fillSoupInfo(menu, regex, elementText)
		return
	}

	if strings.Contains(elementText, "MENU") {
		menu.Menus = append(menu.Menus, model.MenuItem{Name: regex.ReplaceAllString(elementText, " ")})
	} else {
		lastMenu := &menu.Menus[len(menu.Menus)-1]
		lastMenu.Name = lastMenu.Name + regex.ReplaceAllString(elementText, " ")
	}
}

func fillSoupInfo(menu *model.Menu, regex *regexp.Regexp, elementText string) {
	if len(menu.Soup.Name) == 0 {
		menu.Soup.Name += regex.ReplaceAllString(elementText, " ")
	} else {
		menu.Soup.Name += " / " + regex.ReplaceAllString(elementText, " ")
	}
}
