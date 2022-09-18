package form

type errors map[string][]string

// Add error message for a given form field.
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns first error message.
func (e errors) Get(field string) string {
	s := e[field]

	if len(s) == 0 {
		return ""
	}

	return s[0]
}
