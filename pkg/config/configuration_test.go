package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/digital-technology-agency/web-scan/pkg/database"
)

var (
	testConfiguration     = Default()
	testConfigurationPath = fmt.Sprintf("%s/wscan-config.json", os.TempDir())
)

func TestConfiguration_Save(t *testing.T) {
	type fields struct {
		ProcessCount     int
		Alphabet         string
		URLLen           int
		ConcurrencyCount int
		DataStoreType    string
		DataStore        database.DbService
		ProtocolTypes    []string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Save configuration",
			fields: fields{
				ProcessCount:     testConfiguration.ProcessCount,
				Alphabet:         testConfiguration.Alphabet,
				URLLen:           testConfiguration.URLLen,
				ConcurrencyCount: testConfiguration.ConcurrencyCount,
				DataStoreType:    testConfiguration.DataStoreType,
				DataStore:        testConfiguration.DataStore,
				ProtocolTypes:    testConfiguration.ProtocolTypes,
			},
			args: args{
				path: testConfigurationPath,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Configuration{
				ProcessCount:     tt.fields.ProcessCount,
				Alphabet:         tt.fields.Alphabet,
				URLLen:           tt.fields.URLLen,
				ConcurrencyCount: tt.fields.ConcurrencyCount,
				DataStoreType:    tt.fields.DataStoreType,
				DataStore:        tt.fields.DataStore,
				ProtocolTypes:    tt.fields.ProtocolTypes,
			}
			if err := cfg.Save(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *Configuration
		wantErr bool
	}{
		{
			name: "Load configuration",
			args: args{
				path: testConfigurationPath,
			},
			want: &Configuration{
				ProcessCount:     testConfiguration.ProcessCount,
				Alphabet:         testConfiguration.Alphabet,
				URLLen:           testConfiguration.URLLen,
				ConcurrencyCount: testConfiguration.ConcurrencyCount,
				DataStoreType:    testConfiguration.DataStoreType,
				DataStore:        nil,
				ProtocolTypes:    testConfiguration.ProtocolTypes,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfiguration_InitDataStore(t *testing.T) {
	type fields struct {
		ProcessCount     int
		Alphabet         string
		URLLen           int
		ConcurrencyCount int
		DataStoreType    string
		DataStore        database.DbService
		ProtocolTypes    []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Init data store",
			fields: fields{
				ProcessCount:     testConfiguration.ProcessCount,
				Alphabet:         testConfiguration.Alphabet,
				URLLen:           testConfiguration.URLLen,
				ConcurrencyCount: testConfiguration.ConcurrencyCount,
				DataStoreType:    testConfiguration.DataStoreType,
				DataStore:        nil,
				ProtocolTypes:    testConfiguration.ProtocolTypes,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Configuration{
				ProcessCount:     tt.fields.ProcessCount,
				Alphabet:         tt.fields.Alphabet,
				URLLen:           tt.fields.URLLen,
				ConcurrencyCount: tt.fields.ConcurrencyCount,
				DataStoreType:    tt.fields.DataStoreType,
				DataStore:        tt.fields.DataStore,
				ProtocolTypes:    tt.fields.ProtocolTypes,
			}
			cfg.InitDataStore()
			if cfg.DataStore == nil {
				t.Errorf("Data store not init!")
			}
		})
	}
}
