package env

const (
	SQLITE_STORE        = `sqlite`
	JSON_EACH_ROW_STORE = `jsoneachrow`
)

// CheckStore validate store type
func CheckStore(input string) bool {
	switch input {
	default:
		return false
	case
		SQLITE_STORE,
		JSON_EACH_ROW_STORE:
		return true
	}
}
