all:
	$(MAKE) clean && $(MAKE) build && $(MAKE) test && $(MAKE) install

build:
	cd position && $(MAKE) build
	cd rover && $(MAKE) build
	cd plateau && $(MAKE) build

clean:
	cd position && $(MAKE) clean
	cd rover && $(MAKE) clean
	cd plateau && $(MAKE) clean

install:
	cd position && $(MAKE) install
	cd rover && $(MAKE) install
	cd plateau && $(MAKE) install

test:
	cd position && $(MAKE) test
	cd rover && $(MAKE) test
	cd plateau && $(MAKE) test

