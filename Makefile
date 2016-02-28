build: main.go
	go build

test:
	go test ./...

assets: frontend/Gulpfile.js frontend/package.json
	cd frontend; npm install; gulp
	make assetsbin

getdeps:
	go get -u -v

assetsbin:
	go-bindata -pkg bindata -nomemcopy -o bindata/bindata.go frontend/static/ frontend/templates/
