package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/digital-technology-agency/web-scan/pkg/database"
	"github.com/digital-technology-agency/web-scan/pkg/database/sqlite"
	"github.com/digital-technology-agency/web-scan/pkg/env"
	generators "github.com/digital-technology-agency/web-scan/pkg/services/generators"
)

// Configuration configuration type.
type Configuration struct {
	ProcessCount     int                  `json:"process_count"`
	Alphabet         string               `json:"alphabet"`
	URLLen           int                  `json:"url_len"`
	ConcurrencyCount int                  `json:"concurrency_count"`
	DataStoreType    string               `json:"data_store_type"`
	DataStore        database.DbService   `json:"-"`
	GeneratorType    string               `json:"generator_type"`
	Generator        generators.Generator `json:"-"`
	ProtocolTypes    []string             `json:"protocol_types"`
}

// Default default configuration.
func Default() Configuration {
	return Configuration{
		ProcessCount:     1,
		Alphabet:         "abcdefg",
		URLLen:           2,
		ConcurrencyCount: 5,
		DataStoreType:    env.SQLiteStore,
		DataStore:        sqlite.SqLite{},
		GeneratorType:    env.SimpleGeneratorVar,
		Generator: &generators.SimpleGenerator{
			Alphabet: "abcdefg",
			Len:      2,
		},
		ProtocolTypes: []string{env.HTTPProtocol, env.HTTPSProtocol},
	}
}

// Load configuration from file.
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

// Validate validate property params.
func (cfg Configuration) Validate() error {
	if cfg.ProcessCount <= 0 {
		return errors.New("поле [process_count] - должно быть больше 0")
	}
	if strings.TrimSpace(strings.ToLower(cfg.Alphabet)) == "" {
		return errors.New("поле [alphabet] - должно содержать набор символов")
	}
	if cfg.URLLen <= 0 {
		return errors.New("поле [url_len] - должно быть больше 0")
	}
	if cfg.ConcurrencyCount <= 0 {
		return errors.New("поле [concurrency_count] - должно быть больше 0")
	}
	if !env.CheckStore(cfg.DataStoreType) {
		return errors.New("поле [data_store_type] - должно содержать значение из предложенных вариантов")
	}
	if !env.CheckGenerator(cfg.GeneratorType) {
		return errors.New("поле [generator_type] - должно содержать значение из предложенных вариантов")
	}
	if len(cfg.ProtocolTypes) == 0 {
		return errors.New("поле [protocol_types] - должно содержать одно или несколько занчение из предложенных вариантов")
	}
	for _, protocolType := range cfg.ProtocolTypes {
		switch protocolType {
		default:
			return errors.New("поле [protocol_types] - должно содержать одно или несколько занчение из предложенных вариантов")
		case env.HTTPProtocol:
		case env.HTTPSProtocol:
			continue
		}
	}
	return nil
}

// Save configuration save to file.
func (cfg Configuration) Save(path string) error {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bytes, 0o644)
	return err
}

// InitGenerator init generator from type.
func (cfg *Configuration) InitGenerator() {
	switch cfg.GeneratorType {
	default:
		cfg.Generator = &generators.SimpleGenerator{
			Alphabet: cfg.Alphabet,
			Len:      cfg.URLLen,
		}
	case env.SimpleGeneratorVar:
		cfg.Generator = &generators.SimpleGenerator{
			Alphabet: cfg.Alphabet,
			Len:      cfg.URLLen,
		}
	}
}

// InitDataStore init data store from type.
func (cfg *Configuration) InitDataStore() {
	switch cfg.DataStoreType {
	default:
		cfg.DataStore = sqlite.SqLite{}
	case env.SQLiteStore:
		cfg.DataStore = sqlite.SqLite{}
	}
}
