TINYGO=../tinygo

all: import

import:
	sh import.sh $(TINYGO)

dist:	dist.tar

dist.tar:
	sh mkdist.sh
	cd dist && tar cf ../$@ .
	rm -rf dist

clean:
	rm -rf dist

.PHONY: \
	clean\
	dist\
	import\
