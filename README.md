# POS App Backend

A Point of Sale (POS) backend system built with Go. This application provides a RESTful API for managing products, processing transactions, and generating reports, complete with user authentication and Role-Based Access Control (RBAC).

## 🚀 Tech Stack

- **Language:** [Go (Golang)](https://golang.org/) v1.26.3
- **Framework:** [Gin Gonic](https://gin-gonic.com/)
- **Database:** [MongoDB](https://www.mongodb.com/)
- **Authentication:** JWT (JSON Web Token)
- **Middleware:** CORS, Auth Middleware, RBAC

## 📂 Folder Structure

```text
.
├── config/             # Database connection and configuration
├── handlers/           # Request handlers (Controllers)
│   ├── auth.go         # Login and Register logic
│   ├── product.go      # Product CRUD logic
│   └── transaction.go  # Transaction and Reporting logic
├── middleware/         # Custom Gin middleware (Auth, AdminOnly)
├── models/             # MongoDB data models (Schemas)
├── routes/             # API route definitions
├── main.go             # Application entry point
└── .env                # Environment variables (ignored by git)
```

## 🛠️ Getting Started

### Prerequisites

- Go installed on your machine
- MongoDB instance (local or Atlas)

### Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd pos-app
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Configure Environment Variables**
   Create a `.env` file in the root directory and add the following:
   ```env
   MONGODB_URI=mongodb://localhost:27017
   JWT_SECRET=your_secret_key_here
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```
   The server will start at `http://localhost:8080`.

## 🛣️ API Endpoints

### Authentication
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login and receive JWT

### Products (Protected)
- `GET /api/products` - List all products
- `POST /api/products` - Create a product (Admin only)
- `PUT /api/products/:id` - Update a product (Admin only)
- `DELETE /api/products/:id` - Delete a product (Admin only)

### Transactions (Protected)
- `GET /api/transactions` - List all transactions
- `POST /api/transactions` - Create a new transaction

### Reports (Protected)
- `GET /api/reports/daily` - Get daily sales report (Admin only)
