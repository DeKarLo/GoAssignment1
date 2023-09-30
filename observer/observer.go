package main

import (
	"fmt"
)


type Observer interface {
	Update(string)
}


type SpellCaster struct {
	observers []Observer
	spell     string
}


func (sc *SpellCaster) Attach(observer Observer) {
	sc.observers = append(sc.observers, observer)
}

func (sc *SpellCaster) CastSpell(spell string) {
	sc.spell = spell
	fmt.Printf("Casting %s spell...\n\n", spell)
	sc.NotifyObservers()
}


func (sc *SpellCaster) NotifyObservers() {
	for _, observer := range sc.observers {
		observer.Update(sc.spell)
	}
}

type WizardObserver struct {
	name string
}

func (wo *WizardObserver) Update(spell string) {
	fmt.Printf("%s received notification: Spellcaster cast %s spell.\n", wo.name, spell)
}

func main() {
	spellCaster := &SpellCaster{}
	wizard1 := &WizardObserver{name: "Wizard1"}
	wizard2 := &WizardObserver{name: "Wizard2"}

	spellCaster.Attach(wizard1)
	spellCaster.Attach(wizard2)

	spellCaster.CastSpell("🔥 Fireball 🔥 ")
	println()
	spellCaster.CastSpell("💀 Curse OF Pain 💀  ")
	println()
	spellCaster.CastSpell("❄️  Frost Bolt ❄️  ")

}