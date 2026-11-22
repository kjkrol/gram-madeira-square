# gram-madeira-square

Madeira on a square grid, a cell 100 m: two demos of the [gram](https://github.com/kjkrol/gram)
engine and the tool that makes their maps of the real island.

| Demo | What it shows |
|---|---|
| [`relief`](relief) | Madeira in relief, its true heights: light, water, the sea over its bed, roads, bridges and tunnels, the weather, the ground textured from photographs (`-look ghibli` paints it), forests and settlements as fields unfolding near and photographed into a sheet far off, the settlements' buildings as solids, caravans and travellers on the roads, the camera on a stick or behind a unit |
| [`flat`](flat) | The same island on the simple map: a flat board, its settlements as cells of their own, bands for the ways, bridges and fords, a minimap round the camera |

```bash
make relief           # -look ghibli: make relief-ghibli; -hud widgets: make relief-widgets
make flat
```

K shows the controls in either, F10 the settings. The demos keep what is dear to make between
runs in the user's cache (`gram-madeira-square/relief`: the sheets of what stands far off and the
hierarchy the long routes are climbed through).

## Maps of the Earth

Both maps are made by [gramap](https://github.com/kjkrol/gramap), which makes a map a gram board
lays of any stretch of the Earth at any scale out of a recipe; `madeira.yaml` is Madeira's:

```bash
make maps
```

which runs gramap at the version the Makefile names (`GRAMAP`):

```bash
go run github.com/kjkrol/gramap@v0.2.1 -recipe madeira.yaml -out relief/map.yaml \
    -geography relief/geography
go run github.com/kjkrol/gramap@v0.2.1 -recipe madeira.yaml -out flat/map.yaml -flat
```

Madeira in relief is laid by its map in metres: `relief/geography/`, the ground's heights as
an image and what lies on it as GeoJSON (gram's `board/mapfile`), and `relief/geography.yaml`,
the level over it. `relief/map.yaml` is the same map as the level file it takes over from,
kept beside it till it can go; the units are written there and carried into
`geography.yaml` as the maps are made.

The recipe says where, how many cells and how many metres a cell spans, what the ground is by
its slope and height, how much drained land makes a brook, a stream and a river, and the roads:
here OpenStreetMap's, with its bridges, tunnels and settlements, the land drawn as CORINE Land
Cover says (laurisilva within its Natura 2000 site, woodland, pinewood, heath and farmland).
What a recipe can say is gramap's README's.

`relief/geography.yaml` and `relief/map.yaml` each carry one line written by hand, the lantern
after the bridge from `[181, 73]` (`{kind: lantern, at: [184, 76]}` among the board's
entities): a map made anew gets it back by hand.

## Building

The module needs gram and its siblings (goke, ceremony, aabbworld, hpastar) as its `go.mod`
says. `make build` and `make test` build and test everything; the Makefile runs with
`GOEXPERIMENT=simd`, as gram's does.

# Credits

- The heights of both maps and of every map gramap makes: the Terrain Tiles on AWS,
  Mapzen's Tilezen terrarium — SRTM, EU-DEM (Copernicus), GMTED2010, ETOPO1 and others; see their
  [attribution](https://github.com/tilezen/joerd/blob/master/docs/attribution.md).
- The roads, tunnels, bridges and settlements of Madeira in relief (`relief/map.yaml`), and of
  every map gramap makes with `source: osm`: ©
  [OpenStreetMap](https://www.openstreetmap.org/copyright) contributors, under the Open Database
  License.
- What grows where on Madeira in relief, and on every map gramap makes with `cover:
  {source: corine}`: CORINE Land Cover 2018, © European Union, Copernicus Land Monitoring
  Service, European Environment Agency (EEA); the Natura 2000 sites, © EEA.
- The ground's photographs (`relief/textures`): Poly Haven, CC0.

# License

[MIT](LICENSE)
