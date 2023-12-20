package models

import "fmt"

// A Search Term Data struct takes values that are resolved from a regular expression
// where a regex has a (?<name>) named group, the key for the data is name, and the
// value is whatever was resolved from that regex. The parent object will have a
// slice of these. The values, once set are read only as they should never get overwritten.
type SearchTermData struct {
	data map[string]string
}

func (stdata *SearchTermData) GetKeys() (keys_returned []string) {
	if len(stdata.data) == 0 {
		return keys_returned
	}

	keys_returned = make([]string, len(stdata.data))

	keys_index := 0
	for key := range stdata.data {
		keys_returned[keys_index] = key
		keys_index++
	}

	return keys_returned
}

func (stdata *SearchTermData) GetValue(key string) (ret_data string, err error) {
	ret_data, success := stdata.data[key]
	if !success {
		err = fmt.Errorf("could not find key: %s in searchtermdata", key)
	}

	return ret_data, err
}

func (stdata *SearchTermData) AddValue(key, value string) (err error) {
	_, success := stdata.data[key]
	if success {
		err = fmt.Errorf("unable to set value: %s for key: %s - data is readonly", value, key)
	} else {
		stdata.data[key] = value
	}

	return err
}
