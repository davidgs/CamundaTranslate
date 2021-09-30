package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	datesi18n "github.com/davidgs/datesi18n"
)

//var XlatedDates = DateLocales{}
func (d *DateLocales) TranslateDateLocales(lang string) {
	dts := datesi18n.NewDatesI18N(lang)
	d.LongDateFormat = dls.LongDateFormat
	d.Week = dls.Week
	mp := make(map[string]interface{})
	trs := make(map[string]interface{})
	reqBodyBytes := new(bytes.Buffer)
	// Calendar first
	json.NewEncoder(reqBodyBytes).Encode(dls.Calendar)
	err := json.Unmarshal(reqBodyBytes.Bytes(), &mp)
	if err != nil {
		fmt.Println("Calendar Unmarshall error: ", err)
	}
	trs = d.translateMap(lang, mp)
	//XlatedDates.Calendar = Calendar{}
	json.NewEncoder(reqBodyBytes).Encode(trs)
	bb, err := json.Marshal(trs)
	err = json.Unmarshal(bb, &d.Calendar)
	if err != nil {
		fmt.Println("Calendar Marshal error: ", err)
	}
	// Months
	for k := range dls.Months {
		d.Months[k] = dts.MonthName(k)
		d.MonthsShort[k] = dts.ShortMonthName(k)
	}
	// Days
	for k := range dls.Weekdays {
		d.Weekdays[k] = dts.DayName(k)
		d.WeekdaysShort[k] = dts.ShortDayName(k)
		d.WeekdaysMin[k] = dts.MinDayName(k)
	}
	// Relative Time
	rtBodyBytes := new(bytes.Buffer)
	json.NewEncoder(rtBodyBytes).Encode(dls.RelativeTime)
	mt := make(map[string]interface{})
	rt := make(map[string]interface{})
	err = json.Unmarshal(rtBodyBytes.Bytes(), &mt)
	if err != nil {
		fmt.Println("Relative Time Unmarshall error: ", err)
	}
	rt = d.translateMap(lang, mt)
	json.NewEncoder(reqBodyBytes).Encode(rt)
	ff, err := json.Marshal(rt)
	err = json.Unmarshal(ff, &d.RelativeTime)
	if err != nil {
		fmt.Println("Marshal error: ", err)
	}
}


func (d *DateLocales) translateMap(lang string, mp map[string]interface{}) map[string]interface{} {
	trs := make(map[string]interface{})
	trans := Translator{}
	for k, v := range mp {
		s, err := trans.TranslateTextWithModel(lang, fmt.Sprintf("%v", v), "nmt")
		if err != nil {
			fmt.Println("Translate Error: ", err)
			return nil
		}
		trs[k] = s
	}
	return trs
}


var dls = DateLocales{
	Months:        [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
	MonthsShort:   [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
	Weekdays:      [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	WeekdaysShort: [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
	WeekdaysMin:   [7]string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
	LongDateFormat: LongDate {
		LT: "HH:mm",
		L:    "DD/MM/YYYY",
		LL:   "D MMMM YYYY",
		LLL:  "D MMMM YYYY LT",
		LLLL: "dddd, D MMMM YYYY LT",
	},
	Calendar: Calendar {
		SameDay:  "[Today at] LT",
		NextDay:  "[Tomorrow at] LT",
		NextWeek: "dddd [at] LT",
		LastDay:  "[Yesterday at] LT",
		LastWeek: "[Last] dddd [at] LT",
		SameElse: "L",
	},
	RelativeTime: RelativeTime {
		Future: "in %s",
		Past:   "%s ago",
		S:      "a few seconds",
		M:      "a minute",
		MM:     "%d minutes",
		H:      "an hour",
		HH:     "%d hours",
		D:      "a day",
		DD:     "%d days",
		M_:     "a month",
		M_M:    "%d months",
		Y:      "a year",
		YY:     "%d years",
	},
	Week: Week{
      Dow: 1,
      Doy: 4,
	},
}

type Calendar struct {
	SameDay  string `json:"sameDay"`
	NextDay  string `json:"nextDay"`
	NextWeek string `json:"nextWeek"`
	LastDay  string `json:"lastDay"`
	LastWeek string `json:"lastWeek"`
	SameElse string `json:"sameElse"`
}

type LongDate struct {
	LT   string `json:"LT"`
	L    string `json:"L"`
	LL   string `json:"LL"`
	LLL  string `json:"LLL"`
	LLLL string `json:"LLLL"`
}

type DateLocales struct {
	Months         [12]string   `json:"months"`
	MonthsShort    [12]string   `json:"monthsShort"`
	Weekdays       [7]string    `json:"weekdays"`
	WeekdaysShort  [7]string    `json:"weekdaysShort"`
	WeekdaysMin    [7]string    `json:"weekdaysMin"`
	LongDateFormat LongDate     `json:"longDateFormat"`
	Calendar       Calendar     `json:"calendar"`
	RelativeTime   RelativeTime `json:"relativeTime"`
	Week           Week         `json:"week"`
}

type Week struct {
	Dow int `json:"dow"`
	Doy int `json:"doy"`
}

type RelativeTime struct {
	Future string `json:"future"`
	Past   string `json:"past"`
	S      string `json:"s"`
	M      string `json:"m"`
	MM     string `json:"mm"`
	H      string `json:"h"`
	HH     string `json:"hh"`
	D      string `json:"d"`
	DD     string `json:"dd"`
	M_     string `json:"M"`
	M_M    string `json:"MM"`
	Y      string `json:"y"`
	YY     string `json:"yy"`
}
