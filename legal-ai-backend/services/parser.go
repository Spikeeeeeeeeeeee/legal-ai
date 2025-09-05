package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yourusername/legal-ai-backend/utils"
)

func ParseDocument(path string) (string, error) {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".txt":
		b, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case ".pdf":
		utils.LogInfo("parser", "PDF file detected — using mock extractor")
		return mockExtract(path)
	case ".docx", ".doc":
		utils.LogInfo("parser", "DOCX/DOC file detected — using mock extractor")
		return mockExtract(path)
	default:
		return "", fmt.Errorf("unsupported file type: %s", ext)
	}
}

func mockExtract(path string) (string, error) {
	name := filepath.Base(path)
	text := fmt.Sprintf("Extracted text from %s.\n\n[This is placeholder content.]", name)
	return text, nil
}
