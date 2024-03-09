package dates

import (
	"fmt"
	"slices"
	"time"
)

type Date string

func (d *Date) UnmarshalJSON(bytes []byte) error {

	if string(bytes) == `""` {
		return nil
	}

	dd, err := time.Parse(`"2006-01-02"`, string(bytes))
	if err != nil {
		return err
	}
	*d = Date(dd.Format("2006-01-02"))

	return nil
}

var holidays = []time.Time{
	// New Years Day
	// Jan 1st
	// 2023 NYD falls on Sunday, observed on Monday
	time.Date(2023, time.January, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC),

	// 2028 NYD falls on Satuday, observed on Friday
	time.Date(2027, time.December, 31, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.January, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.January, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2032, time.January, 1, 0, 0, 0, 0, time.UTC),

	// 2033 NYD falls on Satuday, observed on Friday
	time.Date(2032, time.December, 31, 0, 0, 0, 0, time.UTC),

	// 2034 NYD falls on Sunday, observed on Monday
	time.Date(2034, time.January, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.January, 1, 0, 0, 0, 0, time.UTC),

	// Birthday of Martin Luther King, Jr.
	// Third Monday of January
	time.Date(2023, time.January, 16, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.January, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2027, time.January, 18, 0, 0, 0, 0, time.UTC),
	time.Date(2028, time.January, 17, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.January, 15, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.January, 21, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.January, 20, 0, 0, 0, 0, time.UTC),
	time.Date(2032, time.January, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2033, time.January, 17, 0, 0, 0, 0, time.UTC),
	time.Date(2034, time.January, 16, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.January, 15, 0, 0, 0, 0, time.UTC),

	// Washington's Birthday/President's Day
	// Third Monday of February
	time.Date(2023, time.February, 20, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.February, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.February, 17, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.February, 16, 0, 0, 0, 0, time.UTC),
	time.Date(2027, time.February, 15, 0, 0, 0, 0, time.UTC),
	time.Date(2028, time.February, 21, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.February, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.February, 18, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.February, 17, 0, 0, 0, 0, time.UTC),
	time.Date(2032, time.February, 16, 0, 0, 0, 0, time.UTC),
	time.Date(2033, time.February, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2034, time.February, 13, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.February, 19, 0, 0, 0, 0, time.UTC),

	// Memorial Day
	// Last Monday of May
	time.Date(2023, time.May, 29, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.May, 27, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.May, 26, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.May, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2027, time.May, 31, 0, 0, 0, 0, time.UTC),
	time.Date(2028, time.May, 29, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.May, 28, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.May, 27, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.May, 26, 0, 0, 0, 0, time.UTC),
	time.Date(2032, time.May, 31, 0, 0, 0, 0, time.UTC),
	time.Date(2033, time.May, 30, 0, 0, 0, 0, time.UTC),
	time.Date(2034, time.May, 29, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.May, 28, 0, 0, 0, 0, time.UTC),

	// Juneteenth
	// 19th of June
	time.Date(2023, time.June, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.June, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.June, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.June, 19, 0, 0, 0, 0, time.UTC),

	// Juneteenth is a Saturday, observed Friday
	time.Date(2027, time.June, 18, 0, 0, 0, 0, time.UTC),
	time.Date(2028, time.June, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.June, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.June, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.June, 19, 0, 0, 0, 0, time.UTC),

	// Juneteenth is a Saturday, observed Friday
	time.Date(2032, time.June, 18, 0, 0, 0, 0, time.UTC),

	// Juneteenth is a Sunday, observed Monday
	time.Date(2033, time.June, 20, 0, 0, 0, 0, time.UTC),
	time.Date(2034, time.June, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.June, 19, 0, 0, 0, 0, time.UTC),

	// Independence Day
	// 4th of July
	time.Date(2023, time.July, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.July, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.July, 4, 0, 0, 0, 0, time.UTC),

	// Independence Day is a Saturday, observed Friday
	time.Date(2026, time.July, 3, 0, 0, 0, 0, time.UTC),

	// Independence Day is a Sunday, observed Monday
	time.Date(2027, time.July, 5, 0, 0, 0, 0, time.UTC),
	time.Date(2028, time.July, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.July, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.July, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.July, 4, 0, 0, 0, 0, time.UTC),

	// Independence Day is a Sunday, observed Monday
	time.Date(2032, time.July, 5, 0, 0, 0, 0, time.UTC),
	time.Date(2033, time.July, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2034, time.July, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.July, 4, 0, 0, 0, 0, time.UTC),

	// Labor Day
	// First Monday of September
	time.Date(2023, time.September, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.September, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.September, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.September, 7, 0, 0, 0, 0, time.UTC),
	time.Date(2027, time.September, 6, 0, 0, 0, 0, time.UTC),
	time.Date(2028, time.September, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.September, 3, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.September, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.September, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2032, time.September, 6, 0, 0, 0, 0, time.UTC),
	time.Date(2033, time.September, 5, 0, 0, 0, 0, time.UTC),
	time.Date(2034, time.September, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.September, 3, 0, 0, 0, 0, time.UTC),

	// Columbus/Indigenous People's Day
	// Second Monday of October
	time.Date(2023, time.October, 9, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.October, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.October, 13, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.October, 12, 0, 0, 0, 0, time.UTC),
	time.Date(2027, time.October, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2028, time.October, 9, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.October, 8, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.October, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.October, 13, 0, 0, 0, 0, time.UTC),
	time.Date(2032, time.October, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2033, time.October, 10, 0, 0, 0, 0, time.UTC),
	time.Date(2034, time.October, 9, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.October, 8, 0, 0, 0, 0, time.UTC),

	// Veteran's Day
	// 11th of November
	// Veteran's Day is a Saturday, observed on Friday
	time.Date(2023, time.November, 10, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.November, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.November, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.November, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2027, time.November, 11, 0, 0, 0, 0, time.UTC),

	// Veteran's Day is a Saturday, observed on Friday
	time.Date(2028, time.November, 10, 0, 0, 0, 0, time.UTC),

	// Veteran's Day is a Sunday, observed on Monday
	time.Date(2029, time.November, 12, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.November, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.November, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2032, time.November, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2033, time.November, 11, 0, 0, 0, 0, time.UTC),

	// Veteran's Day is a Saturday, observed on Friday
	time.Date(2034, time.November, 10, 0, 0, 0, 0, time.UTC),

	// Veteran's Day is a Sunday, observed on Monday
	time.Date(2035, time.November, 12, 0, 0, 0, 0, time.UTC),

	// Thanksgiving Day
	// 4th Thursday of November
	time.Date(2023, time.November, 23, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.November, 28, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.November, 27, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.November, 26, 0, 0, 0, 0, time.UTC),
	time.Date(2027, time.November, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2028, time.November, 23, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.November, 22, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.November, 28, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.November, 27, 0, 0, 0, 0, time.UTC),
	time.Date(2032, time.November, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2033, time.November, 24, 0, 0, 0, 0, time.UTC),
	time.Date(2034, time.November, 23, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.November, 22, 0, 0, 0, 0, time.UTC),

	// Christmas Day
	// 25th of December
	time.Date(2023, time.December, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2024, time.December, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2025, time.December, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC),

	// Christmas Day is a Saturday, observed Friday
	time.Date(2027, time.December, 24, 0, 0, 0, 0, time.UTC),
	time.Date(2028, time.December, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2029, time.December, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2030, time.December, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2031, time.December, 25, 0, 0, 0, 0, time.UTC),

	// Christmas Day is a Saturday, observed Friday
	time.Date(2032, time.December, 24, 0, 0, 0, 0, time.UTC),

	// Christmas Day is a Sunday, observed Monday
	time.Date(2033, time.December, 26, 0, 0, 0, 0, time.UTC),
	time.Date(2034, time.December, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2035, time.December, 25, 0, 0, 0, 0, time.UTC),
}

// Checks if a given date is an SEC Holiday. SEC Holidays
// mean that we get an extra day to file.
// The SEC is closed on the following holidays:
// New Year's Day,
// Birthday of Martin Luther King, Jr.,
// Washington's Birthday,
// Memorial Day,
// Juneteenth,
// Independence Day,
// Labor Day,
// Columbus Day,
// Veterans Day,
// Thanksgiving Day,
// Christmas Day.
//
// If a holiday falls on a weekend, then the SEC observes the holiday
// on the nearest weekday, either Friday or Monday.
func IsSECHoliday(timeString string) (bool, error) {

	theTime, err := time.Parse("2006-01-02", timeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}

	// Throw an error if we're out of dates. The list unfortunately needs
	// manual updating.
	if theTime.Year() > 2035 {
		return false, fmt.Errorf("out of holidays, need to update list")
	}

	// Some months never have a holiday so we can return early if the date
	// is one of those.
	if theTime.Month() == time.March || theTime.Month() == time.April || theTime.Month() == time.August {
		return true, nil
	}

	return slices.Contains(holidays, theTime), nil
}
