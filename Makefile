compile:
	go build ./src/userio
	go build -o "Homeworlds" ./src/main

test: compile
	./Homeworlds

clean:
	rm Homeworlds
