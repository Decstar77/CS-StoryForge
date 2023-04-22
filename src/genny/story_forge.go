package main

import (
	"fmt"
	"io"
	"log"

	dem "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
)

type Proompt struct {
	text string
}

func AddProomptData(proopt *Proompt, s string) {
	proopt.text += s + "\n"
}

func ProomptMatchStartData(proopt *Proompt, gs dem.GameState) {
	tTeam := gs.TeamTerrorists()
	ctTeam := gs.TeamCounterTerrorists()

	tName := tTeam.ClanName()
	ctName := ctTeam.ClanName()

	if len(tName) == 0 {
		tName = "Terrorists"
	}

	if len(ctName) == 0 {
		ctName = "Counter Terrorists"
	}

	AddProomptData(proopt, fmt.Sprintf("%s vs %s", tName, ctName))

	tPlayerProompt := tName + ": "
	tPlayers := tTeam.Members()
	for _, p := range tPlayers {
		tPlayerProompt += p.Name + ", "
	}

	AddProomptData(proopt, tPlayerProompt)

	ctPlayerProompt := ctName + ": "
	ctPlayers := ctTeam.Members()
	for _, p := range ctPlayers {
		ctPlayerProompt += p.Name + ", "
	}

	AddProomptData(proopt, ctPlayerProompt)
}

func ProomptPlayerKillData(proopt *Proompt, e events.Kill) {
	hs := ""
	if e.IsHeadshot {
		hs = "(HeadShot)"
	}

	AddProomptData(proopt, fmt.Sprintf("%s killed %s with %s%s", e.Killer.Name, e.Victim.Name, e.Weapon.String(), hs))
}

func ProomptPlayerStats(proopt *Proompt, player *common.Player) {
	kills := player.Kills()
	deaths := player.Deaths()
	assists := player.Assists()
	totalDamage := player.TotalDamage()
	monies := player.MoneySpentTotal()

	AddProomptData(proopt, fmt.Sprintf("%s Kills:%d, Deaths:%d, Assists:%d, TotalDamage:%d, TotalMoneySpent:%d", player.Name, kills, deaths, assists, totalDamage, monies))
}

func IsKnifeRound(gs dem.GameState) bool {
	for _, p := range gs.TeamCounterTerrorists().Members() {
		w := p.Weapons()
		if len(w) == 1 && w[0].Type == common.EqKnife {
			return true
		}
	}

	return false
}

func GenerateProompt(readStream io.Reader) Proompt {
	p := dem.NewParser(readStream)
	defer p.Close()

	proompt := Proompt{}

	AddProomptData(&proompt, "Write a flamboyant narrative story about the following csgo match:")

	roundNum := 1
	p.RegisterEventHandler(func(e events.RoundStart) {
		if p.GameState().IsWarmupPeriod() == false && IsKnifeRound(p.GameState()) == false {

			if roundNum == 1 {
				ProomptMatchStartData(&proompt, p.GameState())
			}

			AddProomptData(&proompt, fmt.Sprintf("Round %d", roundNum))
			roundNum += 1
		}
	})

	p.RegisterEventHandler(func(e events.Kill) {
		if p.GameState().IsWarmupPeriod() == false && IsKnifeRound(p.GameState()) == false && roundNum > 1 {
			ProomptPlayerKillData(&proompt, e)
		}
	})

	p.RegisterEventHandler(func(e events.RoundEnd) {
		if p.GameState().IsWarmupPeriod() == false && IsKnifeRound(p.GameState()) == false && roundNum > 1 {
			AddProomptData(&proompt, fmt.Sprintf("Round over: %s", e.Message))
		}
	})

	p.RegisterEventHandler(func(e events.AnnouncementWinPanelMatch) {
		AddProomptData(&proompt, "======Player stats======")
		for _, pl := range p.GameState().Participants().Playing() {
			ProomptPlayerStats(&proompt, pl)
		}

		AddProomptData(&proompt, "======Game over======")
	})

	err := p.ParseToEnd()
	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}

	return proompt
}
