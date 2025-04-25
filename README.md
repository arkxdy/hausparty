# 🎉 HausParty - Host & Join House Parties

A microservices-based web and mobile platform to host and join small-scale house parties. Built for web (React.js) and mobile (React Native + WebView), with scalable microservices using Go. API Gateway to handle request from UI using Node.js and Express.js

## 🚀 Features

- ✅ Host or join parties
- ✅ Admin approval workflow for new hosts
- ✅ Ratings & feedback for hosts and users
- ✅ Realtime booking using Kafka (first-come-first-serve)
- ✅ Shared UI for web & mobile via WebView
- ✅ Fully containerized with Docker (dev & prod setups)

## 🧱 Tech Stack

- Frontend: React.js
- Backend: Go (Gin), Node.js, Express.js, Docker, Kafka
- Mobile: React Native with WebView
- DB: PostgreSQL + MongoDB + Redis
- Infra: Docker Compose

## 🏗️ Microservices

- `auth-service`: Login/signup, JWT
- `user-service`: User profiles
- `party-service`: Party creation/joining
- `rating-service`: Ratings/feedback
- `admin-service`: Approvals, admin controls

## 📦 Run Locally

```bash
docker-compose -f docker-compose.dev.yml up --build
