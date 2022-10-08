build-authentication-guest:
	docker build -f app/authentication/guest/Dockerfile . -t authentication_guest:local

build-advancewars-battlemap-create:
	docker build -f app/advancewars/battlemapcreate/Dockerfile . -t advancewars_battlemap_create:local

build-authentication: build-authentication-guest

build-advancewars: build-advancewars-battlemap-create

build: build-authentication build-advancewars
