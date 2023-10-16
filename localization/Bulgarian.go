package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Bulgarian] = map[string]string{
		"#allow-role":                   "разрешени-роли",
		"#look":                         "виж",
		"#move":                         "ход",
		"#review":                       "преглед",
		"#allow-role.value":             "стойност",
		"#.subject":                     "мишена",
		"#look.info":                    "информация",
		"#look.review-list":             "списък-за-преглед",
		"#allow-role.Description":       "Само администратор - По подразбиране: False",
		"#look.Description":             "виж нещо",
		"#move.Description":             "Преместете рецензията в този канал",
		"#review.Description":           "Потребителски преглед",
		"#allow-role.value.Description": "Задайте стойност",
		"#.subject.Description":         "Изберете цел",
		"#look.info.Description":        "Вижте потребителска информация",
		"#look.review-list.Description": "Вижте списъка с отзиви, получени от потребители",
		"#allow-role.NeedPermissions":   "На бота липсват разрешения - Управление на роли",
		"#allow-role.InProgress":        "работа в прогрес",
		"#allow-role.proc.Title":        "Опции за промяна",
		"#allow-role.proc.Description":  "Разрешение за роля",
		"#allow-role.proc.InProgress":   "Процедура",
		"#allow-role.proc.Done":         "пълен",
		"#allow-role.Keep":              "Няма променени настройки",
		"#look.info.IsNone":             "Няма отзиви",
		"#look.review-list.IsNone":      "Няма отзиви",
		"#look.review-list.menu.Title":  "Отзиви за %s",
		"#look.review-list.menu.Page":   "%d страница",
		"#move.IsNone":                  "Няма написани рецензии за целта",
		"#move.Move":                    "премести се тук",
		"#review.SelfReview":            "Не можете да прегледате себе си",
		"#review.modal.Title":           "Преглед на %s",
		"#review.lable.Score":           "резултат",
		"#review.lable.Title":           "заглавие",
		"#review.lable.Content":         "детайл",
		"$review.IsEdited":              "Тази рецензия е редактирана",
		"$review.NoAuthor":              "Изтритите рецензии не могат да бъдат възстановени, защото авторът не съществува тук",
		"$review.DM":                    "Написана е нова рецензия",
	}
}