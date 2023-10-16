package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Czech] = map[string]string{
		"#allow-role":                   "role-povolena",
		"#look":                         "koukni-se",
		"#move":                         "hýbat-se",
		"#review":                       "posouzení",
		"#allow-role.value":             "hodnota",
		"#.subject":                     "cílová",
		"#look.info":                    "informace",
		"#look.review-list":             "seznam-recenzí",
		"#allow-role.Description":       "Pouze správce – Výchozí: False",
		"#look.Description":             "něco vidět",
		"#move.Description":             "Přesunout recenzi na tento kanál",
		"#review.Description":           "Uživatelská recenze",
		"#allow-role.value.Description": "Nastavit hodnotu",
		"#.subject.Description":         "Vyberte cíl",
		"#look.info.Description":        "Zobrazit informace o uživateli",
		"#look.review-list.Description": "Zobrazit seznam recenzí přijatých uživateli",
		"#allow-role.NeedPermissions":   "Bot postrádá oprávnění – Správa rolí",
		"#allow-role.InProgress":        "probíhající práce",
		"#allow-role.proc.Title":        "Upravit možnosti",
		"#allow-role.proc.Description":  "Povolit roli",
		"#allow-role.proc.InProgress":   "Pokračování",
		"#allow-role.proc.Done":         "kompletní",
		"#allow-role.Keep":              "Žádná změna nastavení",
		"#look.info.IsNone":             "Žádné recenze",
		"#look.review-list.IsNone":      "Nejsou žádné recenze",
		"#look.review-list.menu.Title":  "Recenze pro %s",
		"#look.review-list.menu.Page":   "%d stránka",
		"#move.IsNone":                  "Pro cíl nejsou napsány žádné recenze",
		"#move.Move":                    "Pojď sem",
		"#review.SelfReview":            "Nemůžete hodnotit sami sebe",
		"#review.modal.Title":           "Zkontrolujte %s",
		"#review.lable.Score":           "skóre",
		"#review.lable.Title":           "titul",
		"#review.lable.Content":         "detail",
		"$review.IsEdited":              "Tato recenze byla upravena",
		"$review.NoAuthor":              "Smazané recenze nelze obnovit, protože zde autor neexistuje",
		"$review.DM":                    "Byla napsána nová recenze",
	}
}
