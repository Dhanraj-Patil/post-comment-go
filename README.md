# 📝 Post-Comment Service (Go + Gin + MongoDB)

A simple backend service built using **Go**, **Gin**, and **MongoDB** that allows users to create **Markdown-supported rich text posts** and comment on them. Comments are stored as posts with a `thread` field referencing their parent post.

---

## 📦 Features

- Create new posts
- Comment on existing posts
- Retrieve posts with nested comments
- Swagger API documentation at `/swagger/index.html`

---

## 🛠️ Tech Stack

- Go (Gin framework)
- MongoDB
- Swaggo for Swagger UI

---

## 📄 Environment Setup

Create a `.env` file in the project root with the following variables:

```env
MONGODB_URI=""
DATABASE="post-comments"
COLLECTION="posts"
```

## Run the project

Install Dependencies and run the project.

```bash
go mod tidy
go run ./cmd .
```

To see and test API's visit http://localhost:8080/swagger/index.html .