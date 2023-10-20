package strings

import "fmt"

var kilobytesInAMegabyte float64 = 1024
var kilobytesInAGigabyte float64 = 1000000

func KbSizeToString(size float64) string {
	if size > kilobytesInAGigabyte {
		return fmt.Sprintf("%.2f GB", size/kilobytesInAGigabyte)
	} else if size > kilobytesInAMegabyte {
		return fmt.Sprintf("%.2f MB", size/kilobytesInAMegabyte)
	}

	return fmt.Sprintf("%.2f KB", size)
}
