# Roadmap

What is left to do in the Madeira demos and in gramap, and only that: an item goes out as it lands.
What is left in the engine's plugins is in gram's own `doc/roadmap.md`.

- **The sea's bed off the south coast** (2026-11-16) — the coarser tiles tell land up to 1.5 km past
  the coast, so Madeira's steep south has shallows wider than in life.
- **Madeira's buildings, what is left** (2026-11-16) — a cottage, a town's house and a chapel of 6
  by 4 m, no wider than the island's world holds a thing (`Entities.MaxSize` a cell) and turned by
  quarters alone.
- **Maps from OpenStreetMap, what is left** (2026-11-11) — Madeira at 100 m lays three quarters of
  its roads' length: 308 stretches of four to seven cells are left out where a road of
  `ground.RoadSlope` (10°) does not fit a hundred metres of that ground (hairpins tighter than a
  cell are gone), so the network has gaps a unit crosses on bare ground; of 114 tunnels 25 lie with
  a road at either head, 7 are dug and left out for no road could be laid to a head cut into the
  ground (a road planned from the head outward would reach them), and 82 stay roads, most for a room
  of 45 m (figures 15 times life) the ground does not close over within the reshape allowed, or a
  head beside running water, which never moves; a tunnel shorter than four cells and a bridge
  shorter than one are plain road. A head drawn outward along its road, a steeper class of road (a
  mountain track) and the map of 10 m with figures as in life are the ways on. gramap asks one
  server and its mirror and names the left-out stretches only by their count; the flat demo's
  committed map predates the settlements and was not made anew.
- **What grows where, what is left** (2026-11-11) — the kinds of growth are CORINE's classes at 25
  ha a shape, the species within a kind picked by the height of the ground alone (no open map of
  Madeira's species exists); the map of 10 m wants ESA WorldCover's 10 m tree cover for the edges.
  Madeira's start view draws 8 % slower with the growth over 55 000 cells where the forest took 19
  000: a third of it follows the number of trees, the rest was not found (the sheets holding more
  sectors is the guess). The flat demo keeps its forest by noise.
- **Buildings of many kinds** (stage 0d, to plan) — OpenStreetMap has 113 762 buildings on Madeira,
  some 13 000 with a type and 11 745 with their levels: the kind of building a cell by the tags, 34
  kinds proposed (dwellings, tourism and trade, public, industry and infrastructure) and 18
  landmarks in their true places.
- **The world is full on a big screen** (2026-11-17) — the relief demo's world holds 40 000 entities
  (`MaxEntCount`); a big window zoomed in unfolds more sectors than that, the world refuses the
  parts past it (a log line a part) and the far sectors stand without some of their trees and of
  their buildings' entities. A higher cap, or the full levels held to a ring round the camera
  (`lod.Ladder.Rings`).
