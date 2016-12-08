go=go
cd=@cd

build:
	$(go) build -o mars

clean:
	$(go) clean

run:
	$(go) run main.go

test:
	$(cd) position && $(go) test
	$(cd) rover && $(go) test
	$(cd) plateau && $(go) test

install:
	$(go) install
