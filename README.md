# Receipt Processor API
## Running the Application

### Without Docker
1. Run the application: go run main.go
4. The application will start on `http://localhost:8080`.

### With Docker
1. Build the Docker image: docker build -t receipt-processor .
2. Run the Docker container: docker run -p 8080:8080 receipt-processor
3. The application will start on `http://localhost:8080`.

## API Endpoints

### POST /receipts/process
- **Description:** 
        Path: /receipts/process
        Method: POST
        Payload: Receipt JSON
        Response: JSON containing an id for the receipt.
- **Request (testing done through CMD and Postman)**:
  Example: curl -X POST http://localhost:8080/receipts/process ^
        -H "Content-Type: application/json" ^
        -d @examples/*****.json
- **Response:**
  { "id": "res-pon-se-id123123123" }

### **GET /receipts/{id}/points**
- **Description:**
        Path: /receipts/{id}/points
        Method: GET
        Response: A JSON object containing the number of points awarded.
- **Request (testing done through CMD and Postman)**:
    Example: curl -X GET http://localhost:8080/receipts/res-pon-se-id123123123/points
- **Response:**
  { "points": 32 }
