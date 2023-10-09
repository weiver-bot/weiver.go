package simple

import "github.com/bwmarrin/discordgo"

func Modal(data *discordgo.InteractionResponseData) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: data,
	}
}
