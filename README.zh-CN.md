# NWUCA 管理系统

本项目是为西北大学计算机与网络空间安全协会 (NWUCA) 开发的官方管理系统。

## 项目架构

本项目遵循 Monorepo 架构：

-   `/client`: 前端应用，基于 Vue 3 和 uni-app 构建。
-   `/server`: 后端应用，基于 Go 和 Gin 构建。
-   `/docs`: 项目文档。

## 快速开始

### 环境准备

-   Go (1.18 或更高版本)
-   Node.js (16.x 或更高版本)
-   pnpm (或 npm/yarn)
-   PostgreSQL (或其他 GORM 支持的数据库)

### 后端 (`server`)

1.  **进入 server 目录:**
    ```bash
    cd server
    ```

2.  **安装依赖:**
    ```bash
    go mod tidy
    ```

3.  **配置环境:**
    基于 `.env.example` 文件创建一个 `.env` 文件，并填入您的数据库连接信息。

4.  **运行服务:**
    ```bash
    go run cmd/api/main.go
    ```
    服务将在 `http://localhost:8080` 启动。

### 前端 (`client`)

1.  **进入 client 目录:**
    ```bash
    cd client
    ```

2.  **安装依赖:**
    ```bash
    pnpm install
    ```

3.  **以 Web (H5) 模式运行:**
    ```bash
    pnpm dev:h5
    ```

4.  **运行到其他平台 (例如：微信小程序):**
    ```bash
    pnpm dev:mp-weixin
    ```

## 了解更多

-   [Gin Web 框架](https://gin-gonic.com/)
-   [GORM](https://gorm.io/)
-   [uni-app](https://uniapp.dcloud.io/)
-   [Ant Design Vue](https://www.antdv.com/)
