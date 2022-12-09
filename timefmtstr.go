package verbose

import (
	"strings"
)

// Converts Unix/Linux date stamp to Go Date Format.
// date +"%A %d %B %Y, %I:%M:%S %P %Z"
// Returns: Thursday December 12 2022, 06:03:08 P EST
// verbose.TimeFormatStr("%A %B %d %Y, %I:%M:%S %P %Z")
// Returns: "Monday January 02 2006, 03:05:05 PM MST". When put into time.Now().Format("Monday January 02 2006, 03:05:05 PM MST")
// Would give you: "Thursday December 12 2022, 06:03:08 PM EST"
func TimeFormatStr(tformat string) (fmtStr string) {
	tFormat := strings.Split(tformat, "")
	for i := 0; i < len(tFormat); i++ {
		if tFormat[i] == "%" {
			i++
			switch tFormat[i] {
			case "Y":
				fmtStr += "2006"
			case "m":
				fmtStr += "01"
			case "D":
				fmtStr += "01/02/06"
			case "B":
				fmtStr += "January"
			case "b":
				fmtStr += "Jan"
			case "d":
				fmtStr += "02"
			case "j":
				fmtStr += "002"
			case "A":
				fmtStr += "Monday"
			case "a":
				fmtStr += "Mon"
			case "H":
				fmtStr += "15"
			case "I":
				fmtStr += "03"
			case "M":
				fmtStr += "04"
			case "S":
				fmtStr += "05"
			case "N":
				fmtStr += ".000"
			case "F":
				fmtStr += "2006-01-02"
			case "T":
				fmtStr += "15:04:05"
			case "Z":
				fmtStr += "MST"
			case "P":
				fmtStr += "PM"
			}
		} else {
			fmtStr += tFormat[i]
		}

	}
	fmtStr = strings.TrimSpace(fmtStr)
	//fmtStr += " "
	return fmtStr
}
