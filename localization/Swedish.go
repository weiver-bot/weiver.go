package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Swedish] = map[string]string{
		"#allow-role":                   "roll-tillåten",
		"#look":                         "se",
		"#move":                         "flytta",
		"#review":                       "recension",
		"#allow-role.value":             "värde",
		"#.subject":                     "mål",
		"#look.info":                    "information",
		"#look.review-list":             "granskningslista",
		"#allow-role.Description":       "Endast administratör - Standard: False",
		"#look.Description":             "se något",
		"#move.Description":             "Flytta recensionen till den här kanalen",
		"#review.Description":           "Användarrecension",
		"#allow-role.value.Description": "satt värde",
		"#.subject.Description":         "Välj mål",
		"#look.info.Description":        "Visa användarinformation",
		"#look.review-list.Description": "Se listan över recensioner som användare tagit emot",
		"#allow-role.NeedPermissions":   "Bot saknar behörigheter - Hantera roller",
		"#allow-role.InProgress":        "pågående arbete",
		"#allow-role.proc.Title":        "Ändra alternativ",
		"#allow-role.proc.Description":  "Tillåt roll",
		"#allow-role.proc.InProgress":   "Förfarande",
		"#allow-role.proc.Done":         "komplett",
		"#allow-role.Keep":              "Inga inställningar har ändrats",
		"#look.info.IsNone":             "Inga recensioner",
		"#look.review-list.IsNone":      "Det finns inga recensioner",
		"#look.review-list.menu.Title":  "Recensioner för %s",
		"#look.review-list.menu.Page":   "%d sida",
		"#move.IsNone":                  "Det finns inga recensioner skrivna för målet",
		"#move.Move":                    "flytta hit",
		"#review.SelfReview":            "Du kan inte recensera dig själv",
		"#review.modal.Title":           "Granska %s",
		"#review.lable.Score":           "Göra",
		"#review.lable.Title":           "titel",
		"#review.lable.Content":         "detalj",
		"$review.IsEdited":              "Denna recension har redigerats",
		"$review.NoAuthor":              "Raderade recensioner kan inte återställas eftersom författaren inte finns här",
		"$review.DM":                    "En ny recension har skrivits",
	}
}
