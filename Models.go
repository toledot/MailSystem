package main

type Error struct {
	Msg string
}

type Package struct {
	Name   string
	Weight float32
}

type Branch struct {
	Id        int
	MinWeight float32
	MaxWeight float32
	Packages  []Package
}

type City struct {
	Name          string
	Branches      []Branch
	NumOfPackages int
}

const (
	OperationsStr = "Number Of Operations"
	OpStr         = "Operation"
	CitiesStr     = "Number Of Cities"
	CityStr       = "City"
	BranchesStr   = "Number Of Branches"
	BranchStr     = "Branch"
	PackageStr    = "Package"

	OpTwoLine   = 5
	BranchLine  = 3
	PackageLine = 2
	OpOneLine   = 2
	NumberLine  = 1

	OpOneStr   = "1"
	OpTwoStr   = "2"
	OpThreeStr = "3"

	BranchPackageNumInd = 0
	BranchMinInd        = 1
	BranchMaxInd        = 2

	SrcCityInd   = 1
	SrcBranchInd = 2
	DstCityInd   = 3
	DstBranchInd = 4
)

var (
	countryArr = make([]City, 0)
	countryMap = make(map[string]int) // city name to index in countryArr
)
