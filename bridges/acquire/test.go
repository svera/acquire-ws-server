package acquirebridge

/*
import (
	"github.com/svera/acquire"
	"github.com/svera/acquire/board"
	"github.com/svera/acquire/fsm"
	"github.com/svera/acquire/interfaces"
	"github.com/svera/acquire/tile"
	"github.com/svera/acquire/tileset"
)

func (b *AcquireBridge) NewGameMergeTest() {
	bd := board.New()
	ts := tileset.New()
	b.corporations = createCorporations()
	tiles := []interfaces.Tile{
		tile.New(5, "E"),
		tile.New(6, "E"),
	}
	tiles2 := []interfaces.Tile{
		tile.New(8, "E"),
		tile.New(9, "E"),
		tile.New(10, "E"),
	}

	ts.DiscardTile(tiles[0])
	ts.DiscardTile(tiles[1])
	ts.DiscardTile(tiles2[0])
	ts.DiscardTile(tiles2[1])
	ts.DiscardTile(tiles2[2])
	bd.SetOwner(b.corporations[0], tiles)
	bd.SetOwner(b.corporations[1], tiles2)
	b.corporations[0].Grow(2)
	b.corporations[1].Grow(3)

	b.game, _ = acquire.New(
		bd,
		b.players,
		b.corporations,
		ts,
		&fsm.PlayTile{},
	)

	b.players[0].DiscardTile(b.players[0].Tiles()[0])
	b.players[0].PickTile(tile.New(7, "E"))
	b.players[0].AddShares(b.corporations[0], 5)
	b.players[1].AddShares(b.corporations[0], 5)
}

func (b *AcquireBridge) NewGameTiedMergeTest() {
	bd := board.New()
	ts := tileset.New()
	b.corporations = createCorporations()
	tiles := []interfaces.Tile{
		tile.New(4, "E"),
		tile.New(5, "E"),
		tile.New(6, "E"),
	}
	tiles2 := []interfaces.Tile{
		tile.New(8, "E"),
		tile.New(9, "E"),
		tile.New(10, "E"),
	}

	ts.DiscardTile(tiles[0])
	ts.DiscardTile(tiles[1])
	ts.DiscardTile(tiles[2])
	ts.DiscardTile(tiles2[0])
	ts.DiscardTile(tiles2[1])
	ts.DiscardTile(tiles2[2])
	bd.SetOwner(b.corporations[0], tiles)
	bd.SetOwner(b.corporations[1], tiles2)
	b.corporations[0].Grow(3)
	b.corporations[1].Grow(3)

	b.game, _ = acquire.New(
		bd,
		b.players,
		b.corporations,
		ts,
		&fsm.PlayTile{},
	)

	b.players[0].DiscardTile(b.players[0].Tiles()[0])
	b.players[0].PickTile(tile.New(7, "E"))
	b.players[0].AddShares(b.corporations[0], 5)
	b.players[1].AddShares(b.corporations[0], 5)
	b.players[0].AddShares(b.corporations[1], 3)
	b.players[1].AddShares(b.corporations[1], 3)
}

func (b *AcquireBridge) NewGameMultiMergeTest() {
	bd := board.New()
	ts := tileset.New()
	b.corporations = createCorporations()
	tiles := []interfaces.Tile{
		tile.New(5, "E"),
		tile.New(6, "E"),
	}
	tiles2 := []interfaces.Tile{
		tile.New(8, "E"),
		tile.New(9, "E"),
		tile.New(10, "E"),
	}
	tiles3 := []interfaces.Tile{
		tile.New(7, "C"),
		tile.New(7, "D"),
	}
	tiles4 := []interfaces.Tile{
		tile.New(7, "F"),
		tile.New(7, "G"),
	}

	ts.DiscardTile(tiles[0])
	ts.DiscardTile(tiles[1])
	ts.DiscardTile(tiles2[0])
	ts.DiscardTile(tiles2[1])
	ts.DiscardTile(tiles2[2])
	ts.DiscardTile(tiles3[0])
	ts.DiscardTile(tiles3[1])
	ts.DiscardTile(tiles4[0])
	ts.DiscardTile(tiles4[1])
	bd.SetOwner(b.corporations[0], tiles)
	bd.SetOwner(b.corporations[1], tiles2)
	bd.SetOwner(b.corporations[2], tiles3)
	bd.SetOwner(b.corporations[3], tiles4)
	b.corporations[0].Grow(2)
	b.corporations[1].Grow(3)
	b.corporations[2].Grow(2)
	b.corporations[3].Grow(2)

	b.game, _ = acquire.New(
		bd,
		b.players,
		b.corporations,
		ts,
		&fsm.PlayTile{},
	)

	b.players[0].DiscardTile(b.players[0].Tiles()[0])
	b.players[0].PickTile(tile.New(7, "E"))
	b.players[0].AddShares(b.corporations[0], 5)
	b.players[1].AddShares(b.corporations[0], 5)
	b.players[0].AddShares(b.corporations[2], 2)
	b.players[1].AddShares(b.corporations[2], 2)
	b.players[0].AddShares(b.corporations[3], 2)
	b.players[1].AddShares(b.corporations[3], 2)
}

func (b *AcquireBridge) NewGameEndTest() {
	bd := board.New()
	ts := tileset.New()
	b.corporations = createCorporations()
	tiles := []interfaces.Tile{
		tile.New(5, "E"),
		tile.New(6, "E"),
		tile.New(7, "E"),
		tile.New(8, "E"),
		tile.New(9, "E"),
		tile.New(10, "E"),
		tile.New(7, "C"),
		tile.New(7, "D"),
		tile.New(7, "F"),
		tile.New(7, "G"),
		tile.New(7, "H"),
	}

	for i := range tiles {
		ts.DiscardTile(tiles[i])
	}
	bd.SetOwner(b.corporations[0], tiles)
	b.corporations[0].Grow(11)

	b.game, _ = acquire.New(
		bd,
		b.players,
		b.corporations,
		ts,
		&fsm.PlayTile{},
	)

	b.players[0].AddShares(b.corporations[0], 5)
	b.players[1].AddShares(b.corporations[0], 5)
	b.players[0].AddShares(b.corporations[2], 2)
	b.players[1].AddShares(b.corporations[2], 2)
	b.players[0].AddShares(b.corporations[3], 2)
	b.players[1].AddShares(b.corporations[3], 2)
}

func (b *AcquireBridge) NewGameAllCorpsOnBoardTest() {
	corpsTiles := make(map[int][]interfaces.Tile)

	corpsTiles[0] = []interfaces.Tile{
		tile.New(1, "A"),
		tile.New(1, "B"),
	}
	corpsTiles[1] = []interfaces.Tile{
		tile.New(3, "A"),
		tile.New(3, "B"),
	}
	corpsTiles[2] = []interfaces.Tile{
		tile.New(5, "A"),
		tile.New(5, "B"),
	}
	corpsTiles[3] = []interfaces.Tile{
		tile.New(7, "A"),
		tile.New(7, "B"),
	}
	corpsTiles[4] = []interfaces.Tile{
		tile.New(9, "A"),
		tile.New(9, "B"),
	}
	corpsTiles[5] = []interfaces.Tile{
		tile.New(11, "A"),
		tile.New(11, "B"),
	}
	corpsTiles[6] = []interfaces.Tile{
		tile.New(1, "D"),
		tile.New(1, "E"),
	}

	unincorporatedTiles := []interfaces.Tile{
		tile.New(3, "D"),
		tile.New(5, "D"),
		tile.New(7, "D"),
		tile.New(9, "D"),
		tile.New(11, "D"),
		tile.New(3, "F"),
		tile.New(5, "F"),
		tile.New(7, "F"),
		tile.New(9, "F"),
		tile.New(11, "F"),
	}

	for i := range unincorporatedTiles {
		b.board.PutTile(unincorporatedTiles[i])
	}

	for i := range corpsTiles {
		b.tileset.(*tileset.Tileset).DiscardTile(corpsTiles[i][0])
		b.tileset.(*tileset.Tileset).DiscardTile(corpsTiles[i][1])
		b.board.SetOwner(b.corporations[i], corpsTiles[i])
		b.corporations[i].Grow(len(corpsTiles[i]))
	}
}
*/
