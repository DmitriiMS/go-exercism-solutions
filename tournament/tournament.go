//package tournament calculates teams statistics for a football tournament
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode"
)

// statistics for a team
type statistics struct {
	teamName      string
	points        int
	matchesPlayed int
	wins          int
	draws         int
	losses        int
}

// rating that holds all statistics for all teams
type rating map[string]*statistics

func Tally(input io.Reader, output io.Writer) error {
	src, err := readStrings(input) // read input
	if err != nil {                // check for errors, if there is one, return it
		return err
	}
	rating := make(rating)
	err = rating.processTournament(src) // process source and calculate statistics
	if err != nil {                     // check for errors
		return err
	}
	err = rating.outputResults(output) // Output results to buffer
	if err != nil {                    // check for errors
		return err
	}
	return nil
}

// readStrings takes input, reads all strings from it and returns slice of strigs,
// each one of them representing one match
func readStrings(in io.Reader) ([]string, error) {
	lineScanner := bufio.NewScanner(in)
	strCut := []string{}
	for lineScanner.Scan() {
		strCut = append(strCut, lineScanner.Text())
	}
	if err := lineScanner.Err(); err != nil {
		return strCut, err
	}
	return strCut, nil
}

// processTournamen take the string slice of games played, strips comments and new lines,
// then it calls a function that validates team names and if everything is correct, calls a function that calculate statistics for teams
func (r *rating) processTournament(src []string) error {
	for _, m := range src {
		if strings.HasPrefix(m, "#") || len(m) == 0 { // ignore comments and newlines
			continue
		}
		splitMatch := strings.Split(m, ";")
		err1, err2 := isTeamNameValid(splitMatch[0]), isTeamNameValid(splitMatch[1]) // validate team names
		if err1 != nil || err2 != nil {
			return errors.New("invalid team name") // create a new error, because I don't want to check which one of the errors is not nil
		}
		err := r.calculateStats(splitMatch[0], splitMatch[1], splitMatch[2]) // calculate statistics for this game and add them to the rating
		if err != nil {
			return err
		}
	}
	return nil
}

// isTeamNameValid cheks that team name is valid.
// It's possible to just check team names against a list, but I assume that
// team naming operates on this set of rules:
// 1. team name must be composed from exactly two words;
// 2. both words must start with the same letter;
// 3. both words must be capitalized.
// This way we can easily add new teams, like "Oscillating Ocelots".
func isTeamNameValid(teamName string) error {
	words := strings.Split(teamName, " ") // split the name of the team into words
	if !(len(words) == 2) ||
		!(unicode.IsUpper(rune(words[0][0])) && unicode.IsUpper(rune(words[1][0]))) ||
		!(words[0][0] == words[1][0]) {
		return errors.New("invalid team name") // check conditions and return error if anything fails
	}
	return nil
}

// calculateStats checks whether or not statistic for a team was initialised, if not, initialises it,
// checks that match resutl is correct and calls function that does actual score calculations for each team in a match.
func (r *rating) calculateStats(firstTeam, secondTeam, result string) error {
	if _, ok := (*r)[firstTeam]; !ok {
		(*r)[firstTeam] = &statistics{teamName: firstTeam}
	}
	if _, ok := (*r)[secondTeam]; !ok {
		(*r)[secondTeam] = &statistics{teamName: secondTeam}
	}
	if result != "win" && result != "loss" && result != "draw" {
		return errors.New("invalid match result: should be win, loss or a draw")
	}
	(*r)[firstTeam].addPoints("+" + result)
	(*r)[secondTeam].addPoints("-" + result)
	return nil
}

//addPoints adds points to team statistics
func (t *statistics) addPoints(result string) {
	if result == "+win" || result == "-loss" {
		t.wins += 1
	} else if result == "-win" || result == "+loss" {
		t.losses += 1
	} else {
		t.draws += 1
	}
	t.matchesPlayed += 1
	t.points = t.wins*3 + t.draws
}

// outputResults prints sorted results to io.Writer
func (r *rating) outputResults(out io.Writer) error {
	rs := make([]statistics, 0)
	for _, value := range *r {
		rs = append(rs, *value) //for sorting we need a slice of statistics
	}
	// Reverse sort by points, if points are equal, sort by first letter of the team ascending.
	sort.Slice(rs, func(i int, j int) bool {
		switch {
		case rs[i].points > rs[j].points:
			return true
		case rs[i].points == rs[j].points:
			return rs[i].teamName[0] < rs[j].teamName[0]
		default:
			return false
		}
	})
	header := "Team                           | MP |  W |  D |  L |  P\n"
	if _, err := io.WriteString(out, header); err != nil { //print header and check for errors
		return err
	}
	for _, team := range rs { //using slice of sorted teams and points contsruct and write strings to buffer
		s := fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", team.teamName, team.matchesPlayed, team.wins, team.draws, team.losses, team.points)
		if _, err := io.WriteString(out, s); err != nil {
			return err
		}
	}
	return nil
}
