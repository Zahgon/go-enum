package generator

// Stringify returns a string that is all of the enum value names concatenated without a separator
func Stringify(e Enum, forceLower, forceUpper bool) (ret string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Mapify returns a map that is all of the indexes for a string value lookup
func Mapify(e Enum) (ret string, err error) { _ = "STUB: not implemented"; return "", nil }

// Unmapify returns a map that is all of the indexes for a string value lookup
func Unmapify(e Enum, lowercase bool) (ret string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Unmapify returns a map that is all of the indexes for a string value lookup
func UnmapifyStringEnum(e Enum, lowercase bool) (ret string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Namify returns a slice that is all of the possible names for an enum in a slice
func Namify(e Enum) (ret string, err error) { _ = "STUB: not implemented"; return "", nil }

// Namify returns a slice that is all of the possible names for an enum in a slice
func namifyStringEnum(e Enum) (ret string, err error) { _ = "STUB: not implemented"; return "", nil }

func Offset(index int, enumType string, val EnumValue) (strResult string) {
	_ = "STUB: not implemented"
	return ""
}

// Unsigned

// Signed

// DirectValue returns the exact value of the enum, not adjusted for iota at all.
func DirectValue(enumType string, val EnumValue) (strResult string) {
	_ = "STUB: not implemented"
	return ""
}

// Unsigned

// Signed
