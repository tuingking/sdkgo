package osql

import "encoding/json"

type secretstring string

func (s secretstring) MarshalJSON() ([]byte, error) {
	return json.Marshal("SECRET")
}
