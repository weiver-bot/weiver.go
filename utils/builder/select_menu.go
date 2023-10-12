package builder

import "github.com/bwmarrin/discordgo"

type SelectMenuStructure struct {
	discordgo.SelectMenu
}

func SelectMenu() *SelectMenuStructure {
	return &SelectMenuStructure{}
}

func (s *SelectMenuStructure) SetCustomID(value string) *SelectMenuStructure {
	s.CustomID = value
	return s
}

func (s *SelectMenuStructure) SetPlaceholder(value string) *SelectMenuStructure {
	s.Placeholder = value
	return s
}

func (s *SelectMenuStructure) AddOptions(values ...*SelectMenuOptionStructure) *SelectMenuStructure {
	for _, value := range values {
		s.Options = append(s.Options, value.SelectMenuOption)
	}
	return s
}

func (s *SelectMenuStructure) SetOptions(values ...*SelectMenuOptionStructure) *SelectMenuStructure {
	s.Options = make([]discordgo.SelectMenuOption, 0)
	return s.AddOptions(values...)
}
