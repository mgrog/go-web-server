#!/bin/bash
set -e

# Database configuration
DB_NAME="db"
DB_USER="user"
DB_PASSWORD="secret"
DB_HOST="localhost" # Change if not using port mapping
DB_PORT="5432"

echo "Resetting database: $DB_NAME"

# Terminate all connections to the database
PGPASSWORD="$DB_PASSWORD" psql -h $DB_HOST -p $DB_PORT -d postgres -U $DB_USER -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '$DB_NAME' AND pid <> pg_backend_pid();"
echo "Terminated active connections."

# Drop the database if it exists
PGPASSWORD="$DB_PASSWORD" psql -h $DB_HOST -p $DB_PORT -d postgres -U $DB_USER -c "DROP DATABASE IF EXISTS $DB_NAME;"
echo "Database dropped."

# Create the database
PGPASSWORD="$DB_PASSWORD" psql -h $DB_HOST -p $DB_PORT -d postgres -U $DB_USER -c "CREATE DATABASE $DB_NAME;"
echo "Database created."

# Optional: Initialize with schema
# PGPASSWORD="$DB_PASSWORD" psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f /path/to/schema.sql
# echo "Schema initialized."

echo "Database reset complete."