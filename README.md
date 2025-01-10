# Crypto Tracker - Backend API
This project tracks cryptocurrency data and provides APIs for fetching real-time stats and price deviation calculations. The backend is built using Golang, MongoDB is used for data storage, and CoinGecko API is used to fetch cryptocurrency data.


## Assignment Overview
This project consists of the following tasks:

### Background Job: Fetch data on Bitcoin, Matic, and Ethereum every 2 hours, including:
- Price in USD
- Market Cap in USD
- 24-hour price change percentage.

### API Endpoints:
/stats: Returns the latest data for a requested cryptocurrency.
/deviation: Returns the standard deviation of the price for the last 100 records of the requested cryptocurrency. 

### Technologies Used
Golang: For the backend
MongoDB: NoSQL database for storing cryptocurrency data
CoinGecko API: Public API to fetch cryptocurrency data

### Setup Instructions
- Install Golang 1.22.1
- Clone the repository
  ``` git 
  git clone github.com/rimo02/Crypto-Tracker
  ```
- Create a .env file 
  ``` bash
  API_KEY=YOUR_COINGECKO_API_KEY
  MONGO_URI=mongodb://localhost:27017
  ```
- ``` go
  go run main.go
  ```
Application will be running at localhost:3000

### API Endpoints

- Task 1: Background Job
A background job is implemented to fetch cryptocurrency data for Bitcoin, Matic, and Ethereum every 2 hours. This data is stored in the database.

- Task 2: /stats Endpoint
    * URL: localhost:3000/stats?coin=bitcoin
    * Method: GET
    * Sample Response:
    ``` json
    {
      "24HourChange": 2.49759,
      "Coin": "bitcoin",
      "MarketCap": 1882841171018,
      "Price": 95152,
      "Symbol": "btc"
    }
    ```
- Task 3: /deviation Endpoint
    * URL: localhost:3000/deviation?coin=bitcoin
    * Method: GET
    * Sample Response:
    ``` json
    {
        "coin": "bitcoin",
        "deviation": 374.1710838640528,
        "mean_price": 94966,
        "total_records": 14
    }
    ```