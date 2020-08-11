package models

import (
	"fmt"
	"strconv"
	"strings"
)

type PlayerSalaries []PlayerSalary

type PlayerSalary struct {
	Player Player
	Salary *float64
	Year   int
	Level  PlayerLevel
}

func NewPlayerSalary(data []string) (PlayerSalary, error) {
	player, playerErr := NewPlayer(data[0])
	if playerErr != nil {
		return PlayerSalary{}, playerErr
	}

	salary, salaryErr := parseSalary(data[1])
	if salaryErr != nil {
		return PlayerSalary{}, salaryErr
	}

	year, yearErr := strconv.Atoi(data[2])
	if yearErr != nil {
		return PlayerSalary{}, yearErr
	}

	level, levelErr := NewPlayerLevel(data[3])
	if levelErr != nil {
		return PlayerSalary{}, levelErr
	}

	return PlayerSalary{player, salary, year, level}, nil
}

func parseSalary(rawSalary string) (*float64, error) {
	if rawSalary == "no salary data" {
		return nil, nil
	}

	var salary float64

	if rawSalary != "" {
		if rawSalary[0] == '$' {
			rawSalary = rawSalary[1:]
		}

		sal, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(rawSalary, "$", ""), ",", ""), 10)
		if err != nil {
			return nil, err
		}
		salary = sal
	}
	return &salary, nil
}

func (ps PlayerSalary) String() string {
	salary := "no salary data"
	if ps.Salary != nil {
		salary = fmt.Sprintf("$%.2f", *ps.Salary)
	}
	return fmt.Sprintf("[%d %s] %s: %s", ps.Year, ps.Level, ps.Player, salary)
}

type Player struct {
	Raw   string
	First string
	Last  string
}

func NewPlayer(raw string) (Player, error) {
	parts := strings.Split(raw, ",")
	if len(parts) > 2 {
		return Player{}, fmt.Errorf("unique name: %s", strings.Join(parts, " "))
	}

	return Player{
		Raw:   raw,
		First: strings.TrimSpace(parts[1]),
		Last:  strings.TrimSpace(parts[0]),
	}, nil
}

func (p Player) String() string {
	return fmt.Sprintf("%s %s", p.First, p.Last)
}

type PlayerLevel string

const (
	PlayerLevelMLB PlayerLevel = "MLB"
)

func NewPlayerLevel(raw string) (PlayerLevel, error) {
	level := PlayerLevel(raw)
	switch level {
	case PlayerLevelMLB:
		return level, nil
	}
	return level, fmt.Errorf("%s is not a valid player level", raw)
}
