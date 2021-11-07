package env

const (
	// SQLiteStore ...
	SQLiteStore = `sqlite`
	// JSONEachRowStore ...
	JSONEachRowStore = `jsoneachrow`
)

// CheckStore validate store type
func CheckStore(input string) bool {
	switch input {
	default:
		return false
	case
		SQLiteStore,
		JSONEachRowStore:
		return true
	}
}
