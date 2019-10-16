# upload-images
How to setup and launch project.

I using technical following:
  + Gin framework to build the REST api services 
  + PostgreSQL DB to save the user data, 
  + Beego ORM for making queries to the PostgreSQL DB

Install the following dependencies.

  +go get -u github.com/gin-gonic/gin
  +go get github.com/astaxie/beego/orm
  +go get github.com/lib/pq

I using docker-compose for database.
 +Run: cd docker && docker-compose up -d
 +Create table using pgAdmin: http://localhost:5050
       create table images (
          id serial PRIMARY KEY,
          file_name varchar(255) NOT NULL,
          size integer NOT NULL,
          content_type varchar(100) NOT NULL
    );

Run project with token (example: token=123):
 token=123 go run main.go
 



