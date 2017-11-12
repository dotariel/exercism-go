package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type Result string

const (
	Win  Result = "win"
	Loss Result = "loss"
	Draw Result = "draw"
)

type Match struct {
	home   *Team
	away   *Team
	result Result
}

type Team struct {
	name    string
	matches int
	wins    int
	losses  int
	draws   int
}

func (t *Team) Points() int {
	return (3 * t.wins) + (1 * t.draws)
}

type Standings struct {
	teams []*Team
}

func (s *Standings) Team(name string) (*Team, bool) {
	for _, team := range s.teams {
		if team.name == name {
			return team, true
		}
	}

	return &Team{}, false
}

type ByPoints []*Team

func (s ByPoints) Len() int {
	return len(s)
}

func (s ByPoints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPoints) Less(i, j int) bool {
	equal := s[i].Points() == s[j].Points()

	if equal {
		return s[i].name < s[j].name
	}

	return s[j].Points() < s[i].Points()
}

func Tally(in io.Reader, out io.Writer) error {
	input, _ := ioutil.ReadAll(in)
	matches, err := parseInput(string(input))

	if err != nil {
		return err
	}

	standings := compute(matches)
	output := []byte(table(standings))

	out.Write(output)
	return nil
}

const (
	rowDelim = "\n"
	colDelim = ";"
)

func compute(matches []Match) Standings {
	standings := new(Standings)

	for _, m := range matches {
		home, ok := standings.Team(m.home.name)
		if !ok {
			home = m.home
			standings.teams = append(standings.teams, home)
		}

		away, ok := standings.Team(m.away.name)
		if !ok {
			away = m.away
			standings.teams = append(standings.teams, away)
		}

		home.matches++
		away.matches++

		switch m.result {
		case Win:
			home.wins++
			away.losses++
		case Loss:
			home.losses++
			away.wins++
		case Draw:
			home.draws++
			away.draws++
		}
	}

	return *standings
}

func parseInput(s string) ([]Match, error) {
	matches := make([]Match, 0)

	rows := strings.Split(strings.Trim(s, rowDelim), rowDelim)
	for _, row := range rows {
		match, err := parseMatch(row)

		if err != nil {
			return nil, err
		}

		if match != nil {
			matches = append(matches, *match)
		}
	}

	return matches, nil
}

func parseMatch(row string) (match *Match, err error) {
	if strings.HasPrefix(row, "#") {
		return
	}

	if len(row) == 0 {
		return
	}

	cols := strings.Split(row, colDelim)

	if len(cols) != 3 {
		err = fmt.Errorf("outcome could not be parsed; wanted length of %v but got %v", 3, len(cols))
		return
	}

	result := Result(cols[2])

	switch result {
	case Win:
	case Loss:
	case Draw:
	default:
		err = fmt.Errorf("invalid result: %s", result)
		return
	}

	match = &Match{
		home:   &Team{cols[0], 0, 0, 0, 0},
		away:   &Team{cols[1], 0, 0, 0, 0},
		result: Result(cols[2]),
	}

	return
}

func table(s Standings) string {
	sort.Sort(ByPoints(s.teams))
	str := make([]string, 0)
	str = append(str, fmt.Sprintf("%-31s| %v |  %s |  %s |  %s |  %s", "Team", "MP", "W", "D", "L", "P"))
	for _, team := range s.teams {
		str = append(str, fmt.Sprintf("%-31s|  %v |  %v |  %v |  %v |  %v", team.name, team.matches, team.wins, team.draws, team.losses, team.Points()))
	}

	return strings.Join(str, rowDelim) + rowDelim
}
