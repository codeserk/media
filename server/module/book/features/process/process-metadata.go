package process

import (
	"fmt"
	"media/module/book"
	"strings"
)

type analysisResponse struct {
	Description string   `json:"description"`
	Authors     []string `json:"authors"`
	Genres      []string `json:"genres"`
	Themes      []string `json:"themes"`
	Mood        []string `json:"mood"`
	Settings    []string `json:"settings"`
	AgeGroups   []string `json:"ageGroups"`
	Pacing      []string `json:"pacing"`
}

func (s *service) ProcessMetadata(metadata *book.Metadata) (*book.Metadata, error) {
	var res analysisResponse
	text, err := s.ai.Structured(fmt.Sprintf(`
		"Analyze the following book details and extract the specified information:
		Title: %s
		Description: %s
		Author: %s
		Tags: %s

		Response must be JSON, for each field only write one or multiple valid options (empty if none apply)
		Make sure the language is: SPANISH
		Based on the provided information, please identify and list the following:
		- description: Revised description, fix grammar and writing errors. Leave empty if is already okay.
		- authors: Array of authors. Return a revised version of the authors. Remove duplicates and make sure the casing is okay.
		- genres: Select the most relevant genres from the following list: %s.
		- themes: Identify any themes present in the description from this list: %s.
		- mood: Determine the mood or tone from the description using these options: %s.
		- settings: Infer the setting from the description using these options: %s.
		- ageGroups: Suggest the most likely target age group from these options: %s.
		- pacing: Determine the pacing or style from the description using these options: %s"`,
		metadata.Title,
		metadata.Description,
		strings.Join(metadata.Authors, ", "),
		strings.Join(metadata.Tags, ", "),
		book.AllGenres,
		book.AllThemes,
		book.AllMoods,
		book.AllSettings,
		book.AllAgeGroups,
		book.AllPacingTypes,
	), &res)
	if err != nil {
		return nil, fmt.Errorf("ai structured: %v", err)
	}

	fmt.Printf("chat: %v", text)

	if res.Description != "" {
		metadata.Description = res.Description
	}
	if len(res.Authors) > 0 {
		metadata.Authors = res.Authors
	}
	metadata.Genres = book.ToGenres(res.Genres)
	metadata.Themes = book.ToThemes(res.Themes)
	metadata.Moods = book.ToMoods(res.Mood)
	metadata.Settings = book.ToSettings(res.Settings)
	metadata.AgeGroups = book.ToAgeGroups(res.AgeGroups)
	metadata.PacingTypes = book.ToPacingTypes(res.Pacing)

	return metadata, nil
}
