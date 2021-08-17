package soumuradio

// SortingKey indicates sorting key
type SortingKey int

const (
	// LicensedName use name
	LicensedName SortingKey = 1
	// LicensedLocation use location
	LicensedLocation SortingKey = 2
	// StationTarget use staion target
	StationTarget SortingKey = 3
	// Date use date
	Date SortingKey = 4
)

// RegistrationSortingKey indicates sorting key
type RegistrationSortingKey int

const (
	// RegisteredName use name
	RegisteredName SortingKey = 1
	// RegisteredLocation use location
	RegisteredLocation SortingKey = 2
	// RegistarDate use registered date
	RegistarDate SortingKey = 3
	// ExpireDate use expired date
	ExpireDate SortingKey = 4
)
