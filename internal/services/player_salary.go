package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/npm622/qo/internal/models"

	xhtml "golang.org/x/net/html"
)

// PlayerSalaryService is a PlayerSalary service
type PlayerSalaryService interface {
	FindAll(options FindAllOptions) (models.PlayerSalaries, error)
}

// FindAllOptions are the FindAll request options
type FindAllOptions struct {
	Type FindAllType
}

// FindAllType is the FindAll request type
type FindAllType string

// set of supported FindAll request types
const (
	FindAllTypeDefault FindAllType = ""
)

// Check returns an error if the FindAll request type is not valid
func (t FindAllType) Check() error {
	switch t {
	case FindAllTypeDefault:
		return nil

	}
	return fmt.Errorf("%s is not a recognized find all option type value", t)
}

func (t FindAllType) String() string {
	switch t {
	case FindAllTypeDefault:
		return "data.html"
	}
	return ""
}

// NewPlayerSalaryService creates a new PlayerSalary service
func NewPlayerSalaryService(baseURL string) PlayerSalaryService {
	return &playerSalaryService{baseURL}
}

type playerSalaryService struct {
	baseURL string
}

func (service *playerSalaryService) FindAll(options FindAllOptions) (models.PlayerSalaries, error) {
	if err := options.Type.Check(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", service.baseURL, options.Type)

	res, resErr := http.Get(url)
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()

	var rows [][]string
	var rowsErr error
	switch options.Type {
	case FindAllTypeDefault:
		rows, rowsErr = parseDefaultHTMLTable(res.Body)
	}
	if rowsErr != nil {
		return nil, rowsErr
	}
	if rows == nil {
		return nil, fmt.Errorf("unable to parse FindAll response body for type: %s", options.Type)
	}

	playerSalaries := make(models.PlayerSalaries, 0, len(rows))
	for _, row := range rows {
		playerSalary, err := models.NewPlayerSalary(row)
		if err != nil {
			return nil, err
		}
		playerSalaries = append(playerSalaries, playerSalary)
	}
	return playerSalaries, nil
}

func parseDefaultHTMLTable(r io.Reader) ([][]string, error) {
	var rows [][]string

	doc, docErr := xhtml.Parse(r)
	if docErr != nil {
		return nil, docErr
	}

	var parse func(n *xhtml.Node)
	parse = func(n *xhtml.Node) {
		if n.Type == xhtml.ElementNode && n.Data == "tr" {
			cols := map[string]string{}

			for td := n.FirstChild; td != nil; td = td.NextSibling {
				if td.Type == xhtml.ElementNode && td.Data == "td" {
					var label string
					for _, attr := range td.Attr {
						if attr.Key == "class" {
							label = attr.Val
						}
					}
					for v := td.FirstChild; v != nil; v = v.NextSibling {
						cols[label] = v.Data
					}
				}
			}
			if cols["player-name"] != "" {
				rows = append(rows, []string{
					cols["player-name"],
					cols["player-salary"],
					cols["player-year"],
					cols["player-level"],
				})
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c)
		}
	}

	parse(doc)

	return rows, nil
}
