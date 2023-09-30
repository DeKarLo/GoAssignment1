package main

import (
	"fmt"
)

type SpellStrategy interface{
	CastSpell()
}

type Mage struct {
	name	string
	spell	SpellStrategy
}

func (m *Mage) setSpell(spell SpellStrategy) {
	m.spell = spell
}

func (m *Mage) CastSpell() {
	fmt.Printf("%s casts a spell:", m.name)
	m.spell.CastSpell()
}

type FireSpell struct{}

func (fs *FireSpell) CastSpell() {
	fmt.Println("🔥 Fireball 🔥\n")
}

type CurseSpell struct{}

func (cs CurseSpell) CastSpell() {
	fmt.Println("💀 Curse OF Pain 💀\n")
}

type FrostSpell struct{}

func (fs FrostSpell) CastSpell() {
	fmt.Println("❄️  Frost Bolt ❄️\n")
}

func main() {
	supreme_mage := &Mage{name: "Isaac"}

	supreme_mage.setSpell(&FireSpell{})
	supreme_mage.CastSpell()

	supreme_mage.setSpell(&CurseSpell{})
	supreme_mage.CastSpell()

	supreme_mage.setSpell(&FrostSpell{})
	supreme_mage.CastSpell()
}