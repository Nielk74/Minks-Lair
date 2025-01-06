# Cosmink Backend

Welcome to the backend service for **Cosmink**, a graphical insights analyzer. This project provides the server-side infrastructure, including user authentication, database interaction, and API handling, to support the visualization and analysis features of Cosmink.

## Project Structure

The backend is organized into a modular and scalable structure:

```
├── auth
│   ├── core          
│   └── infra         
├── graph
│   ├── controller   
│   ├── core          
│   └── infra        
└── libs
    ├── database
    ├── route        
    └── utils      

```

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.23 or higher)
- [Docker](https://www.docker.com/)
- [migrate](https://github.com/golang-migrate/migrate) (for database migrations)

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd cosmink-backend
   ```

2. Install Go dependencies:

   ```bash
   go mod download
   ```

3. Build the application:

   ```bash
   make build
   ```

4. Run the application:

   ```bash
   make run
   ```

## Docker

To build and run the application using Docker:

1. Build the Docker image:

   ```bash
   make docker-build
   ```

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 go-app:test
   ```

## Database Migration

1. Create a new migration:

   ```bash
   make create-migration name=<migration_name> db=<database_name>
   ```

2. Apply migrations using `migrate`:

   ```bash
   migrate -database postgres://localhost:5432/<db_name> -path migrations up
   ```

## Testing

Run unit tests:

```bash
make test
```

Run tests in Docker:

```bash
make docker-test
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m 'Add feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.


