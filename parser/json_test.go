package parser_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tuingking/sdkgo/parser"
)

func Test__GetJSONToFloat64Array(t *testing.T) {
	t.Run("Success from byte", func(t *testing.T) {
		res, err := parser.GetJSONToFloat64Array([]byte(`[1.23]`))
		assert.Equal(t, res, []float64{1.23})
		assert.Nil(t, err)
	})
	t.Run("Failed from byte", func(t *testing.T) {
		res, err := parser.GetJSONToFloat64Array([]byte("123"))
		assert.Equal(t, res, []float64{})
		assert.NotNil(t, err)
	})
	t.Run("Success from array", func(t *testing.T) {
		res, err := parser.GetJSONToFloat64Array([]float64{1.23})
		assert.Equal(t, res, []float64{1.23})
		assert.Nil(t, err)
	})
	t.Run("Success from nil", func(t *testing.T) {
		res, err := parser.GetJSONToFloat64Array(nil)
		assert.Equal(t, res, []float64{})
		assert.Nil(t, err)
	})
	t.Run("Invalid scheme", func(t *testing.T) {
		res, err := parser.GetJSONToFloat64Array(123)
		assert.Equal(t, res, []float64{})
		assert.NotNil(t, err)
	})
}

func Test__GetJSONToInt64Array(t *testing.T) {
	t.Run("Success from byte", func(t *testing.T) {
		res, err := parser.GetJSONToInt64Array([]byte(`[123]`))
		assert.Equal(t, res, []int64{123})
		assert.Nil(t, err)
	})
	t.Run("Failed from byte", func(t *testing.T) {
		res, err := parser.GetJSONToInt64Array([]byte("123"))
		assert.Equal(t, res, []int64{})
		assert.NotNil(t, err)
	})
	t.Run("Success from array", func(t *testing.T) {
		res, err := parser.GetJSONToInt64Array([]int64{123})
		assert.Equal(t, res, []int64{123})
		assert.Nil(t, err)
	})
	t.Run("Success from nil", func(t *testing.T) {
		res, err := parser.GetJSONToInt64Array(nil)
		assert.Equal(t, res, []int64{})
		assert.Nil(t, err)
	})
	t.Run("Invalid scheme", func(t *testing.T) {
		res, err := parser.GetJSONToInt64Array(123)
		assert.Equal(t, res, []int64{})
		assert.NotNil(t, err)
	})
}

func Test__GetJSONToStringArray(t *testing.T) {
	t.Run("Success from byte", func(t *testing.T) {
		res, err := parser.GetJSONToStringArray([]byte(`["123"]`))
		assert.Equal(t, res, []string{"123"})
		assert.Nil(t, err)
	})
	t.Run("Failed from byte", func(t *testing.T) {
		res, err := parser.GetJSONToStringArray([]byte("123"))
		assert.Equal(t, res, []string{})
		assert.NotNil(t, err)
	})
	t.Run("Success from array", func(t *testing.T) {
		res, err := parser.GetJSONToStringArray([]string{"123"})
		assert.Equal(t, res, []string{"123"})
		assert.Nil(t, err)
	})
	t.Run("Success from nil", func(t *testing.T) {
		res, err := parser.GetJSONToStringArray(nil)
		assert.Equal(t, res, []string{})
		assert.Nil(t, err)
	})
	t.Run("Invalid scheme", func(t *testing.T) {
		res, err := parser.GetJSONToStringArray("123")
		assert.Equal(t, res, []string{})
		assert.NotNil(t, err)
	})
}

func Test__GetJSONToString(t *testing.T) {
	t.Run("Success from byte", func(t *testing.T) {
		res, err := parser.GetJSONToString([]byte("123"))
		assert.Equal(t, res, "123")
		assert.Nil(t, err)
	})
	t.Run("Success from string", func(t *testing.T) {
		res, err := parser.GetJSONToString("123")
		assert.Equal(t, res, "123")
		assert.Nil(t, err)
	})
	t.Run("Success from nil", func(t *testing.T) {
		res, err := parser.GetJSONToString(nil)
		assert.Equal(t, res, "")
		assert.Nil(t, err)
	})
	t.Run("Invalid scheme", func(t *testing.T) {
		res, err := parser.GetJSONToString(123)
		assert.Equal(t, res, "")
		assert.NotNil(t, err)
	})
}

func Test__JsonBytesFromInterface(t *testing.T) {
	type obj struct {
		Test string `json:"test"`
	}
	expectedObj := obj{
		Test: "foo",
	}
	expected, _ := json.Marshal(expectedObj)
	t.Run("Success from byte", func(t *testing.T) {
		res := parser.JsonBytesFromInterface(expected)
		assert.Equal(t, res, expected)
	})
	t.Run("Success from string", func(t *testing.T) {
		res := parser.JsonBytesFromInterface(string(expected))
		assert.Equal(t, res, expected)
	})
	t.Run("Success from object", func(t *testing.T) {
		res := parser.JsonBytesFromInterface(expectedObj)
		assert.Equal(t, res, expected)
	})
}
