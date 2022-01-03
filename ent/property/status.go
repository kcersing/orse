package property
type Status string

const (
	On  Status = "0"
	off  Status = "1"
)

// Values provides list valid values for Enum.
func (Status) Values() (kinds []string) {
	for _, s := range []Status{On, off} {
		kinds = append(kinds, string(s))
	}
	return
}