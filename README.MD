# Energy Project

This project sets up a MySQL database and loads data from a CSV file into a table using Docker.

## Prerequisites

The following tools are required to run this project:

- Docker (https://www.docker.com/products/docker-desktop)
- Docker Compose (https://docs.docker.com/compose/install/)
- Make (https://www.gnu.org/software/make/)

Ensure you have these installed and available in your command line path before proceeding.

## Setting Up

Follow these steps to setup the database and load data:

1. **Build the Docker image**
    ```bash
    make build-db
    ```
   This command builds a Docker image using the Dockerfile in the current directory, which is configured to set up a MySQL 8.0 database.

2. **Start the Docker container**
    ```bash
    make run-db
    ```
   This command starts a Docker container based on the image created in the previous step, running the MySQL service inside the container.

3. **Perform the migration**
    ```bash
    make migrate
    ```
   This command creates the table in the MySQL database using the migration.sql file.

4. **Load Data**
    ```bash
    make load-data
    ```
   This command loads data into the database table from the test_bia.csv file.

## Accessing the Database

Once the Docker container is running, you can access the MySQL database using the MySQL client or any MySQL-compatible tool. The database runs on port 3306 in the Docker container, and this port is mapped to port 3306 on your host machine for your convenience.

The root password for the MySQL database is set to `root`. In a production environment, be sure to use a more secure password and manage it using appropriate secret management tools.

---
