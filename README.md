# golang_next.js
SPA(Go × Next.js)

### How to run api !  
1. git clone <this repository>  
2. input command below
```
make up //docker-compose up -d
```
3. access to localhost:8080 with chrome or somethig else
```
curl -X GET 'localhost:8080'    //Hello World!
```

### How to check the  HTTP response
After run container, command below in Terminal.
```
curl -X GET 'localhost:8080/template'    //GET called!!
curl -X POST 'localhost:8080/template'   //POST called!!
curl -X PUT 'localhost:8080/template'    //PUT called!!
curl -X DELETE 'localhost:8080/template' //DELETE called!!
```

### To do
 - clean architecture  
 - DDD/domain-driven design  
 - git hub action  
 - Terraform  
 - Mongo DB
 - GraphQL
