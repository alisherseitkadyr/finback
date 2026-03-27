#!/bin/bash

# Root backend
mkdir -p backend/cmd/server
mkdir -p backend/internal/platform/{db,logger,config}

# Auth module
mkdir -p backend/internal/auth/{handler,usecase,repository,model}

# Content module
mkdir -p backend/internal/content/{handler,usecase,repository,model}

# Future modules
mkdir -p backend/internal/progress
mkdir -p backend/internal/payment

# Shared
mkdir -p backend/migrations

# Root files
touch backend/cmd/server/main.go
touch backend/Dockerfile
touch backend/go.mod
touch backend/.env

# Platform files
touch backend/internal/platform/db/postgres.go
touch backend/internal/platform/logger/logger.go
touch backend/internal/platform/config/config.go

# Auth files
touch backend/internal/auth/handler/http.go
touch backend/internal/auth/usecase/service.go
touch backend/internal/auth/repository/postgres.go
touch backend/internal/auth/model/user.go

# Content files
touch backend/internal/content/handler/http.go
touch backend/internal/content/usecase/service.go
touch backend/internal/content/repository/postgres.go
touch backend/internal/content/model/topic.go
touch backend/internal/content/model/lesson.go

echo "✅ Modular Monolith Backend Structure Created"