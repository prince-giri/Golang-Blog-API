📚 Project Description
  This is a simple Blog Application API built with Golang and PostgreSQL.
  It demonstrates the fundamental concepts of authentication, authorization, CRUD operations, pagination, and JWT-based secure APIs.
  Users can register, login, create posts, update/delete their own posts, comment on any post, and manage their comments.

🚀 Features
  User Registration and Login with secure Password Hashing (bcrypt).
  JWT Authentication for protected routes.
  Users can Create, Read (with Pagination), Update, and Delete their own:
  Posts
  Comments
  Post Status Management: Posts can be published, drafted, or archived.
  Pagination support for posts listing (/posts?page=1&limit=10).
  Proper Authorization Checks (only owners can edit/delete their posts/comments).
  Clean project structure: Controllers, Models, Middleware, Database, and Utils.
  API tested via Postman.
  Database managed using PostgreSQL, viewed using pgAdmin.

🛠️ Tech Stack
  Language: Golang
  Framework: Gin Gonic
  Database: PostgreSQL
  ORM: GORM
  Authentication: JWT (JSON Web Tokens)
  Password Encryption: bcrypt

📁 Project Structure
  go-blog-app/
  ├── controllers/   // Contains all route handler logic (auth, posts, comments)
  ├── database/      // Database connection logic
  ├── middleware/    // JWT auth middleware
  ├── models/        // Database models
  ├── utils/         // JWT utility functions
  ├── go.mod         // Go module definition
  ├── go.sum         // Dependency lock file
  └── main.go        // Application entry point

