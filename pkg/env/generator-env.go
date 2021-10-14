package env

var (
	SIMPLE_GENERATOR = `simple`
)

// CheckGenerator validate generator type
func CheckGenerator(input string) bool {
	switch input {
	default:
		return false
	case
		SIMPLE_GENERATOR:
		return true
	}
}
