package converter

func StringToBool(s string) bool {
	var b (bool)
	switch s {
	case "true":
		b = true
	case "false":
		b = false
	}
	return b
}
