# Проект `smarthouse-service`

Простой веб-сервис, который можно использовать для организации умного дома или [интернета вещей](https://ru.wikipedia.org/wiki/%D0%98%D0%BD%D1%82%D0%B5%D1%80%D0%BD%D0%B5%D1%82_%D0%B2%D0%B5%D1%89%D0%B5%D0%B9).

## Описание

Сервис позволяет работу со следующими объектами (в иерархии):

* Пользователь 1
	* Дом 1
		* Датчик 1
			* Данные
		* Датчик 2
			* Данные
		* ...
		* Команда управления 1
		* Команда управления 2
		* ...
		* Команда запроса 1
		* Команда запроса 2
		* ...
	* Дом 2
	* ...
* Пользователь 2
* ...

## Технологии

* [Go](https://golang.org/) 
* [PostgreSQL](https://www.postgresql.org/)

## API

Данные будут доступны только если запрос содержит заголовок `Authorization` с хешем пользователя и пароля ([Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication)). Проверка на валидность осуществляется запросом:

* `POST host:port/login` - попытка пройти аутентификацию

### Пользователи

Формат объекта в `JSON`:

	{
		"login": "user",
		"password": "password",
		"name": "Вася"
	}

Запросы:

* `GET host:port/user` - возвращает данные по пользователю
* `GET host:port/users` - возвращает данные по пользователям
* `POST host:port/users` - создает пользователя

### Дома

Формат объекта в `JSON`:

	{
		"name": "Квартира",
		"address": "г. Москва, ул. Строителей, д.22"
	}

Запросы:

* `GET host:port/users/{user_id}/houses` - возвращает список домов пользователя
* `POST host:port/users/{user_id}/houses` - добавляет дом
* `PUT host:port/users/{user_id}/houses/{house_id}` - редактирует дом
* `DELETE host:port/users/{user_id}/houses/{house_id}` - удаляет дом

### Датчики

Формат объекта в `JSON`:

	{
		"name": "Датчик температуры",
		"tag": "temperature_sensor_1"
	}

Запросы:

* `GET host:port/users/{user_id}/houses/{house_id}/sensors` - возвращает список датчиков дома
* `POST host:port/users/{user_id}/houses/{house_id}/sensors` - добавляет датчик в дом
* `PUT host:port/users/{user_id}/houses/{house_id}/sensors/{sensor_id}` - редактирует датчик дома
* `DELETE host:port/users/{user_id}/houses/{house_id}/sensors/{sensor_id}` - удаляет датчик из дома

### Данные датчиков

Формат объекта в `JSON`:

	{
		"data": "...например JSON с данными..."
		"time": "2017-03-03 12:02:33.2345+3000" 
	}

Поле `time` доступно только для чтения и генерируется автоматически.

Запросы:

* `PUT host:port/users/{user_id}/houses/{house_id}/sensors/{sensor_id}/sensordata` - возвращает данные датчика
* `POST host:port/users/{user_id}/houses/{house_id}/sensors/{sensor_id}/sensordata` - записывает данные для датчика. В качестве фильтра в тело запроса можно поместить `JSON` со следующими свойствами:	
	* `after` - дата и время первого показания
	* `before` - дата и время последнего показания

* `POST host:port/sensordata/{sensor_tag}` - записывает данные датчика

### Команды управления

Еще не реализовано

### Команды запроса

Еще не реализовано