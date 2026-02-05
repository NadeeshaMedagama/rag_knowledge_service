# Architecture Documentation

This folder contains documentation about the system architecture and design patterns used in the RAG Knowledge Service.

## 📄 Documents

- **[ARCHITECTURE.md](./ARCHITECTURE.md)** - Comprehensive system architecture overview
- **[PINECONE_ONLY_ARCHITECTURE.md](./PINECONE_ONLY_ARCHITECTURE.md)** - Pinecone vector database specific architecture
- **[AUTOMATIC_INDEXING.md](./AUTOMATIC_INDEXING.md)** - Automatic document indexing design and implementation

## 🏗️ Key Concepts

### Microservices Architecture
The system is built using a microservices architecture pattern with the following services:
- Orchestrator
- Document Scanner
- Content Extractor
- Vision Service
- Summarization Service
- Embedding Service
- Vector Store
- Query Service

### Design Principles
- SOLID principles
- Domain-driven design
- Separation of concerns
- Scalability and maintainability

## 📚 Related Documentation

- [API Reference](../api/) - Service APIs and endpoints
- [Deployment Guide](../deployment/) - How to deploy the architecture
- [Development Guide](../development/) - How to extend the architecture
