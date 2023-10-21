package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Lithuanian] = map[string]string{
		"#admin":                              "admin",
		"#admin.allow-role":                   "leisti-vaidmenį",
		"#look":                               "žiūrėk",
		"#move":                               "judėti",
		"#review":                             "apžvalga",
		"#admin.allow-role.value":             "vertė",
		"#.subject":                           "tikslas",
		"#look.info":                          "informacija",
		"#look.reviews":                       "apžvalgos",
		"#admin.allow-role.Description":       "Rodyti balus kaip vaidmenis. Numatytoji: klaidinga",
		"#look.Description":                   "ką nors pamatyti",
		"#move.Description":                   "Perkelti apžvalgą į šį kanalą",
		"#review.Description":                 "Vartotojo apžvalga",
		"#admin.allow-role.value.Description": "nustatyta vertė",
		"#.subject.Description":               "Pasirinkite tikslą",
		"#look.info.Description":              "Peržiūrėkite vartotojo informaciją",
		"#look.reviews.Description":           "Peržiūrėkite atsiliepimus, kurių subjektas gavo",
		"#admin.allow-role.NeedPermissions":   "Botui trūksta leidimų – Tvarkyti vaidmenis",
		"#admin.allow-role.InProgress":        "darbas vyksta",
		"#admin.allow-role.proc.Title":        "Modifikuoti parinktis",
		"#admin.allow-role.proc.Description":  "Leisti vaidmenis",
		"#admin.allow-role.proc.InProgress":   "Vykdoma",
		"#admin.allow-role.proc.Done":         "užbaigti",
		"#admin.allow-role.Keep":              "Jokie nustatymai nepasikeitė",
		"#look.info.IsNone":                   "Nėra atsiliepimų",
		"#look.reviews.IsNone":                "Apžvalgų nėra",
		"#look.reviews.menu.Title":            "Atsiliepimai apie %v",
		"#look.reviews.menu.Page":             "%v puslapis",
		"#move.IsNone":                        "Nėra parašytų atsiliepimų apie taikinį",
		"#move.Move":                          "persikelti čia",
		"#review.SelfReview":                  "Jūs negalite peržiūrėti savęs",
		"#review.modal.Title":                 "Peržiūrėkite %v",
		"#review.lable.Score":                 "balas",
		"#review.lable.Title":                 "titulą",
		"#review.lable.Content":               "detalė",
		"$review.IsEdited":                    "Ši apžvalga buvo redaguota",
		"$review.NoAuthor":                    "Ištrintų atsiliepimų atkurti negalima, nes autoriaus čia nėra",
		"$review.DM":                          "Parašyta nauja apžvalga",
	}
}
