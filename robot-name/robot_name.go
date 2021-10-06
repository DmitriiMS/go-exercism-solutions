// Package robotname provides methods for unique robot names generation
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// The robot
type Robot struct {
	robotName string
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ") //for randomization
var namePile = make(map[string]bool)               // map for holding names that were already in use
var setPrefixes = []string{}                       // empty set for alphabet prefixes: AA, AB ... ZY, ZZ
var setNumbers = []string{}                        // same for numbers from 000 to 999
// together Prefix+Number -- full name

// Prepare initializes sets for randomization
func Prepare(pref *[]string, num *[]string) {
	//generate all 676 letter combinations
	for _, i := range letters {
		for _, j := range letters {
			p := fmt.Sprintf("%s%s", string(i), string(j))
			setPrefixes = append(setPrefixes, p)
		}
	}
	// generate all 1000 numbers
	for t := 0; t < 1000; t++ {
		n := fmt.Sprintf("%03d", t)
		setNumbers = append(setNumbers, n)
	}
}

// GenerateRandomName takes random prefix and random number and returns a name
func GenerateRandomName() string {
	//Initialize sets once if it's the first time we generate the name
	if len(setPrefixes) < 1 {
		Prepare(&setPrefixes, &setNumbers)
	}
	rand.Seed(time.Now().UnixNano()) // new seed every time
	prefix := setPrefixes[rand.Intn(len(setPrefixes))]
	number := setNumbers[rand.Intn(len(setNumbers))]
	return prefix + number
}

// Name generates the robot name and checks that it's unique
func (r *Robot) Name() (string, error) {
	if r.robotName != "" { //if a robot already has a name, return it, don't do anything else
		return r.robotName, nil
	}
	name := GenerateRandomName() // generate a name
	// Now it's ttime to make sure that that name is unique and was never seen before.
	for {
		//first check that all 676000 possible combinations are not echausted
		if len(namePile) >= 676000 {
			return "", errors.New("name pool exhausted")
		}
		//if the name is already in the pile of used names, generate a new one and check again
		if _, ok := namePile[name]; ok {
			name = GenerateRandomName()
			continue
		}
		break // if the name is unique, leave
	}
	namePile[name] = true // add new name to the pile
	r.robotName = name    // set the name to robot
	return r.robotName, nil
}

//Reset resets the name to zero value, after that a new name can be generated
func (r *Robot) Reset() {
	r.robotName = ""
}
