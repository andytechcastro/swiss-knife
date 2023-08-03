test:
	go test ./kubernetes/builders -v -cover && go test ./kubernetes/actions -v -cover
