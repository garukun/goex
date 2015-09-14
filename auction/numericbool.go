package auction
import "encoding/json"

// A boolean type to be marshaled into 0 or 1.
type NumericBool bool

func (this NumericBool) MarshalJSON() ([]byte, error) {
	if this {
		return json.Marshal(1)
	}
	return json.Marshal(0)
}

func (this *NumericBool) UnmarshalJSON(bytes []byte) error  {
	var num int
	err := json.Unmarshal(bytes, &num)
	if err != nil {
		return err
	}

	if num == 0 {
		*this = NumericBool(false)
	} else if num == 1 {
		*this = NumericBool(true)
	} else {
		return json.UnsupportedValueError
	}

	return nil
}