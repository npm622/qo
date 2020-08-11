package commands

import (
	"sort"

	"github.com/npm622/qo/internal/models"
	"github.com/npm622/qo/internal/services"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var defaultPrinter = message.NewPrinter(language.Make("en"))

type analytics struct {
	size            int
	qualifyingOffer float64
	netAvg          float64
	grossAvg        float64
	noSalary        []string
	zeroSalary      []string
}

func analyzeSalaries(rs models.PlayerSalaries, verbose, ferbose bool) analytics {
	var a analytics

	if !verbose && !ferbose {
		for _, ps := range rs[0:125] {
			a.qualifyingOffer += *ps.Salary
		}
	} else {
		var total float64
		for i, ps := range rs {
			var salary float64

			if ps.Salary == nil {
				a.noSalary = append(a.noSalary, ps.Player.String())
			} else {
				salary = *ps.Salary
				if salary == 0 {
					a.zeroSalary = append(a.zeroSalary, ps.Player.String())
				}
			}

			if i < 125 {
				a.qualifyingOffer += salary
			}
			total += salary
		}
		a.grossAvg = total / float64(len(rs))
		a.netAvg = total / float64(len(rs)-(len(a.noSalary)+len(a.zeroSalary)))
	}

	a.size = len(rs)
	a.qualifyingOffer = a.qualifyingOffer / 125

	return a
}

func getSortedPlayerSalaries(baseURL string) (models.PlayerSalaries, error) {
	service := services.NewPlayerSalaryService(baseURL)

	res, err := service.FindAll(services.FindAllOptions{})
	if err != nil {
		return nil, err
	}

	sort.Slice(res, func(i, j int) bool {
		var s1, s2 float64
		ps1, ps2 := res[i], res[j]

		if s := ps1.Salary; s != nil {
			s1 = *s
		}
		if s := ps2.Salary; s != nil {
			s2 = *s
		}
		if s1 == s2 {
			return ps1.Player.Last < ps2.Player.Last
		}
		return s1 > s2
	})

	return res, nil
}
