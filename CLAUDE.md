# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this
repository.

## What this is

**gram-madeira-square** (module `github.com/kjkrol/gram-madeira-square`) is two demos of the
[gram](https://github.com/kjkrol/gram) engine on the real Madeira, a square grid of cells 100 m
across. They were gram's `examples/board/topography` and `examples/board/flat` till gram v0.22.0.
Their maps are made by [gramap](https://github.com/kjkrol/gramap) (locally
`/home/karol/CODE/gramap`, where its skill is) from the Terrain Tiles, OpenStreetMap and CORINE
Land Cover, as `madeira.yaml` says.

- `relief/` — Madeira in relief (`demo.go` the Stage, `scene.go` how it looks, `names.go` its
  names, a file a figure, `hud/` the hour and the weather on the glass, `geography.yaml` and
  `geography/` the map).
- `flat/` — the same island on a flat board.
- `madeira.yaml` — the recipe of both maps, gramap's to read.

gram lives beside it (`/home/karol/CODE/gram`, joined by `/home/karol/CODE/go.work` with goke,
ceremony, aabbworld and hpastar): **the rules of the house, the skills of the engine's layers and
the way a demo is written are gram's** (`gram/CLAUDE.md`, `gram/.claude/skills`), read there and
never repeated here.

## Commands

```bash
go build ./... && go vet ./... && gofmt -l . && go test ./...   # the verification sequence
make relief                                                     # Madeira in relief
make relief-ghibli                                              # the same ground painted (-look ghibli)
make relief-widgets                                             # the hour and the weather by the corners (-hud widgets)
make flat                                                       # Madeira on the simple map
make maps                                                       # both maps made anew from madeira.yaml, by gramap at the Makefile's version
```

The Makefile runs with `GOEXPERIMENT=simd`, as gram's. The demos are GUI programs: `go test`
does not exercise them; one is checked by building it and running it under `timeout` (exit 124
is healthy), by a shot (`GRAM_SHOT=<file.png>`, gram's variable) or with gram's `gram-probe`
scripts pointed at this repository (`PROBE_DEMOS=<this folder>`, the demo dir `relief` or
`flat`); `.claude/probes/picture` is the probe that sets the relief demo's tactical picture.

## What to know

- A change to a gram plugin these demos use is checked here too: the relief demo is the one
  place the ground, the sky, the forests and the towns run at an island's scale.
- Each demo lays its `geography.yaml` over its `geography/`, the map in metres (gram's
  `board/mapfile`, written by gramap); the units are written in `geography.yaml` and kept by
  gramap as the maps are made, so its world says its scale. The relief's carries one line
  written by hand (the lantern after the bridge from `[181, 73]`): put it back after
  `make maps`. gramap holds itself to Madeira made from
  its own copy of this recipe (`make madeira` there): a change to `madeira.yaml` here moves
  nothing there.
- What the relief demo keeps between runs (the sheets, the routes' hierarchy) lies in the
  user's cache under `gram-madeira-square/relief`; a change to how fields are planted or steps
  are priced is a new `SheetsVersion` or `RoutesVersion` in `relief/names.go`.
- The relief demo's world holds 40 000 entities (`MaxEntCount`): a big screen zoomed in unfolds
  more sectors than that and the world refuses the parts past it (a log line a part).
- Commits, pull requests, docs and comments follow gram's rules; the user tags and merges.
