# HNG DevOps Interns - The Cool Kids : Stage 1
### Create an API that takes a number and returns interesting mathematical properties about it, along with a fun fact.
### Technology Stack: 
- **Go**
- Gorilla/mux router
- NumbersAPI (external data source)

### Endpoint: **GET** *your-url*/api/classify-number?number=371

 ## Steps for any valid number(n):
 - Check if it's prime.
 - Check if it's perfect.
 - Calculate digit sum.
 - Check if it's an Armstrong number.
 - Determine if it's odd or even.
 - Collect all properties that are true.
 - Get the fun_fact from Numbers API.

 ### Example JSON response:
 ```json
 {
    "number": 2,
    "is_prime": true,
    "is_perfect": false,
    "properties": [
        "armstrong",
        "even"
    ],
    "digit_sum": 2,
    "fun_fact": "2 is a primorial, as well as its own factorial."
}
```
### Error Response (400 Bad Request)
```json
{
    "number": "invalid-input",
    "error": true
}
```
### Structure - at a high level:
- Use a router like gorilla/mux 
- Create a handler function that takes the number from the URL path, validates it, then processes it.
- For each computation (prime, perfect, etc.), create helper functions.

## Installation

### Prerequisites
- Go 1.20+ installed
- Git for repository cloning

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/your_repo_name.git
   cd your_repo_name
2. Install dependencies:
    ```bash
    go get github.com/gorilla/mux
    ```
3. Build and run:
    ```bash
    go build -o number-api
    ./number-api
    ```
The API will be available at http://localhost:8080

## Example Requests
- Valid Request
    ```bash
    curl "http://localhost:8080/api/classify-number?number=371"
    ```
- Invalid Input
    ```bash
    curl "http://localhost:8080/api/classify-number?number=abc"
    ```
- Missing Parameter
    ```bash
    curl "http://localhost:8080/api/classify-number"
    ```
## Deployment

1. Build the binary:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o number-api
   ```
2. Deploy to your server

3. Run as a background service using systemd or supervisor

4. Configure reverse proxy (Nginx/Apache) if needed

### Testing
- Run the application and test with:
  ```bash
  curl "http://localhost:8080/api/classify-number?number=28"
  ```
- Expected response:
  ```json
  {
    "number": 28,
    "is_prime": false,
    "is_perfect": true,
    "properties": ["even"],
    "digit_sum": 10,
    "fun_fact": "28 is a perfect number."
   }
   ```
## Dependencies:
- [NumbersAPI](numbersapi.com) - Provides mathematical facts
- [Gorilla/mux](https://github.com/gorilla/mux) - Request router and dispatcher