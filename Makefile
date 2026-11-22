GO = go

# every build, test and run with the simd experiment, as gram's own Makefile has it: aabbworld's
# pair sweep eight boxes at a time on a CPU with AVX2; GOEXPERIMENT= on the command line goes
# without it
GOEXPERIMENT ?= simd
export GOEXPERIMENT

# the map maker, a tool of its own
GRAMAP ?= github.com/kjkrol/gramap@v0.2.1

.PHONY: all relief relief-ghibli relief-widgets flat maps deps tidy build test clean

all: build

## relief: Madeira in relief, a cell 100 m
relief: deps
	$(GO) run ./relief
relief-ghibli: deps
	$(GO) run ./relief -look ghibli
relief-widgets: deps
	$(GO) run ./relief -hud widgets

## flat: Madeira on the simple map, a flat board
flat: deps
	$(GO) run ./flat

## maps: makes the maps anew from the recipe, the relief's in metres too (relief/geography and
## relief/geography.yaml, what the demo lays); relief/geography.yaml and relief/map.yaml each
## carry one line written by hand (the lantern after the bridge from [181, 73]), to be put back
maps:
	$(GO) run $(GRAMAP) -recipe madeira.yaml -out relief/map.yaml -geography relief/geography
	$(GO) run $(GRAMAP) -recipe madeira.yaml -out flat/map.yaml -flat

deps:
	$(GO) mod tidy

tidy:
	$(GO) mod tidy

build:
	$(GO) build ./...

test:
	$(GO) test ./...

clean:
	$(GO) clean
