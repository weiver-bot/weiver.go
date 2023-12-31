package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Danish] = map[string]string{
		"#admin":                              "admin",
		"#admin.allow-role":                   "tillade-rolle",
		"#look":                               "se",
		"#move":                               "bevæge-sig",
		"#review":                             "anmeldelse",
		"#admin.allow-role.value":             "værdi",
		"#.subject":                           "mål",
		"#look.info":                          "information",
		"#look.reviews":                       "anmeldelser",
		"#admin.allow-role.Description":       "Vis score som roller. Standard: Falsk",
		"#look.Description":                   "se noget",
		"#move.Description":                   "Flyt anmeldelse til denne kanal",
		"#review.Description":                 "Brugeranmeldelse",
		"#admin.allow-role.value.Description": "indstillet værdi",
		"#.subject.Description":               "Vælg mål",
		"#look.info.Description":              "Se brugeroplysninger",
		"#look.reviews.Description":           "Se de anmeldelser, emnet har modtaget",
		"#admin.allow-role.NeedPermissions":   "Bot mangler tilladelser - Administrer roller",
		"#admin.allow-role.InProgress":        "arbejde der er i gang",
		"#admin.allow-role.proc.Title":        "Rediger indstillinger",
		"#admin.allow-role.proc.Description":  "Tillad roller",
		"#admin.allow-role.proc.InProgress":   "Fortsætter",
		"#admin.allow-role.proc.Done":         "komplet",
		"#admin.allow-role.Keep":              "Ingen indstillinger ændret",
		"#look.info.IsNone":                   "Ingen anmeldelser",
		"#look.reviews.IsNone":                "Der er ingen anmeldelser",
		"#look.reviews.menu.Title":            "Anmeldelser for %v",
		"#look.reviews.menu.Page":             "%v side",
		"#move.IsNone":                        "Der er ingen anmeldelser skrevet til målet",
		"#move.Move":                          "flytte hertil",
		"#review.SelfReview":                  "Du kan ikke anmelde dig selv",
		"#review.modal.Title":                 "Anmeldelser",
		"#review.lable.Score":                 "score",
		"#review.lable.Title":                 "titel",
		"#review.lable.Content":               "detalje",
		"$review.IsEdited":                    "Denne anmeldelse er blevet redigeret",
		"$review.NoAuthor":                    "Slettede anmeldelser kan ikke gendannes, fordi forfatteren ikke findes her",
		"$review.DM":                          "Der er skrevet en ny anmeldelse",
	}
}
