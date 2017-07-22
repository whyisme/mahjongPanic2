package game

type TileType int
const (
  Dot TileType = iota
  Bamboo
  Character
  Wind
  Dragon
)
const (
  East int = iota
  South
  West
  North
)
const (
  Red int = iota
  Green
  White
)
type Tile struct {
  Type TileType
  Value int
}