package builder

import "github.com/bwmarrin/discordgo"

type TextInputStructure struct {
	discordgo.TextInput
}

func TextInput() *TextInputStructure {
	return &TextInputStructure{}
}

func (t *TextInputStructure) SetCustomId(value string) *TextInputStructure {
	t.CustomID = value
	return t
}

func (t *TextInputStructure) SetLable(value string) *TextInputStructure {
	t.Label = value
	return t
}

func (t *TextInputStructure) SetMaxLength(value int) *TextInputStructure {
	t.MaxLength = value
	return t
}

func (t *TextInputStructure) SetMinLength(value int) *TextInputStructure {
	t.MinLength = value
	return t
}

func (t *TextInputStructure) SetPlaceholder(value string) *TextInputStructure {
	t.Placeholder = value
	return t
}

func (t *TextInputStructure) SetRequired(value bool) *TextInputStructure {
	t.Required = value
	return t
}

func (t *TextInputStructure) SetStyle(value discordgo.TextInputStyle) *TextInputStructure {
	t.Style = value
	return t
}

func (t *TextInputStructure) SetValue(value string) *TextInputStructure {
	t.Value = value
	return t
}
