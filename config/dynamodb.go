package config

import (
    _ "embed"

    yaml "gopkg.in/yaml.v3"
    "github.com/pkg/errors"
)

//go:embed config.yml
var config []byte

type DynamoDbConfig struct {
    TableName string `yaml:"table_name"`
}

func NewDynamoDbConfig(isTest, isProduction bool) (DynamoDbConfig, error) {
    var empty DynamoDbConfig
    if isTest && isProduction {
        return empty, errors.New("only one of test or production must be set")
    }

    if isTest {
        panic("TODO")
    }

    if isProduction {
        panic("TODO")
    }

    var temp map[string]map[string]DynamoDbConfig
    err := yaml.Unmarshal(config, &temp)
    if err != nil {
        return empty, errors.New("only one of test or production must be set")
    }

    return temp["dynamodb"]["development"], nil
}
