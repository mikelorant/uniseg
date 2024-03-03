package uniseg

var GraphemeClusterWidthOverrides map[string]int

func overrideWidth(gc string, w int) int {
	if len(GraphemeClusterWidthOverrides) == 0 {
		return w
	}

	if width, ok := GraphemeClusterWidthOverrides[gc]; ok {
		return width
	}

	return w
}
