✅ Recommended Project Structure
go
Copy
Edit
📁 Build a Fully Server-Rendered Web App in Go/
├── Webapp/
│   ├── config/
│   │   └── .env
│   ├── static/
│   │   └── style.css
│   ├── template/
│   │   └── devTo.html
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── hello/
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── README.md
📄 README.md
markdown
Copy
Edit
# Go Server-Rendered Web Application 🚀

A fully server-rendered web app built using Go’s standard `net/http` and `html/template` packages. No frontend frameworks. Just performance, simplicity, and native Golang.

---

## 🔧 Features

- ✅ Server-side rendering using Go’s `html/template`
- ✅ Environment-based config with `.env`
- ✅ Modular project structure
- ✅ Static asset support (CSS, images, etc.)
- ✅ Clean architecture and routing
- ✅ Built with Go modules

---

## 📂 Project Modules

### Webapp (Main App)
A dynamic HTML server using native Go tools.

- `main.go` — entry point and routing
- `template/devTo.html` — HTML template
- `static/style.css` — basic styling
- `config/.env` — environment variables (can hold port, secrets)

### hello (Demo Module)
A small standalone Go module to test basic Go module support and structure.

---

## 🚀 Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/Mysteriousboy727/Built-a-fully-rendered-server-web-application.git
cd Build\ a\ Fully\ Server-Rendered\ Web\ App\ in\ Go/Webapp
2. Install Go
Make sure you have Go 1.19+ installed.
Download: https://golang.org/dl/

3. Run the Server
bash
Copy
Edit
go run main.go
Server runs on http://localhost:8080

🧠 Concepts Demonstrated
Low-level routing with http.ServeMux

Template parsing with html/template

Static asset handling (/static)

Code separation and modularity

.env config loading using os package

📈 Potential Improvements
Session management with cookies

Integration with PostgreSQL or Redis

Dockerfile for containerized deployment

Unit + integration tests with Go’s testing package

Support for form submissions

💼 Relevance to Industry Roles
This project reflects:

Strong fundamentals in Go’s HTTP stack

Backend experience with no external frameworks

Server-side rendering and templating logic

Practical understanding of file structure, separation of concerns, and performance

Ideal for internships or roles in:

Backend Engineering

Cloud Development

Systems Programming

DevOps & Infrastructure









