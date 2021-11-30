package unmarshal

import (
	"encoding/json"
	"errors"
)

func Marshal(v interface{}) ([]byte, error) {
	var (
		b   []byte
		err error
	)
	if v == nil {
		return nil, errors.New("v is nil")
	}
	if b, err = json.Marshal(v); err != nil {
		return nil, err
	}

	return b, nil
}

func UnMarshal(bs []byte, v interface{}) error {
	var (
		err error
	)
	if bs == nil {
		return nil
	}
	if err = json.Unmarshal(bs, &v); err != nil {
		return err
	}

	return nil
}
