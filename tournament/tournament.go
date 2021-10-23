//package tournament calculates teams statistics for football tournament
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

// Structure to hold information about the match
type match struct {
	first  string
	second string
	result string
}

// Tournament hold information about all mathces
type tournament []match

// Final rating table -- slice of statistics
type statistics struct {
	teamName      string
	points        int
	matchesPlayed int
	wins          int
	draws         int
	losses        int
}

type rating map[string]*statistics

func Tally(input io.Reader, output io.Writer) error {
	src, err := readStrings(input) // read input
	if err != nil {                // check for errors, if there is one, return it
		return err
	}
	tournament := new(tournament)
	err = tournament.processTournament(src) // process source and get tournament information
	if err != nil {                         // check for errors
		return err
	}
	rating := make(rating)
	err = rating.rateTournament(*tournament) // process the tournament and calculate statistics
	if err != nil {                          // check for errors
		return err
	}
	err = rating.outputResults(output) // Output results to buffer
	if err != nil {                    // check for errors
		return err
	}
	return nil
}

// readStrings takes input, reads all from it and returns slice of strigs,
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

// processTournamen take the string slice of games played, checks if everything in every
// string is valid and constructs a slice of matches for further processing.
// Returns an error if abything is wrong
func (t *tournament) processTournament(src []string) error {

	for _, m := range src {
		if strings.HasPrefix(m, "#") || len(m) == 0 { // ignore comments and newlines
			continue
		}
		splitMatch := strings.Split(m, ";")
		err1, err2 := isTeamNameValid(splitMatch[0]), isTeamNameValid(splitMatch[1]) // validate team names
		if err1 != nil || err2 != nil {
			return errors.New("invalid team name") // create a new error, because I don't want to check which one of the errors is not nil
		}
		if splitMatch[2] != "win" && splitMatch[2] != "loss" && splitMatch[2] != "draw" { // check if result of the game is valid
			return errors.New("invalid match result: should be win, loss or a draw")
		}
		// Create a match and append it to the tournament
		*t = append(*t, match{first: splitMatch[0], second: splitMatch[1], result: splitMatch[2]})
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

// rateTournament calculates rating of every team in the tournament.
// Various statistics are represented by map:
// data is saved in this way:
// MP            |W   |D    |L     |P
// matches played|wins|draws|losses|points
func (r *rating) rateTournament(t tournament) error {
	for _, match := range t {
		if _, ok := (*r)[match.first]; !ok {
			(*r)[match.first] = &statistics{teamName: match.first}
		}
		if _, ok := (*r)[match.second]; !ok {
			(*r)[match.second] = &statistics{teamName: match.second}
		}
		switch match.result {
		case "win":
			//statistics for the first team
			(*r)[match.first].matchesPlayed += 1
			(*r)[match.first].wins += 1
			(*r)[match.first].points += 3
			//statistics for the second team
			(*r)[match.second].matchesPlayed += 1
			(*r)[match.second].losses += 1
		case "draw":
			//statistics for the first team
			(*r)[match.first].matchesPlayed += 1
			(*r)[match.first].draws += 1
			(*r)[match.first].points += 1
			//statistics for the second team
			(*r)[match.second].matchesPlayed += 1
			(*r)[match.second].draws += 1
			(*r)[match.second].points += 1
		case "loss":
			//statistics for the first team
			(*r)[match.first].matchesPlayed += 1
			(*r)[match.first].losses += 1
			//statistics for the second team
			(*r)[match.second].matchesPlayed += 1
			(*r)[match.second].wins += 1
			(*r)[match.second].points += 3
		default:
			return errors.New("unknown error while processing tournament")
		}

	}
	return nil
}

// outputResults prints sorted results to io.Writer
func (r *rating) outputResults(out io.Writer) error {
	rs := make([]statistics, 0)
	for _, value := range *r {
		rs = append(rs, *value)
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
