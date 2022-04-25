
APP=Service
PUB=PublisherServ

all: $(APP) $(PUB)

$(APP):
	go build -o $@ main.go

$(PUB):
	go build -o $@ Publisher/publisher.go

fclean: clean
	@rm $(APP) $(PUB) 2> /dev/null &
	@echo "$(APP) и $(PUB) удалены"

clean:
	@bash close.sh
	@echo "порты 4222 и 8080 освобождены"

run: $(APP) $(PUB)
	bash run.sh
	
