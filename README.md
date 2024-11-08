# 📝 Daily-Task ToDo App

A simple and efficient **Daily ToDo App** built using the **Go Fiber** framework for the backend RESTful API and **React+Vite+TypeScript** for the frontend. This app helps you manage your daily tasks effortlessly with **CRUD (Create, Read, Update, Delete)** functionalities.

## 🚀 Features

- **Add New Tasks:** Quickly add new tasks for the day.
- **Edit Tasks:** Update task details in a flash.
- **Delete Tasks:** Remove tasks once they are completed, using a dedicated delete button.
- **Real-Time Updates:** Enjoy an intuitive interface for easy task management.

## 🛠️ Tech Stack

### Backend
- [Go Fiber](https://gofiber.io/) - High-performance web framework in Go
- **REST API** with CRUD methods:
  - `GET` - Retrieve tasks
  - `POST` - Add new tasks
  - `PATCH` - Update task details
  - `DELETE` - Remove tasks

### Frontend
- [React](https://reactjs.org/) - Frontend library for building user interfaces
- [Vite](https://vitejs.dev/) - Next-generation frontend tooling
- [TypeScript](https://www.typescriptlang.org/) - Typed JavaScript for safer code

## 📸 Screenshots
![image](https://github.com/user-attachments/assets/60cb5cbe-1d74-4bb1-8bd0-299443151364)



## 📦 Installation & Setup

1. **Clone the repository:**
    ```bash
    git clone https://github.com/Endale2/GoLang-ReactJs-Todo-App.git
    cd daily-task
    ```

2. **Backend Setup (Go Fiber):**
    - Install Go dependencies:
      ```bash
      go mod tidy
      ```
    - Start the Go Fiber server:
      ```bash
      go run main.go
      ```
    - API endpoints:
      - `GET /api/todos` - List all tasks
      - `POST /api/todos` - Create a new task
      - `PATCH /api/todos/:id` - Update an existing task
      - `DELETE /api/todos/:id` - Delete a task

3. **Frontend Setup (React+Vite+TypeScript):**
    - Navigate to the `client` folder:
      ```bash
      cd client
      ```
    - Install frontend dependencies:
      ```bash
      npm install
      ```
    - Run the frontend:
      ```bash
      npm run dev
      ```

4. **Visit the App:**
    - Open [http://localhost:3000](http://localhost:5173) in your browser to access the app.

## 🧪 Testing

Use tools like **Postman** to test the backend endpoints and verify the CRUD operations are functioning as expected.


