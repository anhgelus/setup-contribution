package utils

func RelativeToAbsolute(path string, contextPath string) string {
	if path == "." {
		return contextPath
	}
	switch string(path[0]) {
	case "/":
		return path
	case "~":
		return path
	case ".":
		return contextPath + path[2:]
	default:
		return contextPath + path
	}
}
