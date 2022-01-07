package scanners

type Scanner interface {
	Scan() string // Must return JSON string
}
