build-guest-auth:
	docker build -f app/auth/guest/Dockerfile . -t guest_authentication:local

build-game-start:
	docker build -f app/game/start/Dockerfile . -t game_initialization:local

build: build-guest-auth build-game-start
