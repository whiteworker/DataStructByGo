package deepcopy

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func DeepCopyByGob(dst, src interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(&buffer).Decode(dst)
}

func DeepCopyByJson(src []Book) (*[]Book, error) {
	var dst = new([]Book)
	b, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, dst)
	return dst, err
}
