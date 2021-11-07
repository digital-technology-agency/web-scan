package env

var (
	// SimpleGeneratorVar ...
	SimpleGeneratorVar = `simple`
)

// CheckGenerator validate generator type
func CheckGenerator(input string) bool {
	switch input {
	default:
		return false
	case
		SimpleGeneratorVar:
		return true
	}
}
