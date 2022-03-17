package age

import (
	"fmt"
	"time"
	"regexp"
)

type Age struct {
	Years, Months int
}

func GetAge(date string) (Age, error) {
	matched, err := regexp.Match(`\d{4}-[A-Za-z]{3}-\d{2}`, []byte(date))
	if err != nil {
		return Age{}, fmt.Errorf("error: compiling regex", err)
	} else if !matched {
		return Age{}, fmt.Errorf("error: string does not match")
	}

	entered, err := time.Parse("2006-Jan-02", date)
	if err != nil {
		return Age{}, fmt.Errorf("error: parsing time", err)
	}

	now := time.Now()

	if now.Before(entered) {
		return Age{}, fmt.Errorf("error: date entered after current date")
	}

	age := Age{
		Years:  int(diff.Hours() / 8760),
		Months: int(math.Mod(diff.Hours()/730, 12)),
	}

	// fmt.Println("Years:", age.Years)
	// fmt.Println("Months:", age.Months)
	// fmt.Println("Days:", age.Days)
	return age, nil
}