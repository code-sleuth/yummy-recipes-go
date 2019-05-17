# Yummy Recipes API Golang

This is a minimalistic REST api written in [Go](https://golang.org/).

### Deps
- [Postgresql](https://www.postgresql.org/download/) - Datastore
- [GORM](http://gorm.io/docs/) - An opensource ORM for golang
- [Gorilla Mux](https://github.com/gorilla/mux) - a URL router and dispatcher for golang
- [Godotenv](https://github.com/joho/godotenv) - loads environment variables

### Endpoints

| Endpoint | Function | Request payload |
| ------ | ------ | ------ |
| POST /users | Creates a user |```{"email": "e@gmail.com", "username": "megatron", "fullname": "Mega Tron", "password": "eX@mpL3"}```
| GET /users | Get all users |
| GET /users/{id} | Get a single user |
| PUT /users/{id} | Update user | ```{"email": "e@gmail.com", "username": "gigatron", "fullname": "Giga Tron"}```
| DELETE /users/{id} | Delete user | 

## Project SetUP