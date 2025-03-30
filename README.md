## getting started as a newb with docker (for me)

### Prerequisites
1. **Install Docker**:
   If Docker is not installed, follow the [official Docker installation guide](https://docs.docker.com/get-docker/).

2. **Install Docker Compose**:
   If `docker-compose` is not installed, you can install it with:
   ```bash
   sudo apt install docker-compose
   ```

---

### Setting Up the Project

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd loadboard
   ```

2. **Set Up the `.env` File**:
   Create a `.env` file in the root directory with the following content:
   ```properties
   DB_HOST=db
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database
   DB_PORT=5432
   ```

   Alternatively, copy the example file:
   ```bash
   cp .env.example .env
   ```

---

### Rebuilding and Starting Your Containers

1. **Stop Existing Containers**:
   ```bash
   docker-compose down
   ```

2. **Rebuild the Containers**:
   ```bash
   docker-compose build
   ```

3. **Start the Containers**:
   ```bash
   docker-compose up
   ```

---

### Testing the Application

1. **Access the Application**:
   Open your browser or use `curl` to test the application:
   ```bash
   curl http://localhost:8080
   ```

2. **Register a User**:
   Use the `/auth/register` endpoint to create a new user:
   ```bash
   curl -X POST http://localhost:8080/auth/register \
   -H "Content-Type: application/json" \
   -d '{"email": "test@example.com", "password": "password123", "role": "carrier"}'
   ```

3. **Log In to Get a Token**:
   Use the `/auth/login` endpoint to log in and get a JWT token:
   ```bash
   curl -X POST http://localhost:8080/auth/login \
   -H "Content-Type: application/json" \
   -d '{"email": "test@example.com", "password": "password123"}'
   ```

4. **Access Protected Routes**:
   Use the token from the login response to access protected routes:
   ```bash
   curl -X GET http://localhost:8080/loads/ \
   -H "Authorization: Bearer your-jwt-token"
   ```

---

### Notes
- Make sure Docker is running before executing the commands.
- If you encounter any issues, check the logs:
  ```bash
  docker-compose logs
  ```
- To stop the containers, press `Ctrl+C` or run:
  ```bash
  docker-compose down
  ```

---

### Summary
This `README.md` now includes:
1. Prerequisites for setting up Docker and Docker Compose.
2. Instructions for setting up the `.env` file.
3. Steps to rebuild and start the containers.
4. Basic testing instructions for the application.
