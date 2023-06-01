package tool

func ModeNameToMode(modeName string) (mode int) {
	switch modeName {
	case "normal":
		mode = 1
	case "hard":
		mode = 2
	default:
		mode = 1
	}
	return mode
}

func ModeToModeName(mode string) (modeName string) {
	switch mode {
	case "1", "normal":
		modeName = "普通"
	case "2", "hard":
		modeName = "困难"
	default:
		modeName = "普通"
	}
	return modeName
}
