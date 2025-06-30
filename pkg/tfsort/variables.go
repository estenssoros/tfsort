package tfsort

func Variables(path string) (string, error) {
	return ParseAndSortBlocks(path, "variable")
}
