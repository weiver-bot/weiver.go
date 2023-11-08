package look

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	db "github.com/y2hO0ol23/weiver/database"
	this "github.com/y2hO0ol23/weiver/handler/slash-commands/look/reviews"
	"github.com/y2hO0ol23/weiver/localization"
	"github.com/y2hO0ol23/weiver/utils/builder"
)

var pageRow int

func init() {
	var err error
	pageRow, err = strconv.Atoi(os.Getenv("PAGE_ROW"))
	if err != nil || pageRow < 1 || 25 < pageRow {
		pageRow = 10
	}
}

func Reviews(s *discordgo.Session, i *discordgo.InteractionCreate, subjectID string) {
	locale := i.Locale

	// load subject data
	subject, err := s.GuildMember(i.GuildID, subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}

	// load reviews from db
	reviews, err := db.GetReviewsByUserID(subjectID)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		return
	}
	if reviews == nil { // if no reviews
		message := builder.Message(&discordgo.InteractionResponseData{
			Content: fmt.Sprintf("`%v`", localization.Load(locale, "#look.reviews.IsNone")),
			Flags:   discordgo.MessageFlagsEphemeral,
		})
		err = s.InteractionRespond(i.Interaction, message)
		if err != nil {
			log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
		}
		return
	}

	// set select menu
	pageNow := 1
	pageCount := (len(*reviews)-1)/pageRow + 1
	selectMenu := SelectMenu(*reviews, locale, subject.User.Username, pageNow, pageCount)

	// send select menu
	err = s.InteractionRespond(i.Interaction, builder.Message(&discordgo.InteractionResponseData{
		Components: []discordgo.MessageComponent{
			builder.ActionRow().AddComponents(selectMenu).ActionsRow,
		},
		Flags: discordgo.MessageFlagsEphemeral,
	}))
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}

	msg, err := s.InteractionResponse(i.Interaction)
	if err != nil {
		log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
	}

	// make handler because bot can not find message that has ephemeral flag
	var handle func(*discordgo.Session, *discordgo.InteractionCreate)
	handle = func(s *discordgo.Session, iter *discordgo.InteractionCreate) {
		if iter.Type != discordgo.InteractionMessageComponent || i.Interaction.Member.User.ID != iter.Interaction.Member.User.ID {
			s.AddHandlerOnce(handle)
			return
		}
		if iter.Interaction.Message.ID != msg.ID {
			return
		}

		data := iter.MessageComponentData()
		if data.CustomID != "reviews" {
			return
		}

		value := data.Values[0]
		locale = iter.Locale

		if strings.HasPrefix(value, "page/") { // move page
			pageNow, err := strconv.Atoi(strings.Split(value, ":")[1])
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}
			selectMenu = SelectMenu(*reviews, locale, subject.User.Username, pageNow, pageCount)
			this.MovePage(s, iter, i, selectMenu)

			s.AddHandlerOnce(handle)
		} else if strings.HasPrefix(value, "review") { // show page link
			id_timestamp := strings.Split(value[7:], "#")
			// parse id
			id, err := strconv.Atoi(id_timestamp[0])
			if err != nil {
				log.Printf("[ERROR] %v\n%v\n", err, string(debug.Stack()))
				return
			}
			this.Selected(s, iter, i, id, id_timestamp[1], locale)
		}
	}
	s.AddHandlerOnce(handle)
}
