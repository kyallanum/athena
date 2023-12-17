package utils

func AddDictToLibrary(ruleDict *map[string](map[string]string), ruleLibrary *map[string](map[string][]string)) {

	for currentRule, value := range *ruleDict {
		if (*ruleLibrary)[currentRule] == nil {
			(*ruleLibrary)[currentRule] = make(map[string][]string)
		}

		for currentKey, value := range value {
			if (*ruleLibrary)[currentRule][currentKey] == nil {
				(*ruleLibrary)[currentRule][currentKey] = make([]string, 0)
			}
			(*ruleLibrary)[currentRule][currentKey] = append((*ruleLibrary)[currentRule][currentKey], value)
		}
	}
}
