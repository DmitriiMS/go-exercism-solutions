//TODO: refactor using suggestions from mentor.
package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"unicode"
)

// Custom type just fo readability
type team string

// Structure to hold information about the match
type match struct {
	first  team
	second team
	result string
}

// Tournament hold information about all mathces
type tournament []match

// Final rating table -- map with team name as a key and anpther map for statistics
type rating map[team]map[string]int

func Tally(input io.Reader, output io.Writer) error {
	src, err := readStrings(input) // read input
	if err != nil {                // check for errors, if there is one, return it
		return err
	}
	tournament, err := processTournament(src) // process source and get tournament information
	if err != nil {                           // check for errors
		return err
	}
	rating, err := rateTournament(tournament) // process the tournament and calculate statistics
	if err != nil {                           // check for errors
		return err
	}
	err = outputResults(rating, output) // Output results to buffer
	if err != nil {                     // check for errors
		return err
	}
	return nil
}

// readStrings takes input, reads all from it and returns slice of strigs,
// each one of them representing one match
func readStrings(in io.Reader) ([]string, error) {
	str, err := ioutil.ReadAll(in) // read everything
	if err != nil {                //if tehre was an error, return it
		return []string{}, err
	}
	strCut := (strings.Split(strings.Trim(string(str), " \n"), "\n")) // trim and cut input into a slice of strings
	return strCut, nil
}

// processTournamen take the string slice of games played, checks if everything in every
// string is valid and constructs a slice of matches for further processing.
// Returns an error if abything is wrong
func processTournament(src []string) (tournament, error) {
	var currentTournament tournament
	for _, m := range src {
		if strings.HasPrefix(m, "#") || len(m) == 0 { // ignore comments and newlines
			continue
		}
		splitMatch := strings.Split(m, ";")
		err1, err2 := isTeamNameValid(splitMatch[0]), isTeamNameValid(splitMatch[1]) // validate team names
		if err1 != nil || err2 != nil {
			return []match{}, errors.New("invalid team name") // create a new error, because I don't want to check which one of the errors is not nil
		}
		if splitMatch[2] != "win" && splitMatch[2] != "loss" && splitMatch[2] != "draw" { // check if result of the game is valid
			return []match{}, errors.New("invalid match result: should be win, loss or a draw")
		}
		// Create a match and append it to the tournament
		currentTournament = append(currentTournament, match{first: team(splitMatch[0]), second: team(splitMatch[1]), result: splitMatch[2]})
	}
	return currentTournament, nil
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
	// check that it has 2 words
	// check capitalization
	// check that both words start with the same letter
	if !(len(words) == 2) ||
		!(unicode.IsUpper(rune(words[0][0])) && unicode.IsUpper(rune(words[1][0]))) ||
		!(words[0][0] == words[1][0]) {
		return errors.New("invalid team name") // if anything fails, return error
	}
	return nil
}

// rateTournament calculates rating of every team in the tournament.
// Various statistics are represented by map:
// data is saved in this way:
// MP            |W   |D    |L     |P
// matches played|wins|draws|losses|points
func rateTournament(t tournament) (rating, error) {
	var r rating = make(map[team]map[string]int)
	for _, match := range t {
		//create submaps if they are not already created
		if _, ok := r[match.first]; !ok {
			r[match.first] = make(map[string]int)
		}
		if _, ok := r[match.second]; !ok {
			r[match.second] = make(map[string]int)
		}
		switch match.result {
		case "win":
			//statistics for the first team
			r[match.first]["MP"] += 1
			r[match.first]["W"] += 1
			r[match.first]["P"] += 3
			//statistics for the second team
			r[match.second]["MP"] += 1
			r[match.second]["L"] += 1
		case "draw":
			//statistics for the first team
			r[match.first]["MP"] += 1
			r[match.first]["D"] += 1
			r[match.first]["P"] += 1
			//statistics for the second team
			r[match.second]["MP"] += 1
			r[match.second]["D"] += 1
			r[match.second]["P"] += 1
		case "loss":
			//statistics for the first team
			r[match.first]["MP"] += 1
			r[match.first]["L"] += 1
			//statistics for the second team
			r[match.second]["MP"] += 1
			r[match.second]["W"] += 1
			r[match.second]["P"] += 3
		default:
			return nil, errors.New("unknown error while processing tournament")
		}
	}
	return r, nil
}

// outputResults prints sorted results to io.Writer
func outputResults(r rating, out io.Writer) error {
	// Helper structure for sorting teams
	type sortingPair struct {
		team   team
		points int
	}
	var sortedTeams []sortingPair //create slice for sorting
	for k, v := range r {
		sortedTeams = append(sortedTeams, sortingPair{k, v["P"]}) //save team name and points
	}
	// Reverse sort by points, if points are equal, sort by first letter of the team ascending.
	sort.Slice(sortedTeams, func(i int, j int) bool {
		switch {
		case sortedTeams[i].points > sortedTeams[j].points:
			return true
		case sortedTeams[i].points == sortedTeams[j].points:
			return sortedTeams[i].team[0] < sortedTeams[j].team[0]
		default:
			return false
		}
	})
	header := "Team                           | MP |  W |  D |  L |  P\n"
	if _, err := io.WriteString(out, header); err != nil { //print header and check for eerors
		return err
	}
	for _, team := range sortedTeams { //using slice of sorted teams and points contsruct and write strings to buffer
		s := fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", string(team.team), r[team.team]["MP"], r[team.team]["W"], r[team.team]["D"], r[team.team]["L"], r[team.team]["P"])
		if _, err := io.WriteString(out, s); err != nil {
			return err
		}
	}
	return nil
}
