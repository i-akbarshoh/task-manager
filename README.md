# Task Manager
This task manager is simple crm project for development team4

## Run project
1. Clone project
2. Install docker and docker-compose
3. Run command `docker-compose up -d --build`

## To test connection rest api
```bash
curl -X GET http://localhost:8080/ping
```

if response is `pong` then connection is ok

