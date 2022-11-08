dev:
	go run .

deploy:
	fly deploy --strategy bluegreen
