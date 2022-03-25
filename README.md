# blog
порт сервиса с базой данных 5300
db - 5434 <br/>
blog-8181
 сборка проекта 
  make
  ./main
  ./db
  
 сборка docker 
 docker build -t blog-db ./
 docker run -d --name blog-db-container -p 5434:5432 blog-db

в users уже добавленны пользователи 
b03d13da-ab8a-11ec-90e5-acde48001122
b03d1916-ab8a-11ec-90e5-acde48001122
b03d1934-ab8a-11ec-90e5-acde48001122
b03d1948-ab8a-11ec-90e5-acde48001122
b03d1952-ab8a-11ec-90e5-acde48001122
b03d1966-ab8a-11ec-90e5-acde48001122
b03d1970-ab8a-11ec-90e5-acde48001122
b03d1984-ab8a-11ec-90e5-acde48001122
b03d198e-ab8a-11ec-90e5-acde48001122
b03d19a2-ab8a-11ec-90e5-acde48001122


GET http://localhost:8181/blog/b03d1916-ab8a-11ec-90e5-acde48001122

POST http://localhost:8181/blog
Content-Type: application/json

{
  "user_id": "b03d1916-ab8a-11ec-90e5-acde48001122",
 "header":"head",
 "text":"text",
 "tags":["one","two"]
}
