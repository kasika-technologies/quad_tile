# quad_tile

Golang implementation of QuadTile to work with OpenStreetMap API database.

## Usage

```go
package main

import (
	"fmt"
	"github.com/kasika-technologies/quad_tile"
)

func main() {
	tile := quad_tile.TileForPoint(139.734531, 35.607090)
	fmt.Println(tile)
	// 3977143975

	tiles := quad_tile.TilesForArea(139.73424, 35.60433, 139.73673, 35.60707)
	fmt.Println(tiles)
	// [3977143974 3977143975]
}
```
