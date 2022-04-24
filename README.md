# Wildberries_L0

## Параметры Базы данных для того что бы запустить это:
```
  username database = [csamuro]
  name of database = [csamuro]
```
___  

## Таблица в базе данных следующего вида:
```
CREATE TABLE models
(
  id SERIAL PRIMARY KEY,
  model JSONB NOT NULL
);
```
___
## Требования!

### У вас должен быть установлен:
```
    1) nats-streaming-server - Брокер сообщений
    2) База данных должна соответствовать требованиям приведенное выше!
    3) go - компилятор что бы скомпилировать это.
```
___ 

### Если с требованиями выше все ОК, то запускаем программу через скрипт следующим образом
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