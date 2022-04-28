There are two ways to start the service you can follow any one of them or both 
## 1. Starting the service from code build
### starting url shortener service
``make up``
### stopping url shortener service
``make down``

## 2. Starting the service using docker compose file
``docker-compose up ``


API Calls

1. Generating a short URL
```azure
    Request Endpoint: http://localhost:8080/api/v1/generate_shortener_url
    Request Type: POST
    Request Body: {
                    "original_url": "https://google.com"
                  }
                  
    Response:
        {
            "message": "visit short url http://localhost:8080/api/v1/url?short_url=99999e"
        }
```
2. Making a request to short URL
```azure
    Request Endpoint: http://localhost:8080/api/v1/url?short_url=99999e
    Request Type: GET
```
