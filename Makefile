compile:
	go build src/userio
	go build src/main -o "Homeworlds"

test: compile
	./Homeworlds

clean:
	rm Homeworlds
