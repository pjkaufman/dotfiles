package config

type Status string

const (
	Ongoing   Status = "O"
	Hiatus    Status = "H"
	Completed Status = "C"
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
