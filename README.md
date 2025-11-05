# Notification Service

This is a Go-based REST API that simulates the processing of notifications from external sources, then processes them and sends them via email to end clients. This application has rate limit settings to prevent system overload and/or avoid excessive emails sent to the end client.

![api schema](https://i.imgur.com/GEIskgE.png)

## How to Run

1.  **Start the server:**

    ```bash
    go run cmd/api/main.go
    ```

2.  **Send a notification:**

    Use a tool like `curl` (or any other Postman-like tools) to send a POST request to the `/send` endpoint:

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{
      "type": "status",
      "user_id": "user123",
      "message": "Your order has been shipped!"
    }' http://localhost:8080/send
    ```

## Rate Limit Rules - Examples

-   **Status:** 2 per minute
-   **News:** 1 per day
-   **Marketing:** 3 per hour

## Quick Demo
Short video demonstrating the functionality with the example of a marketing notification
[![Watch the video](https://img.youtube.com/vi/4ooXk44XNH0/hqdefault.jpg)](https://youtu.be/4ooXk44XNH0)

