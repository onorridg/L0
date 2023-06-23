package castomType

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type FlexFloat float64

func (ff *FlexFloat) Scan(value interface{}) error {
	if value == nil {
		*ff = 0.0
		return nil
	}

	if v, ok := value.(int64); ok {
		*ff = FlexFloat(v)
		return nil
	} else if v, ok := value.(float64); ok {
		*ff = FlexFloat(v)
		return nil
	} else if v, ok := value.(string); ok {
		res, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		*ff = FlexFloat(res)
		return nil
	}

	return fmt.Errorf("failed to scan FlexStr value")
}

func (ff *FlexFloat) UnmarshalJSON(b []byte) error {
	var value interface{}
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	switch v := value.(type) {
	case float64:
		*ff = FlexFloat(v)
	case string:
		fNum, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		*ff = FlexFloat(fNum)
	default:
		return fmt.Errorf("unsupported value type: %T", v)
	}
	return nil

}
