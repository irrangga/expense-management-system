# Expense Management System

An expense management system where employees can submit expenses, managers can approve them, and approved expenses get processed for payment.

## Prerequisites

Before starting, make sure you have the following installed on your system:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Quick Start

Follow these steps to get the application up and running:

### 1. Environment Configuration

Copy the sample environment file and configure it:

```bash
cp .env.sample .env
```

**Important:** Open the `.env` file and update any necessary configuration values according to your environment needs.

### 2. Start the Application

Run the following command to start all services:

```bash
docker-compose up
```

### 3. Access the Application

Once the containers are running, the application should be accessible at `http://localhost:3000`.

## Login Credentials

Use the following test accounts to access the application based on your role:

### Employee Accounts

| Email                   | Name              | Role     |
| ----------------------- | ----------------- | -------- |
| `employee1@example.com` | Arie Untung       | Employee |
| `employee2@example.com` | Bandung Bondowoso | Employee |
| `employee3@example.com` | Cecep Santoso     | Employee |

### Manager Account

| Email                 | Name        | Role    |
| --------------------- | ----------- | ------- |
| `manager@example.com` | Dono Warkop | Manager |
