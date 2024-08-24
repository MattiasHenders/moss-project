# Moss Communication Server

## Overview

The `moss-communication-server` is a simple server application that provides API routes for Stable Diffusion tasks and includes health check endpoints. It uses middleware for API key verification and handles text-to-image and image-to-image transformations.

## Table of Contents

- [Description](#description)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Server](#running-the-server)
- [Launching on Runpods.io](#launching-on-runpodso)

---

### <a name="description"></a> Description

This server provides authenticated routes for performing image transformation tasks using Stable Diffusion. It includes endpoints for generating images from text and modifying images based on existing ones. The server also has a health check route to ensure it's running correctly.

---

### <a name="installation"></a> Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-repo/moss-communication-server.git
   cd moss-communication-server
   ```

2. **Install Dependencies**

   Ensure you have Go installed and run:

   ```bash
   go mod tidy
   ```

3. **Set Up Environment Variables**

   Create a `.env` file in the root directory and fill it with your configuration details. An example `.env` file is provided below.

---

### <a name="configuration"></a> Configuration

Create a `.env` file in the root directory of the project with the following variables:

```ini
# Server config
PORT=8080
ENV=development

# DB config
databaseHost=localhost
databasePort=5432
databaseName=mydatabase
databaseUsername=myuser
databasePassword=mypassword

# Hashing config
hashSalt=your_hash_salt
passwordSecret=your_password_secret

# Demo config
demoAPIKey=your_demo_api_key
runpodAPIKey=your_runpod_api_key
```

Replace the placeholder values with your actual configuration settings.

---

### <a name="running-the-server"></a> Running the Server

1. **Start the Server**

   With the environment variables set, you can start the server with:

   ```bash
   go run main.go
   ```

   The server will start and listen on the port specified in the `.env` file.

2. **Verify API Key Middleware**

   Ensure that the API key used for authentication matches the `demoAPIKey` provided in the `.env` file.

---

### <a name="launching-on-runpodso"></a> Launching on Runpods.io

1. **Create a Runpods.io Account**

   Sign up or log in to your account at [Runpods.io](https://runpods.io).

2. **Create a New Pod**

   - Go to your dashboard and click on "Create New Pod".
   - Choose a suitable runtime environment (e.g., Node.js, Python) depending on your application's requirements.
   - Set up the environment to match the configuration of your server.

3. **Upload Your Code**

   - Upload the project files or connect a Git repository.
   - Ensure that the `.env` file is included or properly configured in the environment settings.

4. **Configure Environment Variables**

   - In the Runpods.io settings, configure the environment variables according to the `.env` file.

5. **Deploy and Run**

   - Start the pod and ensure it is running correctly.
   - Check the logs and health checks to verify that the server is functioning as expected.

6. **Access Your Server**

   - Obtain the URL provided by Runpods.io and use it to interact with your server's endpoints.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.