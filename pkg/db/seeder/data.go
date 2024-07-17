package seeder

import (
	"github.com/argaputra12/go-postgres/pkg/models"
)

var employees = []models.Employee{
	{
		FirstName:       "Bob",
		LastName:        "Smith",
		HireDate:        *stringToDate("2009-06-20"),
		TerminationDate: stringToDate("2014-01-01"),
		Salary:          10000,
	},
	{
		FirstName:       "Joe",
		LastName:        "Jarrod",
		HireDate:								*stringToDate("2010-02-12"),
		TerminationDate: nil,
		Salary:          20000,
	},
	{
		FirstName:       "Nancy",
		LastName:        "Soley",
		HireDate:        *stringToDate("2012-03-14"),	
		TerminationDate: nil,
		Salary:          30000,
	},
	{
		FirstName:       "Keith",
		LastName:        "Widjaja",
		HireDate:        *stringToDate("2013-09-10"),
		TerminationDate: stringToDate("2014-01-01"),
		Salary:          20000,
	},
	{
		FirstName:       "Kelly",
		LastName:        "Smalls",
		HireDate:        *stringToDate("2013-09-10"),
		TerminationDate: nil,
		Salary:          20000,
	},
	{
		FirstName:       "Frank",
		LastName:        "Nguyen",
		HireDate:        *stringToDate("2015-04-10"),
		TerminationDate: stringToDate("2015-05-01"),
		Salary:          60000,
	},
}


var annualReviews = []models.AnnualReviews{
	{ID: 10, EmployeeID: 1, ReviewDate: *stringToDate("2016-01-01")},
	{ID: 20, EmployeeID: 2, ReviewDate: *stringToDate("1016-04-12")},
	{ID: 30, EmployeeID: 10, ReviewDate: *stringToDate("2015-02-13")},
	{ID: 40, EmployeeID: 22, ReviewDate: *stringToDate("2010-10-12")},
	{ID: 50, EmployeeID: 11, ReviewDate: *stringToDate("2009-01-01")},
	{ID: 60, EmployeeID: 12, ReviewDate: *stringToDate("2009-03-03")},
	{ID: 70, EmployeeID: 13, ReviewDate: *stringToDate("2008-12-01")},
	{ID: 80, EmployeeID: 1, ReviewDate: *stringToDate("2003-04-12")},
	{ID: 90, EmployeeID: 1, ReviewDate: *stringToDate("2014-04-30")},
}

	