package commands

func Judge(names []string) string {
	result := ""
	for _, name := range names {
		result += name + "\n"
	}

	return result
}
