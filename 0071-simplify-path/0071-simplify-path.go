func simplifyPath(path string) string {
	stack := []string{}

	parts := strings.Split(path, "/")

	for _, part := range parts {
		switch part {
		case "", ".":
			
			continue
		case "..":
			
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}