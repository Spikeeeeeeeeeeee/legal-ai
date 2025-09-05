package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/yourusername/legal-ai-backend/models"
	"github.com/yourusername/legal-ai-backend/services"
	"github.com/yourusername/legal-ai-backend/utils"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		utils.LogError("parse form", err)
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		utils.LogError("formfile", err)
		http.Error(w, "file missing", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := os.MkdirAll("uploads", 0755); err != nil {
		utils.LogError("mkdir upload", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	id := uuid.New().String()
	ext := filepath.Ext(header.Filename)
	outPath := filepath.Join("uploads", fmt.Sprintf("%s%s", id, ext))

	out, err := os.Create(outPath)
	if err != nil {
		utils.LogError("create file", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	if _, err := io.Copy(out, file); err != nil {
		utils.LogError("save file", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	utils.LogInfo("file saved", outPath)

	originalText, perr := services.ParseDocument(outPath)
	if perr != nil {
		utils.LogError("parse document", perr)
		http.Error(w, "failed to parse document", http.StatusInternalServerError)
		return
	}

	summary, aerr := services.SummarizeLegalText(originalText)
	if aerr != nil {
		utils.LogError("ai summarize", aerr)
		http.Error(w, "AI processing error", http.StatusInternalServerError)
		return
	}

	doc := models.Document{
		ID:           id,
		Filename:     header.Filename,
		UploadedAt:   time.Now().UTC(),
		OriginalText: originalText,
		Summary:      summary,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(doc)
}
