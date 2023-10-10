package builder

import "github.com/bwmarrin/discordgo"

type ModalStructure struct {
	*discordgo.InteractionResponse
}

func Modal() *ModalStructure {
	return &ModalStructure{
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{},
		},
	}
}

func (m *ModalStructure) SetCustomId(value string) *ModalStructure {
	m.Data.CustomID = value
	return m
}

func (m *ModalStructure) SetTitle(value string) *ModalStructure {
	m.Data.Title = value
	return m
}

func (m *ModalStructure) AddComponents(values ...*ActionRowStructure) *ModalStructure {
	for _, value := range values {
		m.Data.Components = append(m.Data.Components, value.ActionsRow)
	}
	return m
}

func (m *ModalStructure) SetComponents(values ...*ActionRowStructure) *ModalStructure {
	m.Data.Components = make([]discordgo.MessageComponent, 0)
	return m.AddComponents(values...)
}
