# RepoGraph Platform Architecture

## Table of Contents
- [Overview](#overview)
- [Architectural Patterns](#architectural-patterns)
- [System Design](#system-design)
- [Microservices](#microservices)
- [Data Flow](#data-flow)
- [Technology Stack](#technology-stack)
- [Design Decisions](#design-decisions)

---

## Overview

RepoGraph Platform is built using a **microservices architecture** with **SOLID principles** at its core. The system is designed for scalability, maintainability, and extensibility.

### Key Architectural Goals

1. **Separation of Concerns**: Each service has a single responsibility
2. **Scalability**: Services can be scaled independently
3. **Reliability**: Fault isolation and graceful degradation
4. **Maintainability**: Clean code with clear boundaries
5. **Extensibility**: Easy to add new file processors or features

---

## Architectural Patterns

### 1. Microservices Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         Client Layer                             │
│                    (CLI, Web UI, API)                            │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                      API Gateway / Router                        │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Orchestrator Service                          │
│              (Workflow Coordination)                             │
└───────────────────────────────┬─────────────────────────────────┘
                                │
        ┌───────────────────────┼───────────────────────┐
        │                       │                       │
        ▼                       ▼                       ▼
┌──────────────┐       ┌──────────────┐       ┌──────────────┐
│   Document   │       │   Content    │       │    Vision    │
│   Scanner    │───────│  Extractor   │───────│   Service    │
└──────────────┘       └──────────────┘       └──────────────┘
        │                       │                       │
        │                       ▼                       │
        │              ┌──────────────┐                │
        │              │Summarization │                │
        │              │   Service    │                │
        │              └──────────────┘                │
        │                       │                       │
        └───────────────────────┼───────────────────────┘
                                ▼
                       ┌──────────────┐
                       │  Embedding   │
                       │   Service    │
                       └──────┬───────┘
                              │
                              ▼
                       ┌──────────────┐
                       │ Vector Store │
                       │   Service    │
                       └──────┬───────┘
                              │
                              ▼
                       ┌──────────────┐
                       │    Query     │
                       │   Service    │
                       └──────────────┘
```

### 2. SOLID Principles Implementation

#### Single Responsibility Principle (SRP)
- **Document Scanner**: Only responsible for file discovery and metadata
- **Content Extractor**: Only handles content extraction
- **Vision Service**: Only analyzes images
- **Summarization Service**: Only generates summaries
- Each processor handles one file format category

#### Open/Closed Principle (OCP)
- New file processors can be added without modifying existing code
- Implement the `ContentExtractor` interface
- Register in the processor registry

#### Liskov Substitution Principle (LSP)
- All processors implement the same interface
- Any processor can be swapped without breaking the system
- All services follow their interface contracts

#### Interface Segregation Principle (ISP)
- Small, focused interfaces (DocumentScanner, VisionAnalyzer, etc.)
- Clients only depend on methods they use
- No "god interfaces"

#### Dependency Inversion Principle (DIP)
- High-level modules depend on abstractions (interfaces)
- Adapters for external services (Azure, Google, Pinecone)
- Dependency injection for all services

### 3. Domain-Driven Design (DDD)

```
internal/
├── domain/              # Core domain (business logic)
│   ├── models/         # Domain entities
│   └── interfaces/     # Service contracts
├── application/        # Application services
│   └── orchestrator/   # Workflow orchestration
└── infrastructure/     # External integrations
    ├── adapters/       # External service adapters
    └── repositories/   # Data persistence
```

---

## System Design

### Component Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                         RepoGraph Platform                       │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │                    Core Services                          │  │
│  │                                                            │  │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐ │  │
│  │  │Document  │  │Content   │  │ Vision   │  │Summarize │ │  │
│  │  │Scanner   │  │Extractor │  │ Service  │  │ Service  │ │  │
│  │  └────┬─────┘  └────┬─────┘  └────┬─────┘  └────┬─────┘ │  │
│  │       │             │              │             │        │  │
│  │  ┌────┴─────┐  ┌────┴─────┐  ┌────┴─────┐  ┌────┴─────┐ │  │
│  │  │Embedding │  │ Vector   │  │  Query   │  │Orchestrate│ │  │
│  │  │Service   │  │  Store   │  │ Service  │  │  Service  │ │  │
│  │  └──────────┘  └──────────┘  └──────────┘  └──────────┘ │  │
│  └──────────────────────────────────────────────────────────┘  │
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │                   Data Layer                              │  │
│  │                                                            │  │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐               │  │
│  │  │PostgreSQL│  │  Redis   │  │ Pinecone │               │  │
│  │  │(Metadata)│  │ (Cache)  │  │(Vectors) │               │  │
│  │  └──────────┘  └──────────┘  └──────────┘               │  │
│  └──────────────────────────────────────────────────────────┘  │
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │                External Services                          │  │
│  │                                                            │  │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐               │  │
│  │  │ Azure    │  │ Google   │  │ Pinecone │               │  │
│  │  │ OpenAI   │  │ Vision   │  │   API    │               │  │
│  │  └──────────┘  └──────────┘  └──────────┘               │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
```

### Deployment Architecture

```
┌──────────────────────────────────────────────────────────────┐
│                    Kubernetes Cluster                         │
├──────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌────────────────────────────────────────────────────────┐ │
│  │                   Ingress Controller                    │ │
│  │                  (NGINX/Traefik)                        │ │
│  └─────────────────────┬──────────────────────────────────┘ │
│                        │                                     │
│  ┌─────────────────────┴──────────────────────────────────┐ │
│  │                   Service Mesh                          │ │
│  │                    (Istio)                              │ │
│  └─────────────────────┬──────────────────────────────────┘ │
│                        │                                     │
│  ┌────────────────────────────────────────────────────────┐ │
│  │               Microservices Pods                        │ │
│  │                                                          │ │
│  │  [Orchestrator] [Scanner] [Extractor] [Vision]         │ │
│  │  [Summarize] [Embedding] [VectorStore] [Query]         │ │
│  │                                                          │ │
│  │  Each service: 2-3 replicas, auto-scaling              │ │
│  └────────────────────────────────────────────────────────┘ │
│                                                               │
│  ┌────────────────────────────────────────────────────────┐ │
│  │                  Data Persistence                       │ │
│  │                                                          │ │
│  │  [PostgreSQL StatefulSet]  [Redis Cluster]             │ │
│  └────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────────────────────┘
```

---

## Microservices

### 1. Document Scanner Service (Port 8081)

**Responsibility**: File discovery and metadata extraction

**Endpoints**:
- `POST /api/v1/scan/directory` - Scan a directory
- `GET /api/v1/metadata/:filePath` - Get file metadata
- `POST /api/v1/compute-hash` - Compute file hash

**Dependencies**: None (leaf service)

### 2. Content Extractor Service (Port 8082)

**Responsibility**: Extract text content from various file formats

**Endpoints**:
- `POST /api/v1/extract` - Extract content from a file
- `GET /api/v1/formats` - List supported formats

**Processors**:
- Image Processor (PNG, JPG, SVG, etc.)
- Diagram Processor (DrawIO, Excalidraw)
- Document Processor (DOCX, PDF, PPTX)
- Spreadsheet Processor (XLSX, XLS)
- Code Processor (Go, Python, JavaScript, etc.)
- Structured Data Processor (JSON, YAML, XML)

### 3. Vision Service (Port 8083)

**Responsibility**: Analyze images and diagrams using Google Vision API

**Endpoints**:
- `POST /api/v1/analyze/image` - Analyze an image
- `POST /api/v1/analyze/diagram` - Analyze a diagram
- `POST /api/v1/detect/text` - OCR text detection

**Dependencies**: Google Vision API

### 4. Summarization Service (Port 8084)

**Responsibility**: Generate summaries using Azure OpenAI

**Endpoints**:
- `POST /api/v1/summarize` - Generate summary
- `POST /api/v1/summarize/context` - Summarize with context

**Dependencies**: Azure OpenAI API

### 5. Embedding Service (Port 8085)

**Responsibility**: Generate embeddings using Azure OpenAI

**Endpoints**:
- `POST /api/v1/embed` - Generate single embedding
- `POST /api/v1/embed/batch` - Generate batch embeddings
- `GET /api/v1/dimension` - Get embedding dimension

**Dependencies**: Azure OpenAI API

### 6. Vector Store Service (Port 8086)

**Responsibility**: Manage vectors in Pinecone

**Endpoints**:
- `POST /api/v1/upsert` - Upsert vectors
- `POST /api/v1/search` - Search similar vectors
- `DELETE /api/v1/document/:id` - Delete by document ID
- `GET /api/v1/stats` - Get index statistics
- `GET /api/v1/exists/:hash` - Check if document exists

**Dependencies**: Pinecone API

### 7. Query Service (Port 8087)

**Responsibility**: Handle RAG queries

**Endpoints**:
- `POST /api/v1/query` - Execute RAG query
- `POST /api/v1/search` - Search documents

**Dependencies**: Embedding Service, Vector Store Service, Azure OpenAI

### 8. Orchestrator Service (Port 8088)

**Responsibility**: Coordinate the complete workflow

**Endpoints**:
- `POST /api/v1/process/document` - Process single document
- `POST /api/v1/process/directory` - Process directory
- `GET /api/v1/status/:documentId` - Get processing status

**Dependencies**: All other services

---

## Data Flow

### Document Indexing Flow

```
1. User → CLI/API
   └─ Request: Index directory

2. CLI → Orchestrator
   └─ POST /process/directory

3. Orchestrator → Document Scanner
   └─ Scan files, compute hashes
   └─ Return: List of file metadata

4. For each file:
   a. Orchestrator → Content Extractor
      └─ Extract text content
      
   b. If image/diagram → Vision Service
      └─ Analyze visual content
      
   c. Orchestrator → Summarization Service
      └─ Generate summary
      
   d. Orchestrator → Chunker
      └─ Split into chunks
      
   e. Orchestrator → Embedding Service
      └─ Generate embeddings
      
   f. Orchestrator → Vector Store
      └─ Upsert vectors to Pinecone

5. Orchestrator → Database
   └─ Store metadata

6. Orchestrator → User
   └─ Return: Processing complete
```

### Query Flow

```
1. User → CLI/API
   └─ Question: "What is the architecture?"

2. CLI → Query Service
   └─ POST /query

3. Query Service → Embedding Service
   └─ Generate query embedding

4. Query Service → Vector Store
   └─ Search similar vectors
   └─ Return: Top K results

5. Query Service → Azure OpenAI
   └─ Generate answer with context
   └─ Return: Answer + sources

6. Query Service → User
   └─ Display answer and sources
```

---

## Technology Stack

### Core Languages & Frameworks
- **Go 1.21**: Primary language
- **Gin**: HTTP web framework
- **Cobra**: CLI framework
- **Viper**: Configuration management
- **Zap**: Structured logging

### External Services
- **Azure OpenAI**: LLM and embeddings
- **Google Vision API**: Image analysis
- **Pinecone**: Vector database

### Data Storage
- **PostgreSQL**: Metadata storage
- **Redis**: Caching layer

### Infrastructure
- **Docker**: Containerization
- **Kubernetes**: Orchestration
- **GitHub Actions**: CI/CD

### Observability
- **OpenTelemetry**: Distributed tracing
- **Prometheus**: Metrics
- **Grafana**: Visualization

---

## Design Decisions

### 1. Why Microservices?

**Pros**:
- Independent scaling
- Technology flexibility
- Fault isolation
- Team autonomy

**Cons**:
- Increased complexity
- Network overhead
- Distributed debugging

**Decision**: Benefits outweigh costs for enterprise use

### 2. Why Go?

- High performance
- Excellent concurrency (goroutines)
- Strong standard library
- Fast compilation
- Single binary deployment
- Great for microservices

### 3. Why Azure OpenAI over OpenAI?

- Enterprise SLA
- Data privacy guarantees
- Regional deployment
- Better compliance

### 4. Why Pinecone?

- Managed service (no ops overhead)
- High performance
- Good free tier
- Strong community

### 5. Synchronous vs Asynchronous

**Current**: Synchronous HTTP
**Future**: Message queue (NATS/RabbitMQ) for long-running tasks

**Decision**: Start simple, add complexity as needed

---

## Future Enhancements

1. **gRPC Communication**: Replace HTTP with gRPC for better performance
2. **Event-Driven Architecture**: Use message queues for async workflows
3. **GraphQL API**: Unified query interface
4. **Caching Strategy**: Multi-level caching (Redis, in-memory)
5. **Batch Processing**: Parallel document processing
6. **A/B Testing**: Compare different LLM models
7. **Multi-tenancy**: Support multiple organizations

---

*Last Updated: February 2026*
