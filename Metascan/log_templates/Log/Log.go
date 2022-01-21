package Log

import (
	"Metascan/main/log_templates/Entry"
)
type SeverityCounter struct {
	High int
	Medium int
	Low int
	Info int
}

type Log struct {
	Scan_date  string
	Scan_types []string
	Severity_counters SeverityCounter
	Entries           []Entry.Entry
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