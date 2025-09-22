# NWUCA Management System

This is the official management system for the Northwest University Computer and Network Security Association (NWUCA).

## Project Structure

This project follows a monorepo architecture:

-   `/client`: Frontend application built with Vue 3 and uni-app.
-   `/server`: Backend application built with Go and Gin.
-   `/docs`: Project documentation.

## Quick Start

### Prerequisites

-   Go (1.18 or higher)
-   Node.js (16.x or higher)
-   pnpm (or npm/yarn)
-   PostgreSQL (or other database supported by GORM)

### Backend (`server`)

1.  **Navigate to the server directory:**
    ```bash
    cd server
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Configure environment:**
    Create a `.env` file based on `.env.example` and fill in your database connection details.

4.  **Run the server:**
    ```bash
    go run cmd/api/main.go
    ```
    The server will start on `http://localhost:8080`.

### Frontend (`client`)

1.  **Navigate to the client directory:**
    ```bash
    cd client
    ```

2.  **Install dependencies:**
    ```bash
    pnpm install
    ```

3.  **Run for web (H5):**
    ```bash
    pnpm dev:h5
    ```

4.  **Run for other platforms (e.g., WeChat Mini Program):**
    ```bash
    pnpm dev:mp-weixin
    ```

## Learn More

-   [Gin Web Framework](https://gin-gonic.com/)
-   [GORM](https://gorm.io/)
-   [uni-app](https://uniapp.dcloud.io/)
-   [Ant Design Vue](https://www.antdv.com/)
