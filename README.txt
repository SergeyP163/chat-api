# Это первый проект, так что он может быть кривым, но возможно будет дорабатываться



# pgAdmin(PostgreSQL)
db name = chat_db
user = chat_user
chat_password = chat_password

# Запуск через Docker
docker-compose up --build

# Использовал PowerShell (На Windows PowerShell лучше использовать Invoke-WebRequest)
    - Создать чат: Invoke-WebRequest -Uri "http://localhost:8080/chats" -Method POST -ContentType "application/json" -Body '{"title":"Мой первый чат"}'
    - Отправить сообщение: Invoke-WebRequest -Uri "http://localhost:8080/chats/1/messages" -Method POST -ContentType "application/json" -Body '{"text":"Привет, это первое сообщение"}'
    - Получение чата с последними сообщениями: Invoke-WebRequest -Uri "http://localhost:8080/chats/1?limit=20" -Method GET
    - Удалить чат: Invoke-WebRequest -Uri "http://localhost:8080/chats/1" -Method DELETE

# Разное (что хочу добавить/изменить)
    - так как в интернете (http://localhost:8080/) не откерывается (404 page not found), то я хочу это исправить;
    - хочу сделать нормальное формление сайта (думаю нужны знания HTML);
    
