package env

// SimpleGeneratorVar ...
var SimpleGeneratorVar = `simple` //nolint:gochecknoglobals

// CheckGenerator validate generator type.
func CheckGenerator(input string) bool {
	switch input {
	default:
		return false
	case SimpleGeneratorVar:
		return true
	}
}
