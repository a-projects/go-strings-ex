package strings_ex

// EmptyToDefault returns default value if string s is empty
func EmptyToDefault(s string, d string) string {
	if IsEmpty(s) {
		return d
	}
	return s
}

// IsEmpty reports whether string s empty or nil
func IsEmpty(s string) bool {
	return s == ""
}

// EmptyOrNilToDefault returns default value if string s is empty or nil
func EmptyOrNilToDefault(s *string, d string) string {
	if IsEmptyOrNil(s) {
		return d
	}
	return *s
}

// IsEmptyOrNil reports whether string s empty or nil
func IsEmptyOrNil(s *string) bool {
	return s == nil || *s == ""
}

// ToPtr converts a string literal to a string pointer
func ToPtr(s string) *string {
	return &s
}
