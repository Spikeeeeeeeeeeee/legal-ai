package services

import (
	"fmt"

	"github.com/yourusername/legal-ai-backend/utils"
)

func SummarizeLegalText(text string) (string, error) {
	utils.LogInfo("ai", "mock summarization called")
	return mockSummarize(text), nil
}

func mockSummarize(original string) string {
	return fmt.Sprintf("Plain-English summary (mock): This clause says you cannot alter the premises without the landlord's written permission.")
}
