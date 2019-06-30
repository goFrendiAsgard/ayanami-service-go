test:
	go test -race ./... -coverprofile=profile.out -count=1 -covermode=atomic

testv:
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic

coverage:
	go tool cover -html=profile.out
