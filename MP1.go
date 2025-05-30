package main

import "fmt"

/*
Machine Project Specification:
Project: Weekly Payroll

Normal work week: 5 days
Normal work day: 8 hours

*employee can work overtime or on holidays (overtime pay, holiday pay)

Regular work hour: 9:00am - 6:00pm (0900-1800)
Night shift: 10:00pm - 6:00am (add 10% hourly rate/hour)
Hourly rate: daily salary/max regular work hours/day

if(work == holiday || work == rest days){
	additional pay
}

Assume work hours == 8 hours (applied to daily rate)
Rest day = 130%
Special non-working = 130%
Special non-working and Rest day = 150%
Regular holiday = 200%
Regular holiday and Rest day = 260%

Overtime (work hour > 8) (applied to hourly rate/hour)
									non-Night shift		Night shift
Normal day = 							125%				137.5%
Rest day =								169%				185.9%
Special non-working = 					169%				185.9%
Special non-working and Rest day =		195%				214.5%
Regular holiday =						260%				286.0%
Regular holiday and Rest day =			338%				371.8%

Default:
daily rate = 500
max reg hours = 8 hours
1st 5 days = work days
all 7 days = normal days
IN = 9:00am
OUT = 9:00am (can be modified by user)

Assume:
no undertime
emplayees may be absent

computation:
user will update OUT time (military format)
rest days/absent: OUT time = 0900
rest day = regular salary unless worked
if absent = no pay
salary = SUM(computed salary/day)

program:
must be menu-driven
default config can be modified
payroll generation must include salary/day & total salary for the week
*/

func computeNight_OTPay(hourly_rate float64, extra_hours float64, day_type uint8) float64 {
	var eveOTPay float64

	switch day_type {
	case 1, 2: //Rest or Special non-working
		eveOTPay = extra_hours * (hourly_rate * 1.859)
	case 3: //Special non-working & Rest
		eveOTPay = extra_hours * (hourly_rate * 2.145)
	case 4: //Regular holiday
		eveOTPay = extra_hours * (hourly_rate * 2.86)
	case 5: //Regular holiday & Rest
		eveOTPay = extra_hours * (hourly_rate * 3.718)
	default: //Normal day
		eveOTPay = extra_hours * (hourly_rate * 1.375)
	}

	return eveOTPay
}

func computeDay_OTPay(hourly_rate float64, extra_hours float64, day_type uint8) float64 {
	var dayOTPay float64

	switch day_type {
	case 1, 2: //Rest or Special non-working
		dayOTPay = extra_hours * (hourly_rate * 1.69)
	case 3: //Special non-working & Rest
		dayOTPay = extra_hours * (hourly_rate * 1.95)
	case 4: //Regular holiday
		dayOTPay = extra_hours * (hourly_rate * 2.6)
	case 5: //Regular holiday & Rest
		dayOTPay = extra_hours * (hourly_rate * 3.38)
	default: //Normal day
		dayOTPay = extra_hours * (hourly_rate * 1.25)
	}

	return dayOTPay
}

func computeAdditionalPay(day_type uint8, daily_rate float64) float64 {
	var add_pay float64

	switch day_type {
	case 1, 2: //Rest or Special non-working
		add_pay = daily_rate * 1.3
	case 3: //Special non-working & Rest
		add_pay = daily_rate * 1.5
	case 4: //Regular holiday
		add_pay = daily_rate * 2.0
	case 5: //Regular holiday & Rest
		add_pay = daily_rate * 2.6
	}

	return add_pay
}

func computeDaySalary(add_pay float64, dayOTPay float64, eveOTPay float64) float64 {
	var salary_day float64

	salary_day = add_pay + dayOTPay + eveOTPay

	return salary_day
}

func main() {
	var daily_rate float64 = 0
	var extra_hours float64 = 0

	fmt.Println("Daily Rate: ")
	fmt.Scan(&daily_rate)
	fmt.Println("Extra Hours: ")
	fmt.Scan(&extra_hours)

	var hourly_rate float64 = daily_rate / 8
	var night_rate float64 = (hourly_rate * 0.1) * extra_hours

	fmt.Println("Hourly Rate: ", hourly_rate)
	fmt.Println("Night Rate: ", night_rate)
}
