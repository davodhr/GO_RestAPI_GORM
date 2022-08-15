** "Golang RestAPI With GORM and PostgreSQL Database" **

** This Microservice get input rectangles and check if any of them overlap with a main rectangle and save it on data base **

1. Run:
```
    docker-compose up --build
```

2. Get All Saved Rectangles and Time as json:
    GET Request Address: http://localhost:3200

3. Post New json Data containing "The main Rectangle" and "Input rectangles":
    POST Request Address: http://localhost:3200


Sample POST REQUEST:
```
curl --request POST --url http://localhost:3200 --header 'Content-Type: application/json' --data '{"main":{"x":3,"y":2,"width":5,"height":10},"inputs":[{"x":4,"y":11,"width":1,"height":1},{"x":9,"y":10,"width":6,"height":9}]}'
```


Sample GET REQUEST:
```
curl --request GET --url http://localhost:3200 
```


