test:
	go test --race --count=1 -coverprofile coverage.out $$(go list ./...)