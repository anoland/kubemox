package utils

import "fmt"

func FormatUptime(uptime int) string {
	// Convert seconds to format like 1d 2h 3m 4s
	days := uptime / 86400
	hours := (uptime - days*86400) / 3600
	minutes := (uptime - days*86400 - hours*3600) / 60
	seconds := uptime - days*86400 - hours*3600 - minutes*60
	uptimeString := fmt.Sprintf("%dd%dh%dm%ds", days, hours, minutes, seconds)
	return uptimeString
}

func SubstractSlices(slice1, slice2 []string) []string {
	elements := make(map[string]bool)
	for _, elem := range slice2 {
		elements[elem] = true
	}
	// Create a result slice to store the difference
	var difference []string
	// Iterate through slice1 and check if the element is present in slice2
	for _, elem := range slice1 {
		if !elements[elem] {
			difference = append(difference, elem)
		}
	}
	return difference
}