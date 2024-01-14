package models

import "fmt"

// The Library struct is a map of Rule Data. With keys that
// relate to the Rule Name. Each map K:V is read only once set
// All data that is committed to this library is readonly by default.
type Library struct {
	rule_data_collection map[string]RuleData
	name                 string
}

func New(name string) *Library {
	return &Library{
		rule_data_collection: make(map[string]RuleData),
		name:                 name,
	}
}

func (library *Library) Name() (string, error) {
	if library.name == "" {
		return "", fmt.Errorf("Library does not have a name assigned")
	}
	return library.name, nil
}

func (library *Library) LibraryKeys() (rule_data_keys []string) {
	if len(library.rule_data_collection) == 0 {
		return rule_data_keys
	}

	rule_data_keys = make([]string, len(library.rule_data_collection))

	rule_data_index := 0
	for key := range library.rule_data_collection {
		rule_data_keys[rule_data_index] = key
		rule_data_index++
	}

	return rule_data_keys
}

func (library *Library) RuleData(key string) (ret_rule_data RuleData, err error) {
	ret_rule_data, success := library.rule_data_collection[key]
	if !success {
		err = fmt.Errorf("could not find key: %s in library", key)
	}
	return ret_rule_data, err
}

func (library *Library) AddRuleData(key string, value *RuleData) (err error) {
	_, success := library.rule_data_collection[key]
	if success {
		err = fmt.Errorf("could not set ruledata for key: %s - data is readonly", key)
	} else {
		library.rule_data_collection[key] = *value
	}

	return err
}
