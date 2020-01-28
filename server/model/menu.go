package model

type Menu struct {
	Info         string      `json:"info,omitempty"`
	Restaurant   string      `json:"restaurant"`
	Url          string      `json:"url"`
	Soup         MenuItem    `json:"soup,omitempty"`
	Menus        []MenuItem  `json:"menus,omitempty"`
	SpecialMenus *[]MenuItem `json:"specialMenus,omitempty"`
	Date         string      `json:"date,omitempty"`
}
