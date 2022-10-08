.PHONY: test clean

test: deploy.done
	curl -fsSL -D - "$$(terraform output -raw url)?name=Lambda"

clean:
	terraform destroy
	rm -f init.done deploy.done coinflip.zip coinflip

init.done:
	terraform init

deploy.done: init.done main.tf coinflip.zip
	terraform apply

coinflip.zip: coinflip
	zip $@ $<

coinflip: main.go
	go get .
	GOOS=linux GOARCH=amd64 go build -o $@
