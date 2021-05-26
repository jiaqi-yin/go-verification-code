# Verification Code Service

This is a microservice to generate and verify a 6 digit verification code. It can be applied in the common verification use case.

Basic use cases

  - The verification code is generated as per mobile number and will expire in 2 minutes.
  - The verification code generation is limited to 3 times a day per mobile number. After exceeding the limit, users have to retry in the next day.

## Getting Started

### Build and Run

Run `docker-compose up --build`

### Sample Requests

- Endpoint: `/ping`

  Request: `curl -X GET 'localhost:8080/ping'`

  Response: `pong`

- Endpoint: `/generate`

  Request: `curl -X POST 'localhost:8080/generate' -H 'Content-Type: application/json' -d '{"phone_number": "0987654321"}'`

  Response: `true`

- Endpoint: `/verify`

  Request: `curl -X POST 'localhost:8080/verify' -H 'Content-Type: application/json' -d '{"phone_number": "0987654321","verification_code": "572631"}'`

  Response: `true`

### Get Verification Code

> The verification code generated via `/generate` is stored in Redis. The SMS service is not included in this app. It would be a future feature to add in. For now, you can get it via the commands below.

1. Run `docker run -it --rm --network go-verification-code_app-tier bitnami/redis:latest redis-cli -h redis` to connect to the `redis` container.
2. Get the verfication code by key, e.g: `get code:0987654321`.
3. Then you can use the verfication code in the `/verify` request.