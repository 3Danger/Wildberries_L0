#!/usr/bin/bash/

fuser -k 4222/tcp && fuser -k 8080/tcp &> /dev/null

sleep 1

nats-streaming-server &> /dev/null &

sleep 2
echo "Запускаю Сервис"
sleep 1

go run main.go &

sleep 2
echo "Сейчас будем публиковать рандомные Json файлы"
echo "_3_"
sleep 1
echo "_2_"
sleep 1
echo "_1_"
sleep 1
echo "Поехали! запускаю "
sleep 1

./testPublisher.sh &> /dev/null

echo "Сервис работает в фоновом режиме.
Теперь можно тестировать в браузере по адресу:
http://localhost:8080/"
