package models

import (
	"time"

	"github.com/google/uuid"
)

// Document represents a processed document in the system
type Document struct {
	ID              uuid.UUID         `json:"id"`
	FileName        string            `json:"file_name"`
	FilePath        string            `json:"file_path"`
	FileType        string            `json:"file_type"`
	FileSize        int64             `json:"file_size"`
	FileHash        string            `json:"file_hash"` // SHA256 hash for deduplication
	Content         string            `json:"content"`
	RawContent      []byte            `json:"raw_content,omitempty"`
	Metadata        map[string]string `json:"metadata"`
	Summary         string            `json:"summary,omitempty"`
	VisionAnalysis  string            `json:"vision_analysis,omitempty"`
	ProcessingState ProcessingState   `json:"processing_state"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	IndexedAt       *time.Time        `json:"indexed_at,omitempty"`
	Chunks          []Chunk           `json:"chunks,omitempty"`
}

// ProcessingState represents the current state of document processing
type ProcessingState string

const (
	StateScanned    ProcessingState = "SCANNED"
	StateExtracted  ProcessingState = "EXTRACTED"
	StateAnalyzed   ProcessingState = "ANALYZED"
	StateSummarized ProcessingState = "SUMMARIZED"
	StateChunked    ProcessingState = "CHUNKED"
	StateEmbedded   ProcessingState = "EMBEDDED"
	StateIndexed    ProcessingState = "INDEXED"
	StateFailed     ProcessingState = "FAILED"
)

// Chunk represents a text chunk from a document
type Chunk struct {
	ID         uuid.UUID         `json:"id"`
	DocumentID uuid.UUID         `json:"document_id"`
	Content    string            `json:"content"`
	StartIndex int               `json:"start_index"`
	EndIndex   int               `json:"end_index"`
	ChunkIndex int               `json:"chunk_index"`
	Metadata   map[string]string `json:"metadata"`
	Embedding  []float32         `json:"embedding,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
}

// FileMetadata contains metadata about a scanned file
type FileMetadata struct {
	Path         string            `json:"path"`
	Name         string            `json:"name"`
	Extension    string            `json:"extension"`
	Size         int64             `json:"size"`
	ModifiedTime time.Time         `json:"modified_time"`
	Hash         string            `json:"hash"`
	MimeType     string            `json:"mime_type"`
	IsDirectory  bool              `json:"is_directory"`
	Custom       map[string]string `json:"custom,omitempty"`
}

// NewDocument creates a new document with generated ID and timestamps
func NewDocument(fileName, filePath, fileType string, fileSize int64, fileHash string) *Document {
	now := time.Now()
	return &Document{
		ID:              uuid.New(),
		FileName:        fileName,
		FilePath:        filePath,
		FileType:        fileType,
		FileSize:        fileSize,
		FileHash:        fileHash,
		Metadata:        make(map[string]string),
		ProcessingState: StateScanned,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

// NewChunk creates a new chunk with generated ID and timestamp
func NewChunk(documentID uuid.UUID, content string, startIdx, endIdx, chunkIdx int) *Chunk {
	return &Chunk{
		ID:         uuid.New(),
		DocumentID: documentID,
		Content:    content,
		StartIndex: startIdx,
		EndIndex:   endIdx,
		ChunkIndex: chunkIdx,
		Metadata:   make(map[string]string),
		CreatedAt:  time.Now(),
	}
}
