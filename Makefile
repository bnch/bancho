build: main.go
	go build

test:
	go test ./...

assets: frontend/Gulpfile.js frontend/package.json
	cd frontend; npm install; gulp
