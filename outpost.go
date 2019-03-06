package main

import (
	"fmt"
)

type Outpost struct {
	Population int
	Morale     int
	Stored     map[string]int
	Atmosphere int
	TechLevel  string
	Holdings   []string
	Credits    int64
	Friction   int
	TroubleTN  int
}

func NewOutpost() *Outpost {
	outpost := &Outpost{}
	outpost.Stored = make(map[string]int)
	outpost.Stored["Supplies"] = roll1dX(100, 0) * 10
	outpost.Morale = 7
	outpost.Population = roll1dX(100, 0) * 1000
	outpost.TroubleTN = 1
	outpost.Holdings = append(outpost.Holdings, "HQ")
	outpost.Holdings = append(outpost.Holdings, "Colony Housing")
	outpost.Holdings = append(outpost.Holdings, "Colony Housing")
	outpost.Holdings = append(outpost.Holdings, "Local Supply")
	outpost.Holdings = append(outpost.Holdings, "Local Supply")
	return outpost
}

func (outpost *Outpost) changePopsBy(increment int) {
	outpost.Population = outpost.Population + increment
	if outpost.Population < 0 {
		outpost.Population = 0
	}
}

func (outpost *Outpost) changeSupplesBy(increment int) {
	outpost.Stored["Supplies"] = outpost.Stored["Supplies"] + increment
	if outpost.Stored["Supplies"] < 0 {
		outpost.Stored["Supplies"] = 0
	}
}

func ColonyTrouble(index int) string {
	if index < 4 {
		return "An unknown local fungus has managed to infect and ruin half the world's Supplies."
	}
	if index < 7 {
		return "Unexpected tectonic activity damages the world's most recently-built holding. Each turn, make a Morale check for the world; the holding is useless until the check succeeds."
	}
	if index < 10 {
		return "A colonist snaps under the pressure and causes a disaster that kills 1d4 x 5% of the population."
	}
	if index < 13 {
		return "Wild economic fluctuations destroy half of the local credits the far trader holds on that world."
	}
	if index < 16 {
		return "Vicious freak weather conditions ruin a random holding and intimidate the colonists, dropping Morale by one point."
	}
	if index < 19 {
		return "Vicious native life forms are attacking the colonists, halting all holding construction or trade possibilities for 1d4 months until the beasts have been driven back."
	}
	if index < 22 {
		return "The Materials on the colony have proven badly unsuitable for local conditions. Half have been scrapped as useless."
	}
	if index < 25 {
		return "A religious or ideological zealot has accrued a following of 1d6 x 5% of the colonists, who will refuse to participate in maintaining holdings and will not count for purposes of determining minimum population. The movement will collapse in 1d6 months if not ended sooner."
	}
	if index < 28 {
		return "Critical mistakes in adapting to local agricultural conditions absolutely halt all local food production for 1d4 months. If the world doesn't have sufficient stockpiled Supplies, they'll starve."
	}
	if index < 31 {
		return "Allergic reactions to native microbial life kill 1d4 x 5% of the colonists before physicians can develop vaccinees. Morale drops by one."
	}
	if index < 34 {
		return "Criminals within the colony's trade bureaus steal half the output of the colony's factories and other production centers, halving available goods for purchase. The ring will be broken up in 1d6 months."
	}
	if index < 37 {
		return "Rebels and malcontents seize an important holding, demanding concessions that amount to at least half the goods and credits the far trader has on the world. Crushing them will destroy the holding."
	}
	if index < 40 {
		return "Sharp cultural distinctions begin to form for colonists engaged in some important role. Friction reduces the colony's Morale by one."
	}
	if index < 43 {
		return "A string of very public misfortunes leaves the colonists shaken. Colony Morale drops by one."
	}
	if index < 46 {
		return "Mistakes made while bringing up the planetary banking system leave half the colonists penniless. Colony Morale drops by one."
	}
	if index < 49 {
		return "Pirates are raiding the world. If no planetary defenses are in place, it loses one holding and all stockpiled Supplies and Materials."
	}
	if index < 52 {
		return "A heretofore-unknown local life form is extremely dangerous, and kills 1d4 x 5% of the colonists before countermeasures are developed."
	}
	if index < 55 {
		return "A terrible disease races through the population, killing 1d4 x 5% of them and lowering colony Morale by one."
	}
	if index < 58 {
		return "Colonial religious or social leaders are discovered to be manipulating supplies to their ow advantage; colony Morale drops by one point."
	}
	if index < 61 {
		return "A small breakaway fraction of colonists has turned bandit out of greed, anger, or ideology. They steal or ruin a quarter of the Supplies and Materials on the world and lower Morale by one point."
	}
	if index < 64 {
		return "Saboteurs in the employ of a rival power manage to destroy half the Supplies or half the Materials on the planet, whichever is greater."
	}
	if index < 67 {
		return "Errors at the central bank threaten to destabilize the local currency; either lose half your local credits or lower colony Morale by one."
	}
	if index < 70 {
		return "Mistakes in the local manufactories destroy half the stockpiled goods produced on the world. Supplies and Materials are unaffected."
	}
	if index < 73 {
		return "Rent-seeking and inefficiencies have become entrenched in the colony's economy; Friction rises by 2 until three consecutive months pass with no patron purchases or sales at all in the colony."
	}
	if index < 76 {
		return "A colony structure was sited on unstable ground; one Colony Housing holding is destroyed."
	}
	if index < 79 {
		return "An embittered colonist arranges the assassination of important colony officials. The world is in turmoil and can do nothing for a month."
	}
	if index < 82 {
		return "Local flora or fauna that seemed valuable and useful are actually dangerously allergenic; colony Morale drops by one."
	}
	if index < 85 {
		return "A get-rich-quick scheme among the colonists cleans most of them out- either sacrifice half your local credit balance to succor them or colony Morale drops by one."
	}
	if index < 88 {
		return "A minor problem is blown out of proportion by a demagogue who wants to use it to gain political power; colony Morale drops by one."
	}
	if index < 91 {
		return "Colonists are rebelling against rationing and harsh lives here. Immediately expend a month's worth of supplies to placate them or colony Morale drops by one."
	}
	if index < 94 {
		return "A random holding has been utterly ruined by local termite-equivalents heretofore unknown to the colonists. Future structures can be built with protection in mind."
	}
	if index < 97 {
		return "The world's factor is breaking down due to a personal tragedy. Replace them, or for the next three months, 3 in 6 chance that nothing can be built on the world that month."
	}
	if index < 100 {
		return "Land deeds are being forged by corrupt colonial administrators, causing fighting among the colonists. Pay half your local credit supply to sort things out or lose one Morale point."
	}
	return "False alarm- no disaster occurs."
}

func (outpost *Outpost) supplyDemand() int {
	return outpost.Population / 50
}

func (outpost *Outpost) step1() {
	sDemand := outpost.supplyDemand()
	fmt.Println(sDemand, " tons of supples people will consume...")
	dangeredPops := 0
	if sDemand > outpost.Stored["Supplies"] {
		dangeredPops = (sDemand - outpost.Stored["Supplies"]) * 50
		fmt.Println(dangeredPops, "people will starve...")
		toDie := dangeredPops / 10
		fmt.Println(toDie, "people will die...")
		outpost.changePopsBy(-toDie)
	}
	outpost.changeSupplesBy(-sDemand)

}

func (outpost *Outpost) step2() {
	outpost.Credits = outpost.Credits + int64(outpost.Population)
	fmt.Println(outpost.Population, "payed taxes...")

}

func (outpost *Outpost) step3() {
	fmt.Println("Buy/Build stuff")
	outpost.produce()
}

func (outpost *Outpost) step4() {
	housing := outpost.Housing()
	pops := outpost.Population
	overcrowd := pops - housing
	if overcrowd > 0 {
		toDie := (outpost.Population - outpost.Housing()) / 10
		fmt.Println(toDie, "people died because of lack of housing...")
		outpost.changePopsBy(-toDie)
	}
}

func (outpost *Outpost) step5() {
	troubleRoll := roll1dX(10, 0)
	fmt.Println("trRoll", troubleRoll, "trTN", outpost.TroubleTN)
	if troubleRoll <= outpost.TroubleTN {
		fmt.Println(ColonyTrouble(roll1dX(100, 0)))
		outpost.TroubleTN = 1
	} else {
		outpost.TroubleTN++
	}
	pplArrived := ((outpost.Population / 100) / 12)
	fmt.Println(pplArrived, "people arrived...")
	outpost.Population = outpost.Population + pplArrived
}

func (outpost *Outpost) produce() {
	for i := range outpost.Holdings {
		if outpost.Holdings[i] == "Local Supply" {
			outpost.changeSupplesBy(200)
		}
	}
}

func (outpost *Outpost) Housing() int {
	housing := 0
	for i := range outpost.Holdings {
		if outpost.Holdings[i] == "Colony Housing" {
			housing = housing + 10000
		}
	}
	return housing
}

func (outpost *Outpost) Status() {
	fmt.Println("Population:", outpost.Population)
	fmt.Println("Housing:", outpost.Housing())
	fmt.Println("Supples:", outpost.Stored["Supplies"])
	fmt.Println("Funds:", outpost.Credits)
}
