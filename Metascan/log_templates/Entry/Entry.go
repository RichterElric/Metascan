package Entry

type Entry struct {
	Filename string
	Issue_name string
	Severity string
	CVE string
	CWE string
	Description string
	Fix string
}

func New(filename string, issue_name string, severity string, CVE string, CWE string, description string, fix string) *Entry {
	return &Entry{filename, issue_name,severity,CVE,CWE,description,fix}
}