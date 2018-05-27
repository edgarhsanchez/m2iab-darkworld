package gameobjects

// ISpellCaster provides all spellcaster capabilities
type ISpellCaster interface {
	Cast(p *Player) (target *Player)
}
