package container

import (
	"encoding/json"
	"fmt"
	"io"
)

type Map map[string]interface{}

func (m Map) ToJSONString() string {
	s := `{`
	for k, v := range m {
		s += fmt.Sprintf(`"%s":"%s",`, k, v)
	}
	s = s[:len(s)-1]
	s += `}`
	return s
}

func (m Map) FindMissingKeys(requiredKeys ...string) []string {
	var missingKeys []string
	for _, key := range requiredKeys {
		_, found := m[key]
		if !found {
			missingKeys = append(missingKeys, key)
		}
	}

	return missingKeys
}

func (m Map) FindForbiddenKeys(allowedKeys ...string) []string {
	var redundantKeys []string
	for key := range m {
		if !ArrayStringContains(allowedKeys, key) {
			redundantKeys = append(redundantKeys, key)
		}
	}
	return redundantKeys
}

func ArrayStringContains(a []string, v string) bool {
	for _, e := range a {
		if e == v {
			return true
		}
	}
	return false
}

// CreateMapFromReader creates map from a JSON reader
func CreateMapFromReader(reader io.Reader) (Map, error) {
	m := Map{}

	if reader == nil {
		return m, nil
	}

	// numbers are represented as string instead of float64
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()

	err := decoder.Decode(&m)
	if err == io.EOF {
		return m, nil
	}

	return m, err
}
