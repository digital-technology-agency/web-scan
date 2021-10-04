package config

import (
	"encoding/json"
	"github.com/digital-technology-agency/web-scan/pkg/database"
	"github.com/digital-technology-agency/web-scan/pkg/database/sqlite"
	"github.com/digital-technology-agency/web-scan/pkg/env"
	"io/ioutil"
)

// Configuration configuration type
type Configuration struct {
	ProcessCount     int                `json:"process_count"`
	Alphabet         string             `json:"alphabet"`
	UrlLen           int                `json:"url_len"`
	ConcurrencyCount int                `json:"concurrency_count"`
	DataStoreType    string             `json:"data_store_type"`
	DataStore        database.DbService `json:"-"`
	ProtocolTypes    []string           `json:"protocol_types"`
}

// Default default configuration
func Default() Configuration {
	return Configuration{
		ProcessCount:     1,
		Alphabet:         "abcdefg",
		UrlLen:           2,
		ConcurrencyCount: 5,
		DataStoreType:    env.SQLITE_STORE,
		DataStore:        sqlite.SqLite{},
		ProtocolTypes:    []string{env.HTTP_PROTOCOL, env.HTTPS_PROTOCOL},
	}
}

// Load configuration from file
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result Configuration
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Save configuration save to file
func (cfg Configuration) Save(path string) error {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bytes, 0644)
	return err
}

// InitDataStore init data store from type
func (cfg *Configuration) InitDataStore() {
	switch cfg.DataStoreType {
	default:
		cfg.DataStore = sqlite.SqLite{}
	case env.SQLITE_STORE:
		cfg.DataStore = sqlite.SqLite{}
	}
}
