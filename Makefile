run:
	air
doc:
	godoc -http=:6060
push:
	git push origin main
mod-refresh:
	go mod tidy