# Crypto Price Tracker

## Overview

The Crypto Price Tracker is a service that calls the Coin Desk API to fetch Bitcoin price data, parses it, and provides the response to the user. The data is cached in internal memory and returned from memory if it is within the configured expiry time to minimize API calls. The expiry time is configurable via environment variables. The service logs all requests and responses using middleware and is designed to be modular to support other API providers in the future.

## Requirements

- Fetch Bitcoin price data from the Coin Desk API.
- Cache the data in internal memory with a configurable expiry time.
- Return data from memory if it is within the expiry time.
- Log requests and responses using middleware.
- Ensure the design is modular to support other API providers.
- Use Go language and any packages you like.
- Use the Gin framework: [Gin Framework](https://github.com/gin-gonic/gin).

## Setup and Configuration

1. **Clone the repository:**

    ```sh
    git clone https://github.com/ravi11kumar/crypto-price-tracker.git
    cd crypto-price-tracker
    ```
  
2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Run the application:**

    ```sh
    go run main.go
    ```

## API Endpoint

### Get Bitcoin Price

- **URL:** `/price`
- **Method:** `GET`
- **Response:**

    ```json
    {
        "data": {
            "bitcoin": {
                "EUR": "61205.9200",
                "GBP": "51574.4240",
                "USD": "66623.9820"
            }
        }
    }
    ```

## Implementation Details

### Main Components

- **Controllers:** Contains the logic for handling HTTP requests.
- **Handlers:** Initializes the services and controllers.
- **Interfaces:** Defines the service interfaces.
- **Middleware:** Contains the logging middleware.
- **Models:** Defines the data structures.
- **Routers:** Sets up the routes for the application.
- **Services:** Contains the business logic and communicates with external APIs.
- **Utils:** Contains utility functions, such as configuration loading.

### Configuration Loading

Configuration is loaded from a `config.json` file located in the `configs` directory. This file contains the server port and cache expiry duration.

### Logging Middleware

All requests and responses are logged using a custom middleware located in `middleware/logger.go`.

### Price Fetching and Caching

The `PriceService` fetches Bitcoin price data from the Coin Desk API and caches it in memory. If the cached data is within the configured expiry time, it is returned directly from memory.

### Error Handling

Errors are handled gracefully, and appropriate error messages are returned to the client.

## Data Source

Coin Desk API: [`https://api.coindesk.com/v1/bpi/currentprice.json`](https://api.coindesk.com/v1/bpi/currentprice.json)
