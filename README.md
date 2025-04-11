# 🎉 HausParty - Host & Join House Parties

A microservices-based web and mobile platform to host and join small-scale house parties. Built for web (Next.js) and mobile (React Native + WebView), with scalable microservices using Go.

## 🚀 Features

- Host or join parties
- Admin approval for new hosts
- User rating system
- Mobile and web support (via WebView)
- Fully dockerized with dev/prod environments

## 🧱 Tech Stack

- Frontend: Next.js
- Backend: Go (Gin), Docker, Kafka (optional)
- Mobile: React Native with WebView
- DB: PostgreSQL + Redis
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
