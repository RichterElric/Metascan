package Log

import (
	"Metascan/main/log_templates/Entry"
)
type SeverityCounter struct {
	high int
	medium int
	low int
	info int
}

type Log struct {
	scan_date string
	scan_types        []string
	severity_counters SeverityCounter
	entries           []Entry.Entry
}

func New(scan_date string, scan_types []string, severity_counters [4]int, entries []Entry.Entry) *Log {
	sev_encounters := SeverityCounter{
		severity_counters[0],
		severity_counters[1],
		severity_counters[2],
		severity_counters[3],
	}
	return &Log{scan_date, scan_types, sev_encounters, entries};
}