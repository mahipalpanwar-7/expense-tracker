# Expense Tracker

A production-ready expense tracker backend built with Golang. 

## Tech Stack

- Golang
- Gin Framework
- MySQL
- HTML, CSS, JS

## Project Structure

expese-tracker/
├── cmd\server/
├── internal/
│   ├── config/
│   ├── handler/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   └── service/
├── migrations/
├── pkg\utils/
├── static/
│   ├── css/
│   ├── img/
│   └── js/
├── templates/
│   ├── auth/
│   ├── expense/
│   └── layouts/
├── .gitignore
└── README.md

## Setup

1. Clone the repository
git clone https://github.com/mahipalpanwar-7/expense-tracker.git
cd expense-tracker

2. Create .env file from .env.example and update  MySQL password

3. Install dependencies
go mod tidy

4. Run the server
go run cmd/server/main.go

Server will run on http://localhost:8080

## Features (Planned)

- User authentication with JWT
- Add, edit, delete expenses
- Category management
- Monthly reports
- Budget tracking
- Dashboard with HTML, CSS, JS

## Branching Strategy

- main - production code
- develop - development code
- feature/* - new features

