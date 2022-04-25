#!/bin/bash/

fuser -k 4222/tcp &> /dev/null
fuser -k 8080/tcp &> /dev/null

sleep 1

echo "Запускаю nats-streaming-serv"
nats-streaming-server &> /dev/null &

sleep 2
echo "Запускаю наш Service"
sleep 1

./Service &

sleep 2
echo
echo "Кэш сейчас возможно пустой, поэтому"
echo "нажмите CTRL+D что бы PublisherServ"
echo "начал что-ни-будь публиковать в канал"
cat ;

echo
echo "Сейчас будем публиковать рандомные Json файлы"
echo "_3_"
sleep 1
echo "_2_"
sleep 1
echo "_1_"
sleep 1
echo "Поехали! запускаю "
sleep 1

./PublisherServ -ms 50 -r 10 -j ./json &> /dev/null

sleep 1

echo "Сервис работает в фоновом режиме.
Теперь можно тестировать в браузере по адресу:
http://localhost:8080/"
