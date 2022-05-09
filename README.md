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

### To Do
 - clean architecture  
    - https://gist.github.com/mpppk/609d592f25cab9312654b39f1b357c60  
 - DDD/domain-driven design  
    - https://codezine.jp/article/detail/11968  
 - GitHub Actions  
    - https://knowledge.sakura.ad.jp/23478/  
 - Terraform
    - https://blog.dcs.co.jp/aws/20210401-terraformaws.html  
    - How to use  
    1. create 3 files(main.tf resource.tf variables.tf)  
        main.tf      : define the provider like AWS, GCP, azure  
        resoure.tf   : declare services  
        variables.tf : set the variables  
    2. check files
        ```
        terraform validate
        ```
    3. execute terraform
        ```
        terraform plan
        terraform apply
        ```
    4. delete terraform
        ```
        terraform destroy
        ```
 - Mongo DB
    - https://openstandia.jp/oss_info/mongodb/  

 - GraphQL
    - https://www.redhat.com/ja/topics/api/what-is-graphql  

