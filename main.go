package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	dem "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
	"github.com/sqweek/dialog"
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

func WritePromptFile(proopt *Proompt, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(proopt.text)
	if err != nil {
		return err
	}

	return nil
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

func DirectoryExists(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}

func main() {
	//filePath := "demos/rolled-16-0.dem"
	//filepath := "demos/og-vs-natus-vincere-m2-mirage.dem"

	startDir := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Counter-Strike Global Offensive\\csgo\\replays"

	if exists, err := DirectoryExists(startDir); err == nil && !exists {
		startDir = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Counter-Strike Global Offensive\\csgo"
	}

	filePath, err := dialog.File().
		Title("Select a demo file").
		Filter("Demo files (*.dem)", "dem").
		SetStartDir(startDir).
		Load()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	f, err := os.Open(filePath)

	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	p := dem.NewParser(f)
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

	err = p.ParseToEnd()
	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}

	fileName := filepath.Base(filePath) + "-proompt.txt"
	WritePromptFile(&proompt, fileName)
}
