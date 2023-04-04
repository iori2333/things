package commands

import (
	"encoding/json"
	"errors"
	"things/base/utils"
)

type Type int

const (
	TypeCreateThing Type = iota
	TypeModifyThing
	TypeDeleteThing

	TypeCreateFeature
	TypeModifyFeature
	TypeOverwriteFeature
	TypeDeleteFeature

	TypeModifyStates
	TypeModifyTransitions

	TypeQueryThing
	TypeQueryThings

	TypeQueryFeature
	TypeQueryFeatures

	TypeQueryState
	TypeQueryStates

	TypeAskThing
	TypeTellThing
)

type Interface interface {
	CommandType() Type
	Execute(*utils.Future)
}

func unmarshal[T any](data []byte) (*T, error) {
	var cmd T
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}
	return &cmd, nil
}

func Unmarshal(t Type, data []byte) (Interface, error) {
	switch t {
	case TypeCreateThing:
		return unmarshal[CreateThing](data)
	case TypeModifyThing:
		return unmarshal[ModifyThing](data)
	case TypeDeleteThing:
		return unmarshal[DeleteThing](data)
	case TypeCreateFeature:
		return unmarshal[CreateFeature](data)
	case TypeModifyFeature:
		return unmarshal[ModifyFeature](data)
	case TypeOverwriteFeature:
		return unmarshal[OverwriteFeature](data)
	case TypeDeleteFeature:
		return unmarshal[DeleteFeature](data)
	case TypeModifyStates:
		return unmarshal[ModifyStates](data)
	case TypeModifyTransitions:
		return unmarshal[ModifyTransitions](data)
	case TypeQueryThing:
		return unmarshal[QueryThing](data)
	case TypeQueryThings:
		return unmarshal[QueryThings](data)
	case TypeQueryFeature:
		return unmarshal[QueryFeature](data)
	case TypeQueryFeatures:
		return unmarshal[QueryFeatures](data)
	case TypeQueryState:
		return unmarshal[QueryState](data)
	case TypeQueryStates:
		return unmarshal[QueryStates](data)
	case TypeAskThing:
		return unmarshal[AskThing](data)
	case TypeTellThing:
		return unmarshal[TellThing](data)
	default:
		return nil, errors.New("unknown command type")
	}
}
