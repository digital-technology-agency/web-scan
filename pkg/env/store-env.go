package env

import (
	"github.com/digital-technology-agency/web-scan/pkg/database"
	"github.com/digital-technology-agency/web-scan/pkg/database/sqlite"
)

const (
	SQLITE_STORE        = `sqlite`
	JSON_EACH_ROW_STORE = `jsoneachrow`
)

func InitDbStore() map[string]database.DbService {
	result := map[string]database.DbService{
		SQLITE_STORE: sqlite.SqLite{},
	}
	return result
}

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

func CheckDataBaseStore(input string) bool {
	switch input {
	default:
		return false
	case SQLITE_STORE:
		return true
	}
}
