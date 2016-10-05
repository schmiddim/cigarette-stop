package main

import (
	"time"
	"fmt"
	"os"
	"strconv"
)


/*
example  ./smoke-stop  2016-09-19 30 0.285714286
 */

import (
	"os"
	"fmt"
	"time"
	"strconv"
)

func getSmokingString(dateStopped time.Time, cigarettesPerDay int64, pricePerCigarette float64) string {
	today := time.Now()
	diff := today.Sub(dateStopped)
	hours := diff.Hours()

	cigarettesNotSmoked := float64(cigarettesPerDay) / float64(24) * float64(hours)
	moneySaved := cigarettesNotSmoked * pricePerCigarette
	daysWithoutSmoking := hours / 24

	return fmt.Sprintf("You stopped smoking %d days ago. Saved %d Euro and you didn't smoked %d cigarettes",

		int(daysWithoutSmoking),
		int(moneySaved),
		int(cigarettesNotSmoked))
}

func main() {
	dateStopped, e := time.Parse(time.RFC3339, os.Args[1] + "T00:00:00+00:00")
	if (nil != e) {
		fmt.Println(e)
	}

	cigarettesPerDay, e := strconv.ParseInt(os.Args[2], 10, 64)
	if (nil != e) {
		fmt.Println(e)
	}

	pricePerCigarette, e := strconv.ParseFloat(os.Args[3], 64)
	if (nil != e) {
		fmt.Println(e)
	}

	fmt.Println(getSmokingString(dateStopped, cigarettesPerDay, pricePerCigarette))
}
