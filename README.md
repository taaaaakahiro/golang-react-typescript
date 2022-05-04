# golang_next.js
SPA(Go × Next.js)

### How to run api !  
1. clone  
2. command
```
make up //docker-compose up -d
```
3. access to localhost:8080 with chrome or somethig else

### How to check the  HTTP response
After run container, command below in Terminal.
```
curl -X GET 'localhost:8080'    //GET called!!
curl -X POST 'localhost:8080'   //POST called!!
curl -X PUT 'localhost:8080'    //PUT called!!
curl -X DELETE 'localhost:8080' //DELETE called!!
```
