package soul

type Ability struct {
	Level int
	EXP   int
}

type SoulComponent struct {
	ClassId int
	RaceId  int

	EXP int // Experience Points

	Level struct {
		Initial int
		Max     int
		Current int
	}
	Abilities map[string]Ability
}

func NewSoulComponent() *SoulComponent {
	return &SoulComponent{}
}
