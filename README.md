# Simple local GOPROXY

	docker build -t goproxy .
	docker run -ti --rm -v cache:/cache -p 8080:80 goproxy
	export GOPROXY=http://localhost:8080
