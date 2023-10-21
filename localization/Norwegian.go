package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Norwegian] = map[string]string{
		"#admin":                              "admin",
		"#admin.allow-role":                   "tillate-rolle",
		"#look":                               "se",
		"#move":                               "bevege-seg",
		"#review":                             "anmeldelse",
		"#admin.allow-role.value":             "verdi",
		"#.subject":                           "mål",
		"#look.info":                          "informasjon",
		"#look.reviews":                       "anmeldelser",
		"#admin.allow-role.Description":       "Vis poeng som roller. Standard: False",
		"#look.Description":                   "se noe",
		"#move.Description":                   "Flytt anmeldelse til denne kanalen",
		"#review.Description":                 "Brukeranmeldelse",
		"#admin.allow-role.value.Description": "Sett verdi",
		"#.subject.Description":               "Velg mål",
		"#look.info.Description":              "Se brukerinformasjon",
		"#look.reviews.Description":           "Se anmeldelsene emnet har fått",
		"#admin.allow-role.NeedPermissions":   "Bot mangler tillatelser - Administrer roller",
		"#admin.allow-role.InProgress":        "arbeid pågår",
		"#admin.allow-role.proc.Title":        "Endre alternativer",
		"#admin.allow-role.proc.Description":  "Tillat roller",
		"#admin.allow-role.proc.InProgress":   "Går videre",
		"#admin.allow-role.proc.Done":         "fullstendig",
		"#admin.allow-role.Keep":              "Ingen innstillinger endret",
		"#look.info.IsNone":                   "Ingen anmeldelser",
		"#look.reviews.IsNone":                "Det er ingen anmeldelser",
		"#look.reviews.menu.Title":            "Anmeldelser for %v",
		"#look.reviews.menu.Page":             "%v side",
		"#move.IsNone":                        "Det er ingen anmeldelser skrevet for målet",
		"#move.Move":                          "Flytt her",
		"#review.SelfReview":                  "Du kan ikke anmelde deg selv",
		"#review.modal.Title":                 "Gjennomgå %v",
		"#review.lable.Score":                 "score",
		"#review.lable.Title":                 "tittel",
		"#review.lable.Content":               "detalj",
		"$review.IsEdited":                    "Denne anmeldelsen er redigert",
		"$review.NoAuthor":                    "Slettede anmeldelser kan ikke gjenopprettes fordi forfatteren ikke eksisterer her",
		"$review.DM":                          "En ny anmeldelse er skrevet",
	}
}
