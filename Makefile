
APP=Service
PUB=PublisherServ

all: $(APP) $(PUB)

$(APP):
	go build -o $@ main.go

$(PUB):
	go build -o $@ Publisher/publisher.go

fclean:
	bash close.sh
	rm $(APP) $(PUB)
clean:
	@bash close.sh

run: $(APP) $(PUB)
	bash run.sh
	
