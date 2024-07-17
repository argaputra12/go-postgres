package migration

var employees = []Employee{
	{
		FirstName:       "Bob",
		LastName:        "Smith",
		HireDate:        "2009-06-20",
		TerminationDate: "2016-01-01",
		Salary:          10000,
	},
	{
		FirstName:       "Joe",
		LastName:        "Jarrod",
		HireDate:        "2010-02-12",
		TerminationDate: "",
		Salary:          20000,
	},
	{
		FirstName:       "Nancy",
		LastName:        "Soley",
		HireDate:        "2012-03-14",
		TerminationDate: "",
		Salary:          30000,
	},
	{
		FirstName:       "Keith",
		LastName:        "Widjaja",
		HireDate:        "2013-09-10",
		TerminationDate: "2014-01-01",
		Salary:          20000,
	},
	{
		FirstName:       "Kelly",
		LastName:        "Smalls",
		HireDate:        "2013-09-10",
		TerminationDate: "",
		Salary:          20000,
	},
	{
		FirstName:       "Frank",
		LastName:        "Nguyen",
		HireDate:        "2015-04-10",
		TerminationDate: "2015-05-01",
		Salary:          60000,
	},
}

var annualReviews = []AnnualReviews{
	{
		ID:         10,
		EmployeeID: 1,
		ReviewDate: "2016-01-01",
	},
	{
		ID:         20,
		EmployeeID: 2,
		ReviewDate: "2016-04-12",
	},
	{
		ID:         30,
		EmployeeID: 10,
		ReviewDate: "2015-02-13",
	},
	{
		ID:         40,
		EmployeeID: 22,
		ReviewDate: "2010-10-12",
	},
	{
		ID:         50,
		EmployeeID: 11,
		ReviewDate: "2009-01-01",
	},
	{
		ID:         60,
		EmployeeID: 12,
		ReviewDate: "2009-03-03",
	},
	{
		ID:         70,
		EmployeeID: 13,
		ReviewDate: "2008-12-01",
	},
	{
		ID:         80,
		EmployeeID: 1,
		ReviewDate: "2003-04-12",
	},
	{
		ID:         90,
		EmployeeID: 1,
		ReviewDate: "2014-04-30",
	},
}
