package parsers

func TranslateDay(day string) string {
	switch day {
	case "Monday":
		return "Pondelok"
	case "Tuesday":
		return "Utorok"
	case "Wednesday":
		return "Streda"
	case "Thursday":
		return "Štvrtok"
	default:
		return "Piatok"
	}
}
