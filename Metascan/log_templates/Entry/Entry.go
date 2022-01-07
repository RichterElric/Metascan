package Entry

type Entry struct {
	filename string
	issue_name string
	severity string
	CVE string
	CWE string
	description string
	fix string
}

func New(filename string, issue_name string, severity string, CVE string, CWE string, description string, fix string) *Entry {
	return &Entry{filename, issue_name,severity,CVE,CWE,description,fix}
}