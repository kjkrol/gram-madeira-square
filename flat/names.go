package main

const (
	WaterCell  = "water"
	EarthCell  = "earth"
	SandCell   = "sand"
	RockCell   = "rock"
	ForestCell = "forest"
	BrookCell  = "brook"
	StreamCell = "stream"
	RiverCell  = "river"
	RoadCell   = "road"
	BuiltCell  = "built"
	IceCell    = "ice"

	MortalRole = "mortal"

	UnitKind = "unit"

	BridgeKind     = "bridge"
	FordKind       = "ford"
	ForestKind     = "forest"
	SettlementKind = "settlement"

	PlayerSide = "player"

	BoardStage = "board"
	MainScene  = "main"
)

func snowyCell(kind string) string { return "snowy " + kind }
