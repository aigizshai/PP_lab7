# PP_lab7
1.	Создание TCP-сервера:
2.	Реализация TCP-клиента:
3.	Асинхронная обработка клиентских соединений:
Для запуска tcp сервера go run tcpServer.go, затем запуск клиента go run tcpClient.go

4.	Создание HTTP-сервера:
5.	Добавление маршрутизации и middleware:
Для запуска http сервера go run httpServer.go.
для клиента используется curl
пример post запроса curl -X POST localhost:8080/data -d '{ "name" : "hello" }'
пример get запроса curl localhost:8080/hello

6.	Веб-сокеты:
Для запуска websocket сервера go run webSocket.go
Запуск клиентов websocket go run client.go
