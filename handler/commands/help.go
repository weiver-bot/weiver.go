package slashcommands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	g "github.com/y2hO0ol23/weiver/handler"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func init() {
	g.CMDList = append(g.CMDList, g.CMDForm{
		Data: &discordgo.ApplicationCommand{
			Name:              "?",
			NameLocalizations: &map[discordgo.Locale]string{"": "?"},
		},
		Message: func(s *discordgo.Session, i *discordgo.InteractionCreate, _ discordgo.Locale, queries []string) string {
			if len(queries) != 0 {
				for locale, _ := range discordgo.Locales {
					if !strings.EqualFold(queries[0], string(locale)) {
						continue
					}

					command := ""
					description := ""

					for _, e := range g.CMDList {
						if e.Message == nil {
							continue
						}
						if cmd, ok := (*e.Data.NameLocalizations)[locale]; ok {
							shorten := ""
							if len(e.Data.Options) != 0 {
								shorten = "..."
							}
							command += fmt.Sprintf("**/%v** %v\n", cmd, shorten)

							if e.Data.DescriptionLocalizations != nil {
								description += fmt.Sprintf("`# %v`\n", (*e.Data.DescriptionLocalizations)[locale])
							} else {
								description += "\n"
							}
						}
					}

					embed := builder.Embed().
						SetTitle("commands:").
						AddFields(
							&discordgo.MessageEmbedField{
								Value:  command,
								Inline: true,
							},
							&discordgo.MessageEmbedField{
								Value:  description,
								Inline: true,
							},
						)

					s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{
							embed.MessageEmbed,
						},
						Flags: discordgo.MessageFlagsEphemeral,
					}))
					return ""
				}
			}

			send := "```\n/? <locale>\nlocale:\n"
			for locale, name := range discordgo.Locales {
				if v := strings.ToLower(string(locale)); v != "" {
					send += fmt.Sprintf("\t%-5v # %v\n", strings.ToLower(string(locale)), name)
				}
			}
			return send + "```"
		},
	})
}
