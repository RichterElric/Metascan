package scanners

type Scanner interface {
	GetDependency() bool
	Scan(path string) string
}