# golang/next(react, typescript)/MongoDB/terraform
SPA(golang × react × typescript)

## How to run api  
1. git clone
2. input below
   ```
   $ make run //docker-compose up -d
   ```
3. access to localhost:8080 with chrome or somethig else
   ```
   $ curl -X GET 'localhost:8080'    //Hello World!
   ```

## uri   
   ### HTTP response
   ```
   $ curl -X GET 'localhost:8080/param/{id}'    //GET param/{id}
   $ curl -X POST 'localhost:8080/param/{id}'   //POST param/{id}
   $ curl -X PUT 'localhost:8080/param/{id}'    //PUT param/{id}
   $ curl -X DELETE 'localhost:8080/param/{id}' //DELETE param/{id}
   ```

# Tech
## Go
   - GraphQL
      - https://www.redhat.com/ja/topics/api/what-is-graphql  
      https://github.com/fatih/gomodifytags


## Next

   - how to run
      ```
      $ npm run dev
      ```

## MongoDB
 - https://openstandia.jp/oss_info/mongodb/  


## github actions
- GitHub Actions  
   - https://knowledge.sakura.ad.jp/23478/  

## Terraform
   - https://blog.dcs.co.jp/aws/20210401-terraformaws.html  
   - How to use  
      1. create 3 files(main.tf resource.tf variables.tf)  
         main.tf      : define the provider like AWS, GCP, azure  
         resoure.tf   : declare services  
         variables.tf : set the variables  
      2. check files
         ```
         terraform validate
         terraform fmt
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

## Architecture
- Clean architecture  
   - https://gist.github.com/mpppk/609d592f25cab9312654b39f1b357c60  
- DDD/domain-driven design  
   - https://codezine.jp/article/detail/11968  
   
   
