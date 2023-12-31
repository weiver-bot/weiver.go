package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Dutch] = map[string]string{
		"#admin":                              "beheerder",
		"#admin.allow-role":                   "rol-toestaan",
		"#look":                               "look",
		"#move":                               "beweging",
		"#review":                             "beoordeling",
		"#admin.allow-role.value":             "waarde",
		"#.subject":                           "doel",
		"#look.info":                          "informatie",
		"#look.reviews":                       "beoordelingen",
		"#admin.allow-role.Description":       "Geef scores weer als rollen. Standaard: Onwaar",
		"#look.Description":                   "Iets zien",
		"#move.Description":                   "Verplaats de recensie naar dit kanaal",
		"#review.Description":                 "Gebruikersrecensie",
		"#admin.allow-role.value.Description": "ingestelde waarde",
		"#.subject.Description":               "Selecteer doel",
		"#look.info.Description":              "Bekijk gebruikersinformatie",
		"#look.reviews.Description":           "Bekijk de beoordelingen die het onderwerp heeft ontvangen",
		"#admin.allow-role.NeedPermissions":   "Bot heeft geen rechten - Beheer rollen",
		"#admin.allow-role.InProgress":        "lopende werkzaamheden",
		"#admin.allow-role.proc.Title":        "Opties wijzigen",
		"#admin.allow-role.proc.Description":  "Rollen toestaan",
		"#admin.allow-role.proc.InProgress":   "Doorgaan",
		"#admin.allow-role.proc.Done":         "compleet",
		"#admin.allow-role.Keep":              "Er zijn geen instellingen gewijzigd",
		"#look.info.IsNone":                   "Geen beoordelingen",
		"#look.reviews.IsNone":                "Er zijn geen beoordelingen",
		"#look.reviews.menu.Title":            "Beoordelingen voor %v",
		"#look.reviews.menu.Page":             "%v pagina",
		"#move.IsNone":                        "Er zijn geen beoordelingen geschreven voor het doel",
		"#move.Move":                          "verhuis naar hier",
		"#review.SelfReview":                  "Je kunt jezelf niet beoordelen",
		"#review.modal.Title":                 "Beoordeel %v",
		"#review.lable.Score":                 "scoren",
		"#review.lable.Title":                 "titel",
		"#review.lable.Content":               "detail",
		"$review.IsEdited":                    "Deze recensie is aangepast",
		"$review.NoAuthor":                    "Verwijderde recensies kunnen niet worden hersteld omdat de auteur hier niet bestaat",
		"$review.DM":                          "Er is een nieuwe recensie geschreven",
	}
}
