package config

type PublisherType string

const (
	YenPress               PublisherType = "YenPress"
	JNovelClub             PublisherType = "JNovelClub"
	SevenSeasEntertainment PublisherType = "SevenSeasEntertainment"
)

func IsPublisherType(val string) bool {
	switch val {
	case string(YenPress):
		return true
	case string(JNovelClub):
		return true
	case string(SevenSeasEntertainment):
		return true
	default:
		return false
	}
}
