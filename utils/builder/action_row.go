package builder

import "github.com/bwmarrin/discordgo"

type ActionRowStructure struct {
	discordgo.ActionsRow
}

func ActionRow() *ActionRowStructure {
	return &ActionRowStructure{}
}

func (a *ActionRowStructure) AddComponents(values ...interface{}) *ActionRowStructure {
	for _, value := range values {
		switch value.(type) {
		case *TextInputStructure:
			a.Components = append(a.Components, value.(*TextInputStructure).TextInput)
		case *ButtonStructure:
			a.Components = append(a.Components, value.(*ButtonStructure).Button)
		}
	}
	return a
}

func (a *ActionRowStructure) SetComponents(values ...interface{}) *ActionRowStructure {
	a.Components = make([]discordgo.MessageComponent, 0)
	return a.AddComponents(values...)
}
