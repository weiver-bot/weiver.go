package events

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/bwmarrin/discordgo"

	g "github.com/y2hO0ol23/weiver/handler"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func init() {
	prefix := os.Getenv("COMMAND_PREFIX")

	g.EventList = append(g.EventList, func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if !strings.HasPrefix(m.Content, prefix) {
			return
		}

		queries := strings.Split(m.Content[len(prefix):], " ")

		if len(queries) == 0 {
			return
		}

		cmd := builder.Button().
			SetStyle(discordgo.SecondaryButton).
			SetLable(fmt.Sprintf("Use Command /%v", queries[0])).
			SetCustomID("cmd")

		reply, err := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: "",
			Components: []discordgo.MessageComponent{
				builder.ActionRow().AddComponents(cmd).ActionsRow,
			},
			Reference: &discordgo.MessageReference{
				MessageID: m.ID,
				ChannelID: m.ChannelID,
				GuildID:   m.GuildID,
			},
			AllowedMentions: &discordgo.MessageAllowedMentions{
				RepliedUser: false,
			},
		})
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			return
		}

		locale, handler := func() (*discordgo.Locale, *g.CMDForm) {
			for _, handler := range g.CMDList {
				for locale, name := range *(handler.Data.NameLocalizations) {
					if name == queries[0] && handler.Message != nil {
						return &locale, &handler
					}
				}
			}
			return nil, nil
		}()

		var handle func(*discordgo.Session, *discordgo.InteractionCreate)
		handle = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionMessageComponent || i.Interaction.Member.User.ID != m.Author.ID {
				s.AddHandlerOnce(handle)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredMessageUpdate,
				})
				return
			}

			data := i.MessageComponentData()
			if i.Message.ID != reply.ID || data.ComponentType != discordgo.ButtonComponent {
				s.ChannelMessageDelete(m.ChannelID, m.ID)
				s.ChannelMessageDelete(reply.ChannelID, reply.ID)
				return
			}

			content := "`/? <locale>`"
			if handler != nil {
				content = handler.Message(s, i, *locale, queries[1:])
			}

			if content != "" {
				s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
					Content: content,
					Flags:   discordgo.MessageFlagsEphemeral,
				}))
			}

			s.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Channel: reply.ChannelID,
				ID:      reply.ID,
				Components: []discordgo.MessageComponent{
					builder.ActionRow().AddComponents(cmd.SetDisabled(true)).ActionsRow,
				},
			})
		}
		s.AddHandlerOnce(handle)
	})
}
