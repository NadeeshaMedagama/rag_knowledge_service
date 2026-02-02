package models

import (
	"time"

	"github.com/google/uuid"
)

// Query represents a user query in the RAG system
type Query struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	TopK      int       `json:"top_k"`
	Namespace string    `json:"namespace,omitempty"`
	Filter    Filter    `json:"filter,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// Filter represents query filters
type Filter struct {
	FileType string            `json:"file_type,omitempty"`
	DateFrom *time.Time        `json:"date_from,omitempty"`
	DateTo   *time.Time        `json:"date_to,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// QueryResult represents the result of a RAG query
type QueryResult struct {
	QueryID   uuid.UUID      `json:"query_id"`
	Answer    string         `json:"answer"`
	Sources   []SearchResult `json:"sources"`
	Timestamp time.Time      `json:"timestamp"`
}

// SearchResult represents a single search result from vector store
type SearchResult struct {
	DocumentID uuid.UUID         `json:"document_id"`
	ChunkID    uuid.UUID         `json:"chunk_id"`
	Score      float32           `json:"score"`
	Content    string            `json:"content"`
	FileName   string            `json:"file_name"`
	FilePath   string            `json:"file_path"`
	FileType   string            `json:"file_type"`
	Metadata   map[string]string `json:"metadata"`
}

// NewQuery creates a new query with generated ID and timestamp
func NewQuery(text string, topK int) *Query {
	return &Query{
		ID:        uuid.New(),
		Text:      text,
		TopK:      topK,
		CreatedAt: time.Now(),
	}
}
