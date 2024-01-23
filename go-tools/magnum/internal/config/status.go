package config

type BookStatus string

const (
	Ongoing   BookStatus = "O"
	Hiatus    BookStatus = "H"
	Completed BookStatus = "C"
)

const (
	OngoingDisplay   string = "Ongoing"
	HiatusDisplay    string = "Hiatus"
	CompletedDisplay string = "Completed"
)

func IsStatus(val string) bool {
	switch val {
	case string(Ongoing):
		return true
	case string(Hiatus):
		return true
	case string(Completed):
		return true
	default:
		return false
	}
}

func BookStatusToDisplayText(val BookStatus) string {
	switch val {
	case Ongoing:
		return OngoingDisplay
	case Hiatus:
		return HiatusDisplay
	case Completed:
		return CompletedDisplay
	default:
		return ""
	}
}
