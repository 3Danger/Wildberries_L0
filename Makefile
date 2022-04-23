
APP=Service
PUB=PublisherServ

all: $(APP) $(PUB)

$(APP):
	go build -o $@ main.go

$(PUB):
	go build -o $@ Publisher/publisher.go

fclean:
	rm $(APP) $(PUB)
	bash close.sh
clean:
	bash close.sh

run: $(APP) $(PUB)
	bash run.sh
	
