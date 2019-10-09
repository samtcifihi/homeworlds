compile:
	go build -o "Homeworlds"

test: compile
	./Homeworlds

clean:
	rm Homeworlds
