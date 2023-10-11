package builder

import "github.com/bwmarrin/discordgo"

type EmbedStructure struct {
	*discordgo.MessageEmbed
}

func Embed() *EmbedStructure {
	return &EmbedStructure{&discordgo.MessageEmbed{}}
}

func (e *EmbedStructure) SetTitle(name string) *EmbedStructure {
	e.Title = name
	return e
}

func (e *EmbedStructure) SetDescription(description string) *EmbedStructure {
	e.Description = description
	return e
}

func (e *EmbedStructure) SetFields(values ...*discordgo.MessageEmbedField) *EmbedStructure {
	e.Fields = values
	return e
}

func (e *EmbedStructure) AddFields(values ...*discordgo.MessageEmbedField) *EmbedStructure {
	e.Fields = append(e.Fields, values...)
	return e
}

func (e *EmbedStructure) SetFooter(value *discordgo.MessageEmbedFooter) *EmbedStructure {
	e.Footer = value
	return e
}

func (e *EmbedStructure) SetThumbnail(value *discordgo.MessageEmbedThumbnail) *EmbedStructure {
	e.Thumbnail = value
	return e
}
