package twelve

import (
	"fmt"
	"strings"
)

var presents []string = []string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves, ",
	"three French Hens, ",
	"four Calling Birds, ",
	"five Gold Rings, ",
	"six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ",
	"eight Maids-a-Milking, ",
	"nine Ladies Dancing, ",
	"ten Lords-a-Leaping, ",
	"eleven Pipers Piping, ",
	"twelve Drummers Drumming, ",
}

var days []string = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

func Verse(d int) string {
	start := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", days[d-1])
	gifts := ""
	for i := 0; i < d; i++ {
		if i == 0 && d > 1 {
			gifts = strings.Join([]string{"and ", presents[i], gifts}, "")
			continue
		}
		gifts = strings.Join([]string{presents[i], gifts}, "")
	}
	return strings.Join([]string{start, gifts}, "")
}

func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song = strings.Trim(strings.Join([]string{song, Verse(i)}, "\n"), "\n")
	}
	return song
}
