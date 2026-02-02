package processors

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

// ProcessorInterface defines the interface for content processors
type ProcessorInterface interface {
	CanProcess(fileType string) bool
	Extract(ctx context.Context, filePath string) (string, error)
}

// TextProcessor handles plain text files
type TextProcessor struct {
	logger *zap.Logger
}

// NewTextProcessor creates a new text processor
func NewTextProcessor(logger *zap.Logger) *TextProcessor {
	return &TextProcessor{logger: logger}
}

// CanProcess checks if this processor can handle the file type
func (p *TextProcessor) CanProcess(fileType string) bool {
	textTypes := []string{".txt", ".md", ".log", ".csv", ".json", ".yaml", ".yml", ".xml", ".toml"}
	for _, t := range textTypes {
		if strings.EqualFold(fileType, t) {
			return true
		}
	}
	return false
}

// Extract extracts content from text files
func (p *TextProcessor) Extract(ctx context.Context, filePath string) (string, error) {
	p.logger.Debug("Extracting text content", zap.String("file", filePath))

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}

// ImageProcessor handles image files
type ImageProcessor struct {
	logger *zap.Logger
}

// NewImageProcessor creates a new image processor
func NewImageProcessor(logger *zap.Logger) *ImageProcessor {
	return &ImageProcessor{logger: logger}
}

// CanProcess checks if this processor can handle the file type
func (p *ImageProcessor) CanProcess(fileType string) bool {
	imageTypes := []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".svg", ".webp"}
	for _, t := range imageTypes {
		if strings.EqualFold(fileType, t) {
			return true
		}
	}
	return false
}

// Extract extracts content from image files (returns path for vision API)
func (p *ImageProcessor) Extract(ctx context.Context, filePath string) (string, error) {
	p.logger.Debug("Image file detected", zap.String("file", filePath))
	// Return file path - actual analysis will be done by Vision service
	return fmt.Sprintf("[IMAGE FILE: %s]", filepath.Base(filePath)), nil
}

// DocumentProcessor handles PDF and DOCX files
type DocumentProcessor struct {
	logger *zap.Logger
}

// NewDocumentProcessor creates a new document processor
func NewDocumentProcessor(logger *zap.Logger) *DocumentProcessor {
	return &DocumentProcessor{logger: logger}
}

// CanProcess checks if this processor can handle the file type
func (p *DocumentProcessor) CanProcess(fileType string) bool {
	docTypes := []string{".pdf", ".docx", ".doc", ".pptx", ".ppt", ".odt"}
	for _, t := range docTypes {
		if strings.EqualFold(fileType, t) {
			return true
		}
	}
	return false
}

// Extract extracts content from documents
func (p *DocumentProcessor) Extract(ctx context.Context, filePath string) (string, error) {
	p.logger.Debug("Extracting document content", zap.String("file", filePath))

	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".pdf":
		return p.extractPDF(filePath)
	case ".docx":
		return p.extractDOCX(filePath)
	case ".pptx":
		return p.extractPPTX(filePath)
	default:
		return "", fmt.Errorf("unsupported document type: %s", ext)
	}
}

func (p *DocumentProcessor) extractPDF(filePath string) (string, error) {
	// For now, return placeholder
	// In production, use a library like github.com/ledongthuc/pdf
	return fmt.Sprintf("[PDF Document: %s]\n(PDF extraction not yet implemented - placeholder)", filepath.Base(filePath)), nil
}

func (p *DocumentProcessor) extractDOCX(filePath string) (string, error) {
	// For now, return placeholder
	// In production, use a library like github.com/nguyenthenguyen/docx
	return fmt.Sprintf("[DOCX Document: %s]\n(DOCX extraction not yet implemented - placeholder)", filepath.Base(filePath)), nil
}

func (p *DocumentProcessor) extractPPTX(filePath string) (string, error) {
	return fmt.Sprintf("[PPTX Presentation: %s]\n(PPTX extraction not yet implemented - placeholder)", filepath.Base(filePath)), nil
}

// SpreadsheetProcessor handles XLSX and CSV files
type SpreadsheetProcessor struct {
	logger *zap.Logger
}

// NewSpreadsheetProcessor creates a new spreadsheet processor
func NewSpreadsheetProcessor(logger *zap.Logger) *SpreadsheetProcessor {
	return &SpreadsheetProcessor{logger: logger}
}

// CanProcess checks if this processor can handle the file type
func (p *SpreadsheetProcessor) CanProcess(fileType string) bool {
	spreadsheetTypes := []string{".xlsx", ".xls", ".csv"}
	for _, t := range spreadsheetTypes {
		if strings.EqualFold(fileType, t) {
			return true
		}
	}
	return false
}

// Extract extracts content from spreadsheets
func (p *SpreadsheetProcessor) Extract(ctx context.Context, filePath string) (string, error) {
	p.logger.Debug("Extracting spreadsheet content", zap.String("file", filePath))

	ext := strings.ToLower(filepath.Ext(filePath))

	if ext == ".csv" {
		return p.extractCSV(filePath)
	}

	// For XLSX, return placeholder
	return fmt.Sprintf("[Spreadsheet: %s]\n(XLSX extraction not yet implemented - placeholder)", filepath.Base(filePath)), nil
}

func (p *SpreadsheetProcessor) extractCSV(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read CSV: %w", err)
	}
	return string(content), nil
}

// CodeProcessor handles source code files
type CodeProcessor struct {
	logger *zap.Logger
}

// NewCodeProcessor creates a new code processor
func NewCodeProcessor(logger *zap.Logger) *CodeProcessor {
	return &CodeProcessor{logger: logger}
}

// CanProcess checks if this processor can handle the file type
func (p *CodeProcessor) CanProcess(fileType string) bool {
	codeTypes := []string{".go", ".py", ".js", ".ts", ".java", ".c", ".cpp", ".h", ".rs", ".rb", ".php", ".sql"}
	for _, t := range codeTypes {
		if strings.EqualFold(fileType, t) {
			return true
		}
	}
	return false
}

// Extract extracts content from code files
func (p *CodeProcessor) Extract(ctx context.Context, filePath string) (string, error) {
	p.logger.Debug("Extracting code content", zap.String("file", filePath))

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read code file: %w", err)
	}

	return string(content), nil
}
