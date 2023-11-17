package config

type SeriesType string

const (
	WebNovel   SeriesType = "WN"
	Manga      SeriesType = "MN"
	LightNovel SeriesType = "LN"
)

func IsSeriesType(val string) bool {
	switch val {
	case string(WebNovel):
		return true
	case string(Manga):
		return true
	case string(LightNovel):
		return true
	default:
		return false
	}
}
