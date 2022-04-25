# Wildberries L0

### Параметры Базы данных для того что бы запустить это:
```
  username database = [csamuro]
  name of database = [csamuro]
```
___  

### Таблица в базе данных следующего вида:
```
CREATE TABLE models
(
  id SERIAL PRIMARY KEY,
  model JSONB NOT NULL
);
```
___

### Запускаем следующим образом
```
$ git clone https://github.com/3Danger/Wildberries_L0.git
$ cd Wildberries_L0
$ make run
```
___
### По завершению использования закрываем соединение
```
$ make clean
```

___