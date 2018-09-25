all:
	$(MAKE) -C src all

build:
	$(MAKE) -C src build

clean:
	$(MAKE) -C src clean
	if [ -f coverage.html ] ; then rm coverage.html ; fi
	if [ -d .cover ] ; then rm -rf .cover ; fi
	docker-compose down --rmi all -v 2>/dev/null || true
	docker-compose stop >/dev/null
	docker-compose rm >/dev/null

rebuild:
	docker-compose build PROJECTNAME
	docker-compose build unit

unit:
	docker-compose run --rm unit

coverage:
	docker-compose run --rm unit && [ -f ./coverage.html ] && xdg-open coverage.html

.PHONY: all build clean unit rebuild coverage
