package quad_tile

import (
	"reflect"
	"testing"
)

func TestTileForPoint(t *testing.T) {
	got := TileForPoint(139.734531, 35.607090)
	want := int64(3977143975)

	if got != want {
		t.Fatalf("want = %d, got = %d", want, got)
	}
}

func TestTilesForArea(t *testing.T) {
	got := TilesForArea(139.73424, 35.60433, 139.73673, 35.60707)
	want := []int64{3977143974, 3977143975}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}
