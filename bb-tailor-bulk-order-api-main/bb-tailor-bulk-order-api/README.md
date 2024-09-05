# Brooks Bingham Tailor And Bulk Orders Backend REST API

## Overview

This repository provides the initial setup for the backend of a REST API built with Go Lang. The framework leverages:

- Echo for routing and handling HTTP requests.
- GORM for interacting with a MySQL database.
- JWT for authentication.
- Casbin for Role-Based Access Control (RBAC).

## Features

### Routing & Handlers:
- Echo framework facilitates routing and handling HTTP requests.
- Handlers are implemented for authentication, user management, and health checks.

### Middleware:
- JWT authentication middleware ensures secure access.
- Access Control middleware implements RBAC with Casbin.

### Database Interactions:
- GORM configuration enables communication with the MySQL database.
- Models for all database tables are defined.

### Command Scripts:
- `cmd/server/main.go`: Script to run the server.
- `cmd/db/sync_schema/main.go`: Script to synchronize the database schema.
- `cmd/db/drop_schema/main.go`: Script to drop the database schema.
- `cmd/db/seed/main.go`: Script to seed the database tables.

### Data Transfer Objects (DTOs):
- Defined DTOs manage request payloads and responses effectively.

### Data Seeding:
- JSON files store data for countries, states, cities, timezones, and currencies.
- Data mappers handle reading and processing of this JSON data.

### Documentation:
- Swagger UI provides API documentation through a comprehensive JSON document.

### Configuration:
- Configuration files are provided for server settings and the Casbin RBAC model.

### Utilities:
- Utility functions are available for hashing, validation, transformations, and JWT operations.

### Server Setup:
- The server initialization includes graceful shutdown and route registration for all resources.

## Folder Structure

```
.
├── api/
│   ├── dtos/          # Data Transfer Objects
│   ├── handlers/      # Handlers for API resources
│   └── middlewares/   # Middlewares for authentication and access control
├── cmd/               # Command scripts
│   ├── db/            # Database scripts (sync, drop, seed)
│   └── server/        # Server script
├── data/              # JSON data files and mappers
├── internal/
│   ├── config/        # Configuration files
│   ├── database/      # Database models and seeders
│   ├── server/        # Server logic and initialization
│   └── utils/         # Utility functions (moved here)
└── web/
    └── static/        # Static files and assets including Swagger UI
```

## How to Test

**Prerequisites:** Ensure you have Go Lang and MySQL installed.

1. **Sync the Database Schema:**
    ```bash
    make sync-schema
    ```

2. **Seed the Database Tables (Optional):**
    This step populates the database with initial data.
    ```bash
    make seed
    ```

3. **Start the Server:**
    ```bash
    make server
    ```

## Extras

**Drop the Database Schema:**
```bash
make drop-schema
```

**Note:** Dropping the schema will delete all data from your database. Use this command with caution.

## Contributing

Pull requests and contributions are welcome! Please refer to the contribution guidelines for details.
