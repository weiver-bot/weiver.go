package builder

import "github.com/bwmarrin/discordgo"

type SelectMenuOptionStructure struct {
	discordgo.SelectMenuOption
}

func SelectMenuOption() *SelectMenuOptionStructure {
	return &SelectMenuOptionStructure{}
}

func (s *SelectMenuOptionStructure) SetLabel(value string) *SelectMenuOptionStructure {
	s.Label = value
	return s
}

func (s *SelectMenuOptionStructure) SetDescription(value string) *SelectMenuOptionStructure {
	s.Description = value
	return s
}

func (s *SelectMenuOptionStructure) SetValue(value string) *SelectMenuOptionStructure {
	s.Value = value
	return s
}
