# Traditional Way:

## Backend Setup:
1. Make sure the PostgreSQL environment variables in `{root}/be/.env` are correctly configured.
* please dont change the env db if you use docker
   
2. Dump the SQL file `toDoListSql.sql` into the database using the following terminal command:
   ```bash
   psql -U <username> -d <database_name> -f toDoListSql.sql
* For example:
    ```bash
    psql -U postgres -d postgres -f toDoListSql.sql

3. navigate it to be 
    ```bash
    "cd/be"

4. in terminal write it  
    ```bash
    * go mod tidy && go mod 


5. run it using 
    ```bash
    *go run main.go

## Front end setup:
1. navigate to the frontend folder 
* ```bash
   cd fe
2. write it syntax on terminal 
    ```bash
     npm install
3. run it using 
    ```bash
    npm run dev 

# using docker:

## just write this syntax on root folder 
    docker-compose up --build


