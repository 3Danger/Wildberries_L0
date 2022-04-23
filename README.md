# Wildberries_L0

Параметры Базы данных для того что бы запустить это:
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

## Требования!

### У вас должен быть установлен:
```
    1) go - кампилятор что бы скампилировать это.
    2) nats-streaming-server - Брокер сообщений
    3) База данных должна соответствовать требованиям приведенное выше!
```
    

### Если с требованиями выше все ОК, то запускаем программу через скрипт следующим образом
```
git clone https://github.com/3Danger/Wildberries_L0.git
cd Wildberries_L0
bash run.sh
```
