package utils

func RemoveDuplicates(suggestions []map[string]string) []map[string]string {
	seenLabels := make(map[string]struct{})
	uniqueSuggestions := []map[string]string{}

	for _, suggestion := range suggestions {
		for key := range suggestion {
			if _, ok := seenLabels[key]; !ok {
				seenLabels[key] = struct{}{}
				uniqueSuggestions = append(uniqueSuggestions, suggestion)
			}
		}
	}

	return uniqueSuggestions
}
