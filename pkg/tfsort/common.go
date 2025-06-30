package tfsort

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pkg/errors"
)

// ParseAndSortBlocks reads a Terraform file, parses it, extracts blocks of the given type,
// sorts them by name, and returns the sorted blocks as a string.
func ParseAndSortBlocks(path, blockType string) (string, error) {
	if ext := filepath.Ext(path); ext != ".tf" {
		return "", errors.New("not a terraform file")
	}

	src, err := os.ReadFile(path)
	if err != nil {
		return "", errors.Wrap(err, "failed to read file")
	}

	file, diags := hclsyntax.ParseConfig(src, path, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return "", errors.Wrap(diags, "failed to parse HCL")
	}

	body, ok := file.Body.(*hclsyntax.Body)
	if !ok {
		return "", errors.New("unexpected body type, expected *hclsyntax.Body")
	}

	var blocks []*hclsyntax.Block
	for _, block := range body.Blocks {
		if block.Type == blockType && len(block.Labels) > 0 {
			blocks = append(blocks, block)
		}
	}

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].Labels[0] < blocks[j].Labels[0]
	})

	var resultBuilder strings.Builder
	for _, block := range blocks {
		rng := block.Range()
		resultBuilder.WriteString(string(src[rng.Start.Byte:rng.End.Byte]))
		resultBuilder.WriteString("\n\n")
	}

	// Trim the trailing newline
	result := strings.TrimSuffix(resultBuilder.String(), "\n\n")
	return result + "\n", nil
}
