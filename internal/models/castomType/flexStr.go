package castomType

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type FlexStr string

func (fs *FlexStr) Scan(value interface{}) error {
	if value == nil {
		*fs = ""
		return nil
	}

	if v, ok := value.(int64); ok {
		*fs = FlexStr(strconv.FormatInt(v, 10))
		return nil
	} else if v, ok := value.(float64); ok {
		*fs = FlexStr(strconv.FormatFloat(v, 'f', -1, 64))
		return nil
	} else if v, ok := value.(string); ok {
		*fs = FlexStr(v)
		return nil
	}

	return fmt.Errorf("failed to scan FlexStr value")
}

func (fs *FlexStr) UnmarshalJSON(b []byte) error {
	var value interface{}
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	switch v := value.(type) {
	case float64:
		*fs = FlexStr(strconv.FormatInt(int64(v), 10))
	case string:
		*fs = FlexStr(v)
	default:
		return fmt.Errorf("unsupported value type: %T", v)
	}
	return nil
}
