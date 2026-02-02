#!/bin/bash

# RepoGraph Platform - Project Verification Script
# Verifies that the project structure and files are correctly set up

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

PASS=0
FAIL=0

check_file() {
    if [ -f "$1" ]; then
        echo -e "${GREEN}✅ $1${NC}"
        ((PASS++))
    else
        echo -e "${RED}❌ $1 (missing)${NC}"
        ((FAIL++))
    fi
}

check_dir() {
    if [ -d "$1" ]; then
        echo -e "${GREEN}✅ $1/${NC}"
        ((PASS++))
    else
        echo -e "${RED}❌ $1/ (missing)${NC}"
        ((FAIL++))
    fi
}

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}  RepoGraph Platform - Project Verification${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""

echo -e "${YELLOW}📁 Checking Directory Structure${NC}"
check_dir "cmd/orchestrator"
check_dir "cmd/document-scanner"
check_dir "cmd/content-extractor"
check_dir "cmd/vision-service"
check_dir "cmd/summarization-service"
check_dir "cmd/embedding-service"
check_dir "cmd/vector-store"
check_dir "cmd/query-service"
check_dir "cmd/repograph-cli"
check_dir "internal/domain/models"
check_dir "internal/domain/interfaces"
check_dir "internal/config"
check_dir "internal/logger"
check_dir "pkg/utils"
check_dir "pkg/health"
check_dir "docs"
check_dir "deployments/docker"
check_dir "scripts"
check_dir ".github/workflows"
echo ""

echo -e "${YELLOW}📄 Checking Core Files${NC}"
check_file "go.mod"
check_file "go.sum"
check_file "Makefile"
check_file "README.md"
check_file "PROJECT_SUMMARY.md"
check_file ".gitignore"
check_file ".env.example"
check_file ".golangci.yml"
echo ""

echo -e "${YELLOW}🔧 Checking Domain Layer${NC}"
check_file "internal/domain/models/document.go"
check_file "internal/domain/models/query.go"
check_file "internal/domain/interfaces/services.go"
check_file "internal/config/config.go"
check_file "internal/logger/logger.go"
check_file "pkg/utils/file_utils.go"
check_file "pkg/health/health.go"
echo ""

echo -e "${YELLOW}🎯 Checking Service Entry Points${NC}"
check_file "cmd/orchestrator/main.go"
check_file "cmd/document-scanner/main.go"
check_file "cmd/repograph-cli/main.go"
echo ""

echo -e "${YELLOW}🐳 Checking Docker Files${NC}"
check_file "deployments/docker/docker-compose.yml"
check_file "deployments/docker/Dockerfile.orchestrator"
check_file "deployments/docker/Dockerfile.document-scanner"
check_file "deployments/docker/Dockerfile.content-extractor"
check_file "deployments/docker/Dockerfile.vision-service"
check_file "deployments/docker/Dockerfile.summarization-service"
check_file "deployments/docker/Dockerfile.embedding-service"
check_file "deployments/docker/Dockerfile.vector-store"
check_file "deployments/docker/Dockerfile.query-service"
echo ""

echo -e "${YELLOW}📚 Checking Documentation${NC}"
check_file "docs/ARCHITECTURE.md"
check_file "docs/API_REFERENCE.md"
check_file "docs/DEPLOYMENT.md"
check_file "docs/DEVELOPMENT.md"
check_file "credentials/README.md"
echo ""

echo -e "${YELLOW}🔄 Checking CI/CD Workflows${NC}"
check_file ".github/workflows/ci-cd.yml"
check_file ".github/workflows/codeql.yml"
check_file ".github/workflows/docker.yml"
check_file ".github/workflows/dependency-updates.yml"
check_file ".github/workflows/release.yml"
check_file ".github/dependabot.yml"
echo ""

echo -e "${YELLOW}🛠️  Checking Scripts${NC}"
check_file "scripts/setup.sh"
echo ""

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}  Verification Summary${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""
echo -e "Passed: ${GREEN}$PASS${NC}"
echo -e "Failed: ${RED}$FAIL${NC}"
echo ""

if [ $FAIL -eq 0 ]; then
    echo -e "${GREEN}🎉 All checks passed! Project structure is complete.${NC}"
    echo ""
    echo "Next steps:"
    echo "1. cp .env.example .env"
    echo "2. Edit .env with your API keys"
    echo "3. Run: ./scripts/setup.sh"
    echo "4. Run: make build"
    echo "5. Run: make test"
    exit 0
else
    echo -e "${RED}⚠️  Some files are missing. Please check the output above.${NC}"
    exit 1
fi
