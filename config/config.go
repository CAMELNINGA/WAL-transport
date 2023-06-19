package config

import (
	"time"

	"github.com/CAMELNINGA/cdc-postgres.git/pkg/postgres"
)

const (
	FilterType  = "filter"
	ReplaseType = "replase"
)

type Config struct {
	Database postgres.DatabaseCfg `json:"database"`
	Listener Listener             `json:"listener"`
	Kafka    Kafka                `json:"kafka"`
	Sanitize []Sanitize           `json:"sanitize"`
}

type Listener struct {
	RefreshConnection time.Duration `json:"refresh_connection"`
	SlotName          string        `json:"slot_name"`
}

// LoggerCfg path of the logger config.
type LoggerCfg struct {
	Caller bool   `long:"caller" env:"CALLER" description:"Caller"`
	Level  string `long:"level" env:"LEVEL" description:"Logger level"`
	Format string `long:"format" env:"FORMAT" description:"Logger format"`
}

type Kafka struct {
	Brokers []string `long:"brokers" env:"BROKERS" env-delim:"," description:"Kafka brokers"`
	Topic   string   `long:"topic" env:"TOPIC" description:"Kafka topic"`
	GroupID string   `long:"group-id" env:"GROUP_ID" description:"Kafka group id"`
}

type Sanitize struct {
	Type     string            `long:"type" env:"TYPE" description:"Sanitize type"`
	Table    string            `long:"table" env:"TABLE" description:"Table name"`
	OldTable string            `long:"old-table" env:"OLD_TABLE" description:"Old table name"`
	Schema   map[string]string `long:"schema" env:"SCHEMA" description:"Schema name"`
	Columns  map[string]string `long:"filter-columns" env:"FILTER_COLUMNS" description:"Filter columns"`
}
