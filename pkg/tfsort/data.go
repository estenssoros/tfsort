package tfsort

func Data(path string) (string, error) {
	return ParseAndSortBlocks(path, "data")
}
