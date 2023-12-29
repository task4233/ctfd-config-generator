setup:
	python -m pip install ctfcli

gen/build:
	cd cmd/generator && go build -o ../../gen

gen: gen/build
	./gen

ctfcli/init:
	python -m ctfcli init
	./cmd/scripts/install.sh

ctfcli/sync:
	./cmd/scripts/sync.sh

