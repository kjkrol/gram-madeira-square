# gram-madeira-square

Madeira on a square grid, a cell 100 m: two demos of the [gram](https://github.com/kjkrol/gram)
engine and the tool that makes their maps of the real island.

| Demo | What it shows |
|---|---|
| [`relief`](relief) | Madeira in relief, its true heights: light, water, the sea over its bed, roads, bridges and tunnels, the weather, the ground textured from photographs (`-look ghibli` paints it), forests and settlements as fields unfolding near and photographed into a sheet far off, the settlements' buildings as solids, caravans and travellers on the roads, the camera on a stick or behind a unit |
| [`flat`](flat) | The same island on the simple map: a flat board, bands for the ways, bridges and fords, a minimap round the camera |

```bash
make relief           # -look ghibli: make relief-ghibli; -hud widgets: make relief-widgets
make flat
```

K shows the controls in either, F10 the settings. The demos keep what is dear to make between
runs in the user's cache (`gram-madeira-square/relief`: the sheets of what stands far off and the
hierarchy the long routes are climbed through).

## Maps of the Earth

`cmd/gramap` makes a map a gram board lays of any stretch of the Earth at any scale out of a
recipe (`cmd/gramap/madeira.yaml` is Madeira's):

```bash
go run ./cmd/gramap -recipe cmd/gramap/madeira.yaml -out relief/map.yaml
go run ./cmd/gramap -recipe cmd/gramap/madeira.yaml -out flat/map.yaml -flat
```

The recipe says where, how many cells and how many metres a cell spans, what the ground is by
its slope and height, how much drained land makes a brook, a stream and a river, and the roads
as chains of named places. The heights come from the
[Terrain Tiles on AWS](https://registry.opendata.aws/terrain-tiles/), fetched once into the
user's cache. The rivers are carved into the ground, the roads routed and graded as Transport
Tycoon lays them, with bridges over water and down valleys and fords where a river is shallow;
with `-flat` the map has no heights. With `bed: {zoom, deepest}` the sea has its bed, as deep
as the tiles tell at a zoom that has the sea's depths (the fine ones have none), evened from the
shore to there. With `source: osm` a map with relief follows
[OpenStreetMap](https://www.openstreetmap.org/copyright) instead: its roads of the classes the
recipe names, its bridges, its tunnels dug where the ground lets one lie, and its settlements
where the buildings stand. With `cover: {source: corine}` it draws the land as
[CORINE Land Cover](https://land.copernicus.eu/en/products/corine-land-cover) says, a kind of
cell for the classes the recipe names, another within a Natura 2000 site (Madeira: laurisilva
within its site, woodland, pinewood, heath and farmland). Its `fields:` lay a field entity over
the cells of a kind, or of several, in each block where enough of them lie, its `settlements:`
cells of a kind round the roads' places, as many as a place's size says in the recipe's own
words (a hamlet, a village, a town); what grows and stands on them is the game's rules', as the
fields unfold.

`relief/map.yaml` carries one line written by hand, the lantern after the bridge from
`[181, 73]`: a map made anew gets it back by hand.

## Building

The module needs gram and its siblings (goke, ceremony, aabbworld, hpastar) as its `go.mod`
says. `make build` and `make test` build and test everything; the Makefile runs with
`GOEXPERIMENT=simd`, as gram's does.

# Credits

- The heights of both maps and of every map `cmd/gramap` makes: the Terrain Tiles on AWS,
  Mapzen's Tilezen terrarium — SRTM, EU-DEM (Copernicus), GMTED2010, ETOPO1 and others; see their
  [attribution](https://github.com/tilezen/joerd/blob/master/docs/attribution.md).
- The roads, tunnels, bridges and settlements of Madeira in relief (`relief/map.yaml`), and of
  every map `cmd/gramap` makes with `source: osm`: ©
  [OpenStreetMap](https://www.openstreetmap.org/copyright) contributors, under the Open Database
  License.
- What grows where on Madeira in relief, and on every map `cmd/gramap` makes with `cover:
  {source: corine}`: CORINE Land Cover 2018, © European Union, Copernicus Land Monitoring
  Service, European Environment Agency (EEA); the Natura 2000 sites, © EEA.
- The ground's photographs (`relief/textures`): Poly Haven, CC0.

# License

[MIT](LICENSE)
