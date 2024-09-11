# Japanese addresses

handle Japanese addresses and property information from a CSV file

## Run

clone the repo, and run

```sh
go run main.go
```

## API

1. `/address/retrieval/{filename}`, [get], CSV Upload
    - Accept a CSV file containing Japanese addresses and property information
2. `/address/upload`, [post], Property Information Retrieval
    - Return all property information in the specified format

When the app is running, browse to http://localhost:8080/swagger/index.html, see [Swagger 2.0 Api documents]((https://github.com/swaggo/swag))