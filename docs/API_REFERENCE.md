# API Reference

## Table of Contents
- [Authentication](#authentication)
- [Orchestrator Service](#orchestrator-service)
- [Document Scanner Service](#document-scanner-service)
- [Content Extractor Service](#content-extractor-service)
- [Vision Service](#vision-service)
- [Summarization Service](#summarization-service)
- [Embedding Service](#embedding-service)
- [Vector Store Service](#vector-store-service)
- [Query Service](#query-service)
- [Error Codes](#error-codes)

---

## Authentication

Currently, services use internal authentication. For production:
- Use API keys in headers: `X-API-Key: your-api-key`
- JWT tokens for user authentication
- Service-to-service mTLS

---

## Orchestrator Service

**Base URL**: `http://localhost:8088`

### Health Check

```http
GET /health
```

**Response**:
```json
{
  "healthy": true,
  "services": {
    "azure_openai": true,
    "pinecone": true,
    "google_vision": true,
    "database": true,
    "redis": true
  },
  "timestamp": "2026-02-02T10:00:00Z"
}
```

### Process Single Document

```http
POST /api/v1/process/document
Content-Type: application/json

{
  "file_path": "/path/to/document.pdf"
}
```

**Response**:
```json
{
  "status": "accepted",
  "document_id": "123e4567-e89b-12d3-a456-426614174000",
  "message": "Document processing started"
}
```

### Process Directory

```http
POST /api/v1/process/directory
Content-Type: application/json

{
  "directory": "/path/to/documents",
  "force_reprocess": false
}
```

**Response**:
```json
{
  "status": "accepted",
  "job_id": "abc123",
  "message": "Directory processing started"
}
```

### Get Processing Status

```http
GET /api/v1/status/:documentId
```

**Response**:
```json
{
  "document_id": "123e4567-e89b-12d3-a456-426614174000",
  "status": "processing",
  "progress": 60,
  "current_stage": "embedding",
  "created_at": "2026-02-02T10:00:00Z",
  "updated_at": "2026-02-02T10:05:00Z"
}
```

**Status Values**:
- `scanned`: File scanned
- `extracted`: Content extracted
- `analyzed`: Vision analysis complete
- `summarized`: Summary generated
- `chunked`: Text chunked
- `embedded`: Embeddings generated
- `indexed`: Stored in vector database
- `completed`: Fully processed
- `failed`: Processing failed

---

## Document Scanner Service

**Base URL**: `http://localhost:8081`

### Scan Directory

```http
POST /api/v1/scan/directory
Content-Type: application/json

{
  "directory": "/path/to/documents",
  "recursive": true,
  "file_types": ["pdf", "docx", "png"]
}
```

**Response**:
```json
{
  "directory": "/path/to/documents",
  "files": [
    {
      "path": "/path/to/doc1.pdf",
      "name": "doc1.pdf",
      "extension": "pdf",
      "size": 1024000,
      "modified_time": "2026-02-01T10:00:00Z",
      "hash": "abc123...",
      "mime_type": "application/pdf"
    }
  ],
  "total_files": 10,
  "total_size": 10240000
}
```

### Get File Metadata

```http
GET /api/v1/metadata/:filePath
```

**Response**:
```json
{
  "path": "/path/to/doc.pdf",
  "name": "doc.pdf",
  "extension": "pdf",
  "size": 1024000,
  "category": "document",
  "mime_type": "application/pdf",
  "hash": "abc123..."
}
```

### Compute File Hash

```http
POST /api/v1/compute-hash
Content-Type: application/json

{
  "file_path": "/path/to/document.pdf"
}
```

**Response**:
```json
{
  "file_path": "/path/to/document.pdf",
  "hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
}
```

---

## Content Extractor Service

**Base URL**: `http://localhost:8082`

### Extract Content

```http
POST /api/v1/extract
Content-Type: application/json

{
  "file_path": "/path/to/document.pdf",
  "file_type": "pdf",
  "options": {
    "include_metadata": true,
    "extract_images": false
  }
}
```

**Response**:
```json
{
  "content": "Extracted text content...",
  "metadata": {
    "author": "John Doe",
    "created_at": "2026-01-01T00:00:00Z",
    "page_count": 10
  },
  "extracted_at": "2026-02-02T10:00:00Z"
}
```

### List Supported Formats

```http
GET /api/v1/formats
```

**Response**:
```json
{
  "formats": [
    {
      "category": "image",
      "extensions": ["png", "jpg", "jpeg", "svg"]
    },
    {
      "category": "document",
      "extensions": ["pdf", "docx", "pptx"]
    }
  ]
}
```

---

## Vision Service

**Base URL**: `http://localhost:8083`

### Analyze Image

```http
POST /api/v1/analyze/image
Content-Type: multipart/form-data

file: <binary>
```

**Response**:
```json
{
  "description": "An architectural diagram showing microservices...",
  "labels": ["architecture", "diagram", "microservices"],
  "text_annotations": ["Service A", "Service B"],
  "confidence": 0.95
}
```

### Analyze Diagram

```http
POST /api/v1/analyze/diagram
Content-Type: multipart/form-data

file: <binary>
```

**Response**:
```json
{
  "description": "System architecture diagram with 5 components...",
  "components": ["API Gateway", "Service Mesh", "Database"],
  "relationships": [
    {"from": "API Gateway", "to": "Service Mesh"}
  ]
}
```

### Detect Text (OCR)

```http
POST /api/v1/detect/text
Content-Type: multipart/form-data

file: <binary>
```

**Response**:
```json
{
  "text": "Detected text from image...",
  "blocks": [
    {
      "text": "Header Text",
      "confidence": 0.98,
      "bounding_box": [10, 20, 100, 50]
    }
  ]
}
```

---

## Summarization Service

**Base URL**: `http://localhost:8084`

### Generate Summary

```http
POST /api/v1/summarize
Content-Type: application/json

{
  "content": "Long text to summarize...",
  "max_tokens": 500,
  "style": "concise"
}
```

**Response**:
```json
{
  "summary": "This document describes...",
  "token_count": 150,
  "model": "gpt-4"
}
```

### Summarize with Context

```http
POST /api/v1/summarize/context
Content-Type: application/json

{
  "content": "Text to summarize...",
  "context": "This is an architecture document",
  "max_tokens": 300
}
```

**Response**:
```json
{
  "summary": "Architecture overview...",
  "token_count": 120
}
```

---

## Embedding Service

**Base URL**: `http://localhost:8085`

### Generate Embedding

```http
POST /api/v1/embed
Content-Type: application/json

{
  "text": "Text to embed"
}
```

**Response**:
```json
{
  "embedding": [0.123, -0.456, 0.789, ...],
  "dimension": 1536,
  "model": "text-embedding-ada-002"
}
```

### Generate Batch Embeddings

```http
POST /api/v1/embed/batch
Content-Type: application/json

{
  "texts": [
    "Text 1",
    "Text 2",
    "Text 3"
  ]
}
```

**Response**:
```json
{
  "embeddings": [
    [0.123, -0.456, ...],
    [0.789, 0.234, ...],
    [-0.567, 0.890, ...]
  ],
  "count": 3
}
```

### Get Dimension

```http
GET /api/v1/dimension
```

**Response**:
```json
{
  "dimension": 1536
}
```

---

## Vector Store Service

**Base URL**: `http://localhost:8086`

### Upsert Vectors

```http
POST /api/v1/upsert
Content-Type: application/json

{
  "vectors": [
    {
      "id": "chunk-1",
      "values": [0.123, -0.456, ...],
      "metadata": {
        "document_id": "doc-123",
        "file_name": "document.pdf",
        "chunk_index": 0
      }
    }
  ],
  "namespace": "default"
}
```

**Response**:
```json
{
  "upserted_count": 1
}
```

### Search Similar Vectors

```http
POST /api/v1/search
Content-Type: application/json

{
  "query_vector": [0.123, -0.456, ...],
  "top_k": 5,
  "namespace": "default",
  "filter": {
    "file_type": "pdf"
  }
}
```

**Response**:
```json
{
  "results": [
    {
      "id": "chunk-1",
      "score": 0.95,
      "metadata": {
        "document_id": "doc-123",
        "file_name": "document.pdf",
        "content": "Relevant text..."
      }
    }
  ]
}
```

### Delete by Document ID

```http
DELETE /api/v1/document/:id
```

**Response**:
```json
{
  "deleted": true,
  "count": 10
}
```

### Get Index Statistics

```http
GET /api/v1/stats
```

**Response**:
```json
{
  "dimension": 1536,
  "index_fullness": 0.25,
  "total_vector_count": 10000,
  "namespaces": {
    "default": {
      "vector_count": 10000
    }
  }
}
```

### Check Document Exists

```http
GET /api/v1/exists/:hash
```

**Response**:
```json
{
  "exists": true,
  "document_id": "doc-123"
}
```

---

## Query Service

**Base URL**: `http://localhost:8087`

### Execute RAG Query

```http
POST /api/v1/query
Content-Type: application/json

{
  "text": "What is the system architecture?",
  "top_k": 5,
  "namespace": "default",
  "filter": {
    "file_type": "pdf"
  }
}
```

**Response**:
```json
{
  "query_id": "query-123",
  "answer": "The system architecture consists of...",
  "sources": [
    {
      "document_id": "doc-123",
      "file_name": "architecture.pdf",
      "content": "Relevant excerpt...",
      "score": 0.95
    }
  ],
  "timestamp": "2026-02-02T10:00:00Z"
}
```

### Search Documents

```http
POST /api/v1/search
Content-Type: application/json

{
  "text": "authentication",
  "top_k": 10,
  "filter": {
    "file_type": "md"
  }
}
```

**Response**:
```json
{
  "results": [
    {
      "document_id": "doc-456",
      "file_name": "auth.md",
      "content": "Authentication is handled by...",
      "score": 0.92
    }
  ],
  "total": 10
}
```

---

## Error Codes

### Standard HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | Success |
| 201 | Created |
| 202 | Accepted (async processing) |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 429 | Too Many Requests |
| 500 | Internal Server Error |
| 503 | Service Unavailable |

### Application Error Codes

```json
{
  "error": {
    "code": "DOCUMENT_NOT_FOUND",
    "message": "Document with ID xyz not found",
    "details": {
      "document_id": "xyz"
    }
  }
}
```

**Error Codes**:
- `DOCUMENT_NOT_FOUND`: Document doesn't exist
- `INVALID_FILE_TYPE`: Unsupported file format
- `EXTRACTION_FAILED`: Content extraction failed
- `EMBEDDING_FAILED`: Embedding generation failed
- `VECTOR_STORE_ERROR`: Pinecone error
- `EXTERNAL_SERVICE_ERROR`: Azure/Google API error
- `RATE_LIMIT_EXCEEDED`: Too many requests

---

## Rate Limits

- **Per Service**: 100 requests/minute
- **Embedding Service**: 50 requests/minute
- **Query Service**: 30 requests/minute

---

## Pagination

For list endpoints:

```http
GET /api/v1/documents?page=1&per_page=50
```

**Response**:
```json
{
  "data": [...],
  "pagination": {
    "page": 1,
    "per_page": 50,
    "total": 500,
    "total_pages": 10
  }
}
```

---

*Last Updated: February 2026*
