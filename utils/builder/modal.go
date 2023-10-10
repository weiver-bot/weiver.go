package builder

import "github.com/bwmarrin/discordgo"

type ModalStructure struct {
	data *discordgo.InteractionResponse
}

func Modal() *ModalStructure {
	return &ModalStructure{
		data: &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{},
		},
	}
}

func (m *ModalStructure) Data() *discordgo.InteractionResponse {
	return m.data
}

func (m *ModalStructure) SetCustomId(value string) *ModalStructure {
	m.data.Data.CustomID = value
	return m
}

func (m *ModalStructure) SetTitle(value string) *ModalStructure {
	m.data.Data.Title = value
	return m
}

func (m *ModalStructure) AddComponents(values ...*ActionRowStructure) *ModalStructure {
	for _, value := range values {
		m.data.Data.Components = append(m.data.Data.Components, value.data)
	}
	return m
}

func (m *ModalStructure) SetComponents(values ...*ActionRowStructure) *ModalStructure {
	m.data.Data.Components = make([]discordgo.MessageComponent, 0)
	return m.AddComponents(values...)
}
