package parser

import (
	"encoding/json"
	"errors"
)

func GetJSONToFloat64Array(param interface{}) ([]float64, error) {
	result := []float64{}
	switch data := param.(type) {
	case []byte:
		if err := json.Unmarshal(data, &result); err != nil {
			return result, err
		}
	case []float64:
		result = data
	case nil:
		return result, nil
	default:
		return result, errors.New("invalid schema")
	}
	return result, nil
}

func GetJSONToInt64Array(param interface{}) ([]int64, error) {
	result := []int64{}
	switch data := param.(type) {
	case []byte:
		if err := json.Unmarshal(data, &result); err != nil {
			return result, err
		}
	case []int64:
		result = data
	case nil:
		return result, nil
	default:
		return result, errors.New("invalid schema")
	}
	return result, nil
}

func GetJSONToStringArray(param interface{}) ([]string, error) {
	result := []string{}
	switch data := param.(type) {
	case []byte:
		if err := json.Unmarshal(data, &result); err != nil {
			return result, err
		}
	case []string:
		result = data
	case nil:
		return result, nil
	default:
		return result, errors.New("invalid schema")
	}
	return result, nil
}

func GetJSONToString(param interface{}) (string, error) {
	var result string
	switch data := param.(type) {
	case []byte:
		result = string(data)
	case string:
		result = data
	case nil:
		return result, nil
	default:
		return result, errors.New("invalid schema")
	}
	return result, nil
}

func JsonBytesFromInterface(param interface{}) []byte {
	var res []byte
	switch data := param.(type) {
	case []byte:
		return data
	case string:
		return []byte(data)
	default:
		res, _ = json.Marshal(param)
	}
	return res
}
