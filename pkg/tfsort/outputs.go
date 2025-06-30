package tfsort

func Outputs(path string) (string, error) {
	return ParseAndSortBlocks(path, "output")
}
