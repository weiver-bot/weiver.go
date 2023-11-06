package admin

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	"github.com/y2hO0ol23/weiver/localization"
	TagUtils "github.com/y2hO0ol23/weiver/utils/bot/tag"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

func AllowRole(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) {
	locale := i.Locale
	value := options[0].Value.(bool)

	guildDB, err := db.LoadGuildByID(i.GuildID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
	if guildDB.InProgress == true {
		err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
			Content:         fmt.Sprintf("`%v`", localization.Load(locale, "#admin.allow-role.InProgress")),
			Flags:           discordgo.MessageFlagsEphemeral,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		}))
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
		return
	}
	if value == true {
		p, err := s.UserChannelPermissions(s.State.User.ID, i.ChannelID)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
		if p&discordgo.PermissionManageRoles == 0 {
			err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
				Content:         fmt.Sprintf("`%v`", localization.Load(locale, "#admin.allow-role.NeedPermissions")),
				Flags:           discordgo.MessageFlagsEphemeral,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			}))
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
			}
			return
		}
	}
	if guildDB.AllowRole != value {
		db.UpdateGuildRoleOption(i.GuildID, value)

		err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				builder.Embed().
					SetDescription(fmt.Sprintf("**%v** `%v` - %v",
						localization.Load(locale, "#admin.allow-role.proc.Title"),
						localization.Load(locale, "#admin.allow-role.proc.Description"),
						localization.Load(locale, "#admin.allow-role.proc.InProgress"),
					)).
					MessageEmbed,
			},
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		}))
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}

		var (
			after string
			wait  sync.WaitGroup
		)
		for {
			res, err := s.GuildMembers(i.GuildID, after, 1000)
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				continue
			}
			if len(res) == 0 {
				break
			}

			wait.Add(1)
			go func(members []*discordgo.Member) {
				defer wait.Done()
				for _, member := range members {
					display, err := TagUtils.GetScoreUIShort(member.User.ID)
					if err != nil {
						log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
						continue
					}
					if value == true {
						err = TagUtils.AddTag(s, i.GuildID, member.User.ID, display)
					} else {
						err = TagUtils.RemoveTag(s, i.GuildID, member.User.ID, display)
					}
					if err != nil {
						log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
						continue
					}
				}
			}(res)

			after = res[len(res)-1].User.ID
		}
		wait.Wait()

		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				builder.Embed().
					SetDescription(fmt.Sprintf("**%v** `%v` - %v",
						localization.Load(locale, "#admin.allow-role.proc.Title"),
						localization.Load(locale, "#admin.allow-role.proc.Description"),
						localization.Load(locale, "#admin.allow-role.proc.Done"),
					)).
					MessageEmbed,
			},
		})
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}

		db.EndOFGuildProgress(i.GuildID)
	} else {
		err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
			Content:         fmt.Sprintf("`%v`", localization.Load(locale, "#admin.allow-role.Keep")),
			Flags:           discordgo.MessageFlagsEphemeral,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		}))
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
	}
}
