package config

import (
	"encoding/json"
	"errors"
	"github.com/digital-technology-agency/web-scan/pkg/database"
	"github.com/digital-technology-agency/web-scan/pkg/database/sqlite"
	"github.com/digital-technology-agency/web-scan/pkg/env"
	generators "github.com/digital-technology-agency/web-scan/pkg/services/generators"
	"io/ioutil"
	"strings"
)

// Configuration configuration type
type Configuration struct {
	ProcessCount     int                  `json:"process_count"`
	Alphabet         string               `json:"alphabet"`
	UrlLen           int                  `json:"url_len"`
	ConcurrencyCount int                  `json:"concurrency_count"`
	DataStoreType    string               `json:"data_store_type"`
	DataStore        database.DbService   `json:"-"`
	GeneratorType    string               `json:"generator_type"`
	Generator        generators.Generator `json:"-"`
	ProtocolTypes    []string             `json:"protocol_types"`
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
		GeneratorType:    env.SIMPLE_GENERATOR,
		Generator: &generators.SimpleGenerator{
			Alphabet: "abcdefg",
			Len:      2,
		},
		ProtocolTypes: []string{env.HTTP_PROTOCOL, env.HTTPS_PROTOCOL},
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

// Validate validate property params
func (cfg Configuration) Validate() error {
	if cfg.ProcessCount <= 0 {
		return errors.New("Поле [process_count] - должно быть больше 0!")
	}
	if strings.TrimSpace(strings.ToLower(cfg.Alphabet)) == "" {
		return errors.New("Поле [alphabet] - должно содержать набор символов!")
	}
	if cfg.UrlLen <= 0 {
		return errors.New("Поле [url_len] - должно быть больше 0!")
	}
	if cfg.ConcurrencyCount <= 0 {
		return errors.New("Поле [concurrency_count] - должно быть больше 0!")
	}
	if !env.CheckStore(cfg.DataStoreType) {
		return errors.New("Поле [data_store_type] - должно содержать значение из предложенных вариантов!")
	}
	if !env.CheckGenerator(cfg.GeneratorType) {
		return errors.New("Поле [generator_type] - должно содержать значение из предложенных вариантов!")
	}
	if len(cfg.ProtocolTypes) == 0 {
		return errors.New("Поле [protocol_types] - должно содержать одно или несколько занчение из предложенных вариантов!")
	}
	for _, protocolType := range cfg.ProtocolTypes {
		switch protocolType {
		default:
			return errors.New("Поле [protocol_types] - должно содержать одно или несколько занчение из предложенных вариантов!")
		case env.HTTP_PROTOCOL:
		case env.HTTPS_PROTOCOL:
			continue
		}
	}
	return nil
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

// InitGenerator init generator from type
func (cfg *Configuration) InitGenerator() {
	switch cfg.GeneratorType {
	default:
		cfg.Generator = &generators.SimpleGenerator{
			Alphabet: cfg.Alphabet,
			Len:      cfg.UrlLen,
		}
	case env.SIMPLE_GENERATOR:
		cfg.Generator = &generators.SimpleGenerator{
			Alphabet: cfg.Alphabet,
			Len:      cfg.UrlLen,
		}
	}
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
