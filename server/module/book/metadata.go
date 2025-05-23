package book

import (
	"media/internal/util"
	"strings"
	"time"

	"github.com/samber/lo"
)

type Metadata struct {
	Title       string
	Description string
	Authors     []string
	ISBN        string
	EAN         string
	Publisher   string
	Tags        []string
	Genres      Genres
	Themes      Themes
	Moods       Moods
	Settings    Settings
	AgeGroups   AgeGroups
	PacingTypes PacingTypes
	PageCount   int
	PublishedAt time.Time
}

func (m *Metadata) IsComplete() bool {
	return m.Title != "" && m.Description != "" && len(m.Authors) > 0 && m.ISBN != "" && len(m.Genres) > 0
}

func (m *Metadata) IsValid() bool {
	return m.Title != "" && m.ISBN != ""
}

func (m *Metadata) Merge(other *Metadata) Metadata {
	merged := Metadata{
		Title:       other.Title,
		Description: other.Description,
		Authors:     other.Authors,
		ISBN:        other.ISBN,
		EAN:         other.EAN,
		Publisher:   other.Publisher,
		Genres:      other.Genres,
		PageCount:   other.PageCount,
		PublishedAt: other.PublishedAt,
	}
	if merged.Title == "" && m.Title != "" {
		merged.Title = m.Title
	}
	if merged.Description == "" && m.Description != "" {
		merged.Description = m.Description
	}
	if len(merged.Authors) == 0 && len(m.Authors) > 0 {
		merged.Authors = m.Authors
	}
	if merged.ISBN == "" && m.ISBN != "" {
		merged.ISBN = m.ISBN
	}
	if merged.EAN == "" && m.EAN != "" {
		merged.EAN = m.EAN
	}
	if merged.Publisher == "" && m.Publisher != "" {
		merged.Publisher = m.Publisher
	}
	if len(merged.Genres) == 0 && len(m.Genres) > 0 {
		merged.Genres = m.Genres
	}
	if merged.PageCount == 0 && m.PageCount != 0 {
		merged.PageCount = m.PageCount
	}
	if merged.PublishedAt.IsZero() && !m.PublishedAt.IsZero() {
		merged.PublishedAt = m.PublishedAt
	}

	return merged
}

// Genres

type Genre string
type Genres = []Genre

const (
	GenreFantasy            Genre = "fantasy"
	GenreScienceFiction     Genre = "scienceFiction"
	GenreMystery            Genre = "mystery"
	GenreThriller           Genre = "thriller"
	GenreRomance            Genre = "romance"
	GenreHorror             Genre = "horror"
	GenreHistoricalFiction  Genre = "historicalFiction"
	GenreAdventure          Genre = "adventure"
	GenreCrime              Genre = "crime"
	GenreChildrenLiterature Genre = "childrenLiterature"
	GenreBiography          Genre = "biography"
	GenreSelfHelp           Genre = "selfHelp"
	GenrePoetry             Genre = "poetry"
	GenreDrama              Genre = "drama"
	GenreComedy             Genre = "comedy"
	GenreNonFiction         Genre = "nonFiction"
	GenreDystopian          Genre = "dystopian"
	GenreParanormal         Genre = "paranormal"
	GenreFiction            Genre = "fiction"
)

func ToGenres(genres []string) []Genre {
	return lo.FilterMap(genres, func(item string, _ int) (Genre, bool) {
		item = util.Sanitize(item)
		if IsGenre(item) {
			return Genre(item), true
		}

		return "", false
	})
}

var validGenres = map[Genre]bool{
	GenreFantasy:            true,
	GenreScienceFiction:     true,
	GenreMystery:            true,
	GenreThriller:           true,
	GenreRomance:            true,
	GenreHorror:             true,
	GenreHistoricalFiction:  true,
	GenreAdventure:          true,
	GenreCrime:              true,
	GenreChildrenLiterature: true,
	GenreBiography:          true,
	GenreSelfHelp:           true,
	GenrePoetry:             true,
	GenreDrama:              true,
	GenreComedy:             true,
	GenreNonFiction:         true,
	GenreDystopian:          true,
	GenreParanormal:         true,
	GenreFiction:            true,
}
var AllGenres = strings.Join(lo.Map(lo.Keys(validGenres), func(item Genre, _ int) string { return string(item) }), ",")

func IsGenre(str string) bool {
	return validGenres[Genre(str)]
}

// Themes

type Theme string
type Themes = []Theme

const (
	ThemeLove                 Theme = "love"
	ThemeBetrayal             Theme = "betrayal"
	ThemeSurvival             Theme = "survival"
	ThemeIdentity             Theme = "identity"
	ThemePowerAndCorruption   Theme = "powerAndCorruption"
	ThemeGoodVsEvil           Theme = "goodVsEvil"
	ThemeComingOfAge          Theme = "comingOfAge"
	ThemeWarAndPeace          Theme = "warAndPeace"
	ThemeFreedomAndOppression Theme = "freedomAndOppression"
	ThemeJourneyAndQuest      Theme = "journeyAndQuest"
	ThemeForcedProximity      Theme = "forcedProximity"
	ThemeFakedDating          Theme = "fakedDating"
	ThemeEnemiesToLovers      Theme = "enemiesToLovers"
	ThemeGrumpyForSunshine    Theme = "grumpyForSunshine"
	ThemeSharedBed            Theme = "sharedBed"
	ThemeFriendsToLovers      Theme = "friendsToLovers"
	ThemeSecondChanceRomance  Theme = "secondChanceRomance"
	ThemeBullyRomance         Theme = "bullyRomance"
	ThemeRoyalRomance         Theme = "royalRomance"
	ThemeSoulmates            Theme = "soulmates"
	ThemeRedemption           Theme = "redemption"
	ThemeFoundFamily          Theme = "foundFamily"
	ThemeReincarnation        Theme = "reincarnation"
	ThemeMortalityAndDeath    Theme = "mortalityAndDeath"
	ThemeRevenge              Theme = "revenge"
	ThemeTimeTravel           Theme = "timeTravel"
	ThemeSmallTownRomance     Theme = "smallTownRomance"
	ThemeOfficeRomance        Theme = "officeRomance"
	ThemeAcademia             Theme = "academia"
	ThemeGreekMythology       Theme = "greekMythology"
)

func ToThemes(themes []string) []Theme {
	return lo.FilterMap(themes, func(item string, _ int) (Theme, bool) {
		item = util.Sanitize(item)
		if IsTheme(item) {
			return Theme(item), true
		}

		return "", false
	})
}

var validThemes = map[Theme]bool{
	ThemeLove:                 true,
	ThemeBetrayal:             true,
	ThemeSurvival:             true,
	ThemeIdentity:             true,
	ThemePowerAndCorruption:   true,
	ThemeGoodVsEvil:           true,
	ThemeComingOfAge:          true,
	ThemeWarAndPeace:          true,
	ThemeFreedomAndOppression: true,
	ThemeJourneyAndQuest:      true,
	ThemeForcedProximity:      true,
	ThemeFakedDating:          true,
	ThemeEnemiesToLovers:      true,
	ThemeGrumpyForSunshine:    true,
	ThemeSharedBed:            true,
	ThemeFriendsToLovers:      true,
	ThemeSecondChanceRomance:  true,
	ThemeBullyRomance:         true,
	ThemeRoyalRomance:         true,
	ThemeSoulmates:            true,
	ThemeRedemption:           true,
	ThemeFoundFamily:          true,
	ThemeReincarnation:        true,
	ThemeMortalityAndDeath:    true,
	ThemeRevenge:              true,
	ThemeTimeTravel:           true,
	ThemeSmallTownRomance:     true,
	ThemeOfficeRomance:        true,
	ThemeAcademia:             true,
	ThemeGreekMythology:       true,
}
var AllThemes = strings.Join(lo.Map(lo.Keys(validThemes), func(item Theme, _ int) string { return string(item) }), ",")

func IsTheme(str string) bool {
	return validThemes[Theme(str)]
}

// Moods

type Mood string
type Moods = []Mood

const (
	MoodDark          Mood = "dark"
	MoodLightHearted  Mood = "lightHearted"
	MoodSuspenseful   Mood = "suspenseful"
	MoodMelancholic   Mood = "melancholic"
	MoodHopeful       Mood = "hopeful"
	MoodHumorous      Mood = "humorous"
	MoodInspirational Mood = "inspirational"
	MoodTense         Mood = "tense"
	MoodWhimsical     Mood = "whimsical"
	MoodNostalgic     Mood = "nostalgic"
)

func ToMoods(moods []string) []Mood {
	return lo.FilterMap(moods, func(item string, _ int) (Mood, bool) {
		item = util.Sanitize(item)
		if IsMood(item) {
			return Mood(item), true
		}

		return "", false
	})
}

var validMoods = map[Mood]bool{
	MoodDark:          true,
	MoodLightHearted:  true,
	MoodSuspenseful:   true,
	MoodMelancholic:   true,
	MoodHopeful:       true,
	MoodHumorous:      true,
	MoodInspirational: true,
	MoodTense:         true,
	MoodWhimsical:     true,
	MoodNostalgic:     true,
}
var AllMoods = strings.Join(lo.Map(lo.Keys(validMoods), func(item Mood, _ int) string { return string(item) }), ",")

func IsMood(str string) bool {
	return validMoods[Mood(str)]
}

// Settings

type Setting string
type Settings = []Setting

const (
	SettingUrban        Setting = "urban"
	SettingRural        Setting = "rural"
	SettingHistorical   Setting = "historical"
	SettingFuturistic   Setting = "futuristic"
	SettingFantasyWorld Setting = "fantasyWorld"
)

func ToSettings(settings []string) []Setting {
	return lo.FilterMap(settings, func(item string, _ int) (Setting, bool) {
		item = util.Sanitize(item)
		if IsSetting(item) {
			return Setting(item), true
		}

		return "", false
	})
}

var validSettings = map[Setting]bool{
	SettingUrban:        true,
	SettingRural:        true,
	SettingHistorical:   true,
	SettingFuturistic:   true,
	SettingFantasyWorld: true,
}
var AllSettings = strings.Join(lo.Map(lo.Keys(validSettings), func(item Setting, _ int) string { return string(item) }), ",")

func IsSetting(str string) bool {
	return validSettings[Setting(str)]
}

// Age groups

type AgeGroup string
type AgeGroups = []AgeGroup

const (
	AgeGroupChildren   AgeGroup = "children"
	AgeGroupYoungAdult AgeGroup = "youngAdult"
	AgeGroupAdult      AgeGroup = "adult"
	AgeGroupSenior     AgeGroup = "senior"
)

func ToAgeGroups(ageGroups []string) []AgeGroup {
	return lo.FilterMap(ageGroups, func(item string, _ int) (AgeGroup, bool) {
		item = util.Sanitize(item)
		if IsAgeGroup(item) {
			return AgeGroup(item), true
		}

		return "", false
	})
}

var validAgeGroups = map[AgeGroup]bool{
	AgeGroupChildren:   true,
	AgeGroupYoungAdult: true,
	AgeGroupAdult:      true,
	AgeGroupSenior:     true,
}
var AllAgeGroups = strings.Join(lo.Map(lo.Keys(validAgeGroups), func(item AgeGroup, _ int) string { return string(item) }), ",")

func IsAgeGroup(str string) bool {
	return validAgeGroups[AgeGroup(str)]
}

// Pacing types

type Pacing string
type PacingTypes = []Pacing

const (
	PacingFastPaced             Pacing = "fastPaced"
	PacingSlowBurn              Pacing = "slowBurn"
	PacingEpistolary            Pacing = "epistolary"
	PacingStreamOfConsciousness Pacing = "streamOfConsciousness"
	PacingNonLinearNarrative    Pacing = "nonLinearNarrative"
)

func ToPacingTypes(pacingTypes []string) []Pacing {
	return lo.FilterMap(pacingTypes, func(item string, _ int) (Pacing, bool) {
		item = util.Sanitize(item)
		if IsPacing(item) {
			return Pacing(item), true
		}

		return "", false
	})
}

var validPacingTypes = map[Pacing]bool{
	PacingFastPaced:             true,
	PacingSlowBurn:              true,
	PacingEpistolary:            true,
	PacingStreamOfConsciousness: true,
	PacingNonLinearNarrative:    true,
}
var AllPacingTypes = strings.Join(lo.Map(lo.Keys(validPacingTypes), func(item Pacing, _ int) string { return string(item) }), ",")

func IsPacing(str string) bool {
	return validPacingTypes[Pacing(str)]
}
