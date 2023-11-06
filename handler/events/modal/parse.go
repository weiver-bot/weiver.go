package modal

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type s struct{}

var (
	Parse s
)

func (_ s) CustomID(value string) (string, string, bool) {
	if strings.HasPrefix(value, "review") {
		data := strings.Split(value, "#")
		if len(data) == 3 {
			return data[1], data[2], true
		}
	}
	return "", "", false
}

func (_ s) ModalComponents(components []discordgo.MessageComponent) (int, string, string, bool) {
	if len(components) == 3 {
		score := func() int {
			count := strings.Count(components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value, "★")
			if count == 0 {
				return 1
			}
			return count
		}()

		title := components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
		content := components[2].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

		return score, title, content, true
	}
	return 0, "", "", false
}
