#  Internal Transfers App

A **Golang Application** to handle internal financial transactions between accounts, with data persistence using **PostgreSQL**.  
The service exposes HTTP APIs for:
- Creating accounts
- Querying account balances
- Submitting transactions (transfers between accounts)

The entire setup is containerized using **Docker + Docker Compose** for easy deployment and local development.

---

##  Features

- Create an account with a specific ID and initial balance  
- Query account balance by account ID  
- Submit transactions between accounts (source ➡ destination)  
- Data persistence with PostgreSQL  
- Clean, modular Go code following clean code principles  
- Dockerized for easy setup  
- Example Postman API requests  

---

##  Tech Stack

- **Golang 1.20**
- **PostgreSQL 15**
- **Docker + Docker Compose**
- **gorilla/mux** (HTTP routing)
- **pq** (PostgreSQL driver)

---

##  Assumptions

- All accounts use the same currency  
- No authentication or authorization required  
- Floating point string format is used for amounts (e.g. `"100.25"`)  
- This is a **local development / internal tool setup** (not production-hardened)

---

##  Getting Started

###  Clone the repo
```bash
git clone https://github.com/anurag8773/Internal-Transfers-App.git
cd Internal-Transfers-App
````

---

###  Run with Docker

```bash
docker-compose up --build
```

 This will start both the **Golang app** and **PostgreSQL DB**.

---

###  Apply DB migrations

Once containers are up:

```bash
docker exec -i <your-db-container-name> psql -U postgres -d transfers < migrations/schema.sql
```

 Example:

```bash
docker exec -i internal-transfers-db-1 psql -U postgres -d transfers < migrations/schema.sql
```

(Replace `internal-transfers-db-1` with the actual container name from `docker ps`)

---

##  API Endpoints

###  **Create Account**

```
POST /accounts
```

Request body:

```json
{
  "account_id": 1,
  "initial_balance": "100.00"
}
```

---

###  **Get Account**

```
GET /accounts/{account_id}
```

Response:

```json
{
  "account_id": 1,
  "balance": 100.00
}
```

---

###  **Submit Transaction**

```
POST /transactions
```

Request body:

```json
{
  "source_account_id": 1,
  "destination_account_id": 2,
  "amount": "50.00"
}
```

---

##  Example Postman Usage

 Import this collection:
 [Postman Collection Link](https://documenter.getpostman.com/view/37271849/2sB34bL3VY)

---

##  Environment Variables

Provide your own `.env` file or use defaults in `docker-compose.yml`.

Example `.env.example`:

```
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=transfers
APP_PORT=8080
```

---

##  Development Notes

* Default DB credentials are safe for **local development only**
* Use `docker-compose logs` to debug container logs
* Use `docker-compose down -v` to tear down containers and volumes

---

##  License

This project is for learning/demo purposes. No license specified.

---

##  Author

 **Anurag Kumar Maurya**
