package localization

import "github.com/bwmarrin/discordgo"

var data = map[discordgo.Locale]map[string]string{}

func Load(locale discordgo.Locale, value string) string {
	if container, ok := data[locale]; ok {
		if result, ok := container[value]; ok {
			return result
		}
	}
	return value
}

func LoadList(value string) *map[discordgo.Locale]string {
	var result = map[discordgo.Locale]string{}

	for locale, container := range data {
		if ele, ok := container[value]; ok {
			result[locale] = ele
		}
	}

	return &result
}
