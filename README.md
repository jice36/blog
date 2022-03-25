# blog
порт сервиса с базой данных 5300
db - 5434 <br/>
blog-8181
 сборка проекта <br/>
  make<br/>
  ./main<br/>
  ./db<br/>
  
 сборка docker <br/>
 docker build -t blog-db ./<br/>
 docker run -d --name blog-db-container -p 5434:5432 blog-db <br/>

в users уже добавленны пользователи <br/>
b03d13da-ab8a-11ec-90e5-acde48001122<br/>
b03d1916-ab8a-11ec-90e5-acde48001122<br/>
b03d1934-ab8a-11ec-90e5-acde48001122<br/>
b03d1948-ab8a-11ec-90e5-acde48001122<br/>
b03d1952-ab8a-11ec-90e5-acde48001122<br/>
b03d1966-ab8a-11ec-90e5-acde48001122<br/>
b03d1970-ab8a-11ec-90e5-acde48001122<br/>
b03d1984-ab8a-11ec-90e5-acde48001122<br/>
b03d198e-ab8a-11ec-90e5-acde48001122<br/>
b03d19a2-ab8a-11ec-90e5-acde48001122<br/>


GET http://localhost:8181/blog/b03d1916-ab8a-11ec-90e5-acde48001122<br/>

POST http://localhost:8181/blog<br/>
Content-Type: application/json<br/>

{<br/>
  "user_id": "b03d1916-ab8a-11ec-90e5-acde48001122",<br/>
 "header":"head",<br/>
 "text":"text",<br/>
 "tags":["one","two"]<br/>
}<br/>
