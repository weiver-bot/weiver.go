package builder

import "github.com/bwmarrin/discordgo"

type TextInputStructure struct {
	data discordgo.TextInput
}

func TextInput() *TextInputStructure {
	return &TextInputStructure{}
}

func (t *TextInputStructure) SetCustomId(value string) *TextInputStructure {
	t.data.CustomID = value
	return t
}

func (t *TextInputStructure) SetLable(value string) *TextInputStructure {
	t.data.Label = value
	return t
}

func (t *TextInputStructure) SetMaxLength(value int) *TextInputStructure {
	t.data.MaxLength = value
	return t
}

func (t *TextInputStructure) SetMinLength(value int) *TextInputStructure {
	t.data.MinLength = value
	return t
}

func (t *TextInputStructure) SetPlaceholder(value string) *TextInputStructure {
	t.data.Placeholder = value
	return t
}

func (t *TextInputStructure) SetRequired(value bool) *TextInputStructure {
	t.data.Required = value
	return t
}

func (t *TextInputStructure) SetStyle(value discordgo.TextInputStyle) *TextInputStructure {
	t.data.Style = value
	return t
}

func (t *TextInputStructure) SetValue(value string) *TextInputStructure {
	t.data.Value = value
	return t
}
