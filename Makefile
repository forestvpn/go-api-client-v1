validate:
	swagger validate api.yaml

client: validate
	swagger generate client --default-scheme https -f api.yaml -A api -t ./

all: client
