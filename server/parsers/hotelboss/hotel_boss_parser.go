package hotelboss

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/vahanito/lunch-app-go/model"
	parsers "github.com/vahanito/lunch-app-go/parsers/utils"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const hotelBossUrl = "https://www.hotel-boss.eu/denne-menu.php"

func ParseHotelBoss() model.Menu {
	resp, _ := http.Get(hotelBossUrl)
	defer resp.Body.Close()

	document, _ := goquery.NewDocumentFromReader(resp.Body)
	todayMenu := document.Find(".contact-form-area").FilterFunction(func(i int, selection *goquery.Selection) bool {
		return strings.Contains(selection.Text(), parsers.TranslateDay(time.Now().Format("Monday")))
	})

	menu := model.Menu{}
	regex, _ := regexp.Compile("[ ]{2,}|[\r\n]+")
	menu.Restaurant = "Hotel Boss"
	menu.Url = hotelBossUrl

	menuTable := todayMenu.Find("table").Eq(int(time.Now().Weekday()) - 1)

	rows := menuTable.Find("tr").FilterFunction(func(i int, selection *goquery.Selection) bool {
		return len(strings.TrimSpace(selection.Text())) != 0
	})

	rows.Each(func(i int, selection *goquery.Selection) {
		processMenuItemElement(selection, i, &menu, regex)
	})
	return menu
}

func processMenuItemElement(selection *goquery.Selection, index int, menu *model.Menu, regex *regexp.Regexp) {
	if index < 2 {
		fillSoupInfo(menu, regex, selection.Find("td").Eq(1).Text())
	} else {
		elementText := selection.Find("td").Eq(1).Text() + " " + selection.Find("td").Eq(2).Text()
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
