run:
	rerun -watch . -run sh -c 'go run main.go'

watch:
	watch -t -n 1 curl 'http://localhost:3333/first/second'

