package builder

import "github.com/bwmarrin/discordgo"

type ActionRowStructure struct {
	data discordgo.ActionsRow
}

func ActionRow() *ActionRowStructure {
	return &ActionRowStructure{}
}

func (a *ActionRowStructure) AddComponents(values ...*TextInputStructure) *ActionRowStructure {
	for _, value := range values {
		a.data.Components = append(a.data.Components, value.data)
	}
	return a
}

func (a *ActionRowStructure) SetComponents(values ...*TextInputStructure) *ActionRowStructure {
	a.data.Components = make([]discordgo.MessageComponent, 0)
	return a.AddComponents(values...)
}
