package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (pack Package) String() string {
	return fmt.Sprintf("%v %f", pack.Name, pack.Weight)
}

func (branch Branch) String() string {
	var packges string
	for _, pack := range branch.Packages {
		packges += fmt.Sprintf("\n\t\t%v", pack.Name)
	}
	return fmt.Sprintf("\t%v:%s", branch.Id, packges)
}

func (city City) String() string {
	var branches string
	for _, branch := range city.Branches {
		branches += fmt.Sprintf("%s\n", branch)
	}
	return fmt.Sprintf("%v:\n%s", city.Name, branches)
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}

func manyToOne(s string) string {
	switch s {
	case OperationsStr:
		return OpStr

	case CitiesStr:
		return CityStr

	case BranchesStr:
		return BranchStr
	}
	return ""
}

func getLastCity() *City {
	currCity := len(countryArr) - 1
	if currCity < 0 {
		return nil
	}
	return &countryArr[currCity]
}

func getLastBranch() *Branch {
	currCity := getLastCity()
	if currCity == nil {
		return nil
	}
	index := len(currCity.Branches) - 1
	if index < 0 {
		return nil
	}
	return &currCity.Branches[index]
}

func parsePackage(words []string) error {
	if len(words) != PackageLine {
		return errors.New("Invaild Input")
	}

	if packWeightTmp, errMax := strconv.ParseFloat(words[1], 32); errMax == nil || packWeightTmp < 0 {
		branch := getLastBranch()
		packWeight := float32(packWeightTmp)

		if packWeight < branch.MinWeight || packWeight > branch.MaxWeight {
			return errors.New("Invaild Input")
		}

		packageObj := Package{words[0], packWeight}
		branch.Packages = append(branch.Packages, packageObj)
		currCity := getLastCity()
		currCity.NumOfPackages += 1
		return nil
	}
	return errors.New("Invaild Input")
}

func parseBranch(words []string) (int, error) {
	if len(words) != BranchLine {
		return 0, &Error{"Invaild Input"}
	}
	packagesNum, errNum := strconv.Atoi(words[BranchPackageNumInd])
	minWeight, errMin := strconv.ParseFloat(words[BranchMinInd], 32)
	maxWeight, errMax := strconv.ParseFloat(words[BranchMaxInd], 32)
	if errNum != nil || errMin != nil || errMax != nil || packagesNum < 0 || minWeight < 0 || maxWeight < 0 || maxWeight < minWeight {
		return 0, &Error{"Invaild Input"}
	}

	currCity := getLastCity()
	if currCity == nil {
		return 0, &Error{"Invaild Input"}
	}

	br := Branch{len(currCity.Branches), float32(minWeight), float32(maxWeight), make([]Package, 0)}
	currCity.Branches = append(currCity.Branches, br)
	return packagesNum, nil
}

func parseCity(cityName string) error {
	city := City{cityName, make([]Branch, 0), 0}
	index := len(countryArr)
	countryArr = append(countryArr, city)
	if _, ok := countryMap[cityName]; !ok { // we already have city with this name
		countryMap[cityName] = index
		return nil
	}
	return errors.New("Invaild Input")
}

func parseOp(words []string) (string, error) {
	switch words[0] {
	case OpOneStr:
		if len(words) != OpOneLine {
			return "", &Error{"Invaild Input"}
		}
		return printCity(words[1])

	case OpTwoStr:
		if len(words) != OpTwoLine {
			return "", &Error{"Invaild Input"}
		}
		return movePackages(words)

	case OpThreeStr:
		if len(words) != NumberLine {
			return "", &Error{"Invaild Input"}
		}
		return printCityWithMostPack()
	}

	return "", &Error{"Invaild Input"}
}

func printCity(cityName string) (string, error) {
	city, ok := countryMap[cityName]
	if ok {
		ctStr := fmt.Sprintf("%s", countryArr[city])
		return ctStr, nil
	}
	return "", &Error{"Cant find City"}
}

func printCityWithMostPack() (string, error) {
	maxPack := -1
	cityName := "No Town"
	for _, city := range countryArr {
		if maxPack < city.NumOfPackages {
			maxPack = city.NumOfPackages
			cityName = city.Name
		}
	}
	if maxPack > -1 {
		ctStr := fmt.Sprintf("Town with the most number of packages is %s", cityName)
		return ctStr, nil
	} else {
		return cityName, nil
	}
}

func movePackagesBetweenBranches(br1 *Branch, br2 *Branch) int {
	minWeiBr2 := br2.MinWeight
	maxWeiBr2 := br2.MaxWeight
	newElemArrayBr1 := make([]Package, 0)
	addedElem := 0
	for _, pack := range br1.Packages {
		if pack.Weight >= minWeiBr2 && pack.Weight <= maxWeiBr2 {
			br2.Packages = append(br2.Packages, pack)
			addedElem++
		} else {
			newElemArrayBr1 = append(newElemArrayBr1, pack)
		}
	}
	br1.Packages = newElemArrayBr1
	return addedElem
}

func movePackages(words []string) (string, error) {
	srcBr, srcBrFind := strconv.Atoi(words[SrcBranchInd])
	dstBr, dstBrFind := strconv.Atoi(words[DstBranchInd])
	srcCityInd, srcCtFind := countryMap[words[SrcCityInd]]
	dstCityInd, dstCtFind := countryMap[words[DstCityInd]]
	if srcBrFind != nil || dstBrFind != nil || !dstCtFind || !srcCtFind {
		return "", &Error{"Invaild Input"}
	}

	srcCity := &countryArr[srcCityInd]
	dstCity := &countryArr[dstCityInd]

	if srcBr >= len(srcCity.Branches) || dstBr >= len(dstCity.Branches) || srcBr < 0 || dstBr < 0 {
		return "", &Error{"Invaild Input"}
	}

	transfered := movePackagesBetweenBranches(&srcCity.Branches[srcBr], &dstCity.Branches[dstBr])
	dstCity.NumOfPackages += transfered
	srcCity.NumOfPackages -= transfered

	return "", nil
}

func handleLine(line, commandType string, commandSt *Stack) error {
	words := strings.Fields(line)
	var (
		err error
		str string
		num int
	)

	if len(words) < NumberLine {
		return errors.New("Invaild Input")
	}

	switch commandType {

	case OpStr:
		if str, err = parseOp(words); err == nil {
			fmt.Print(str)
			commandSt.Pop()
			return nil
		}

	case CityStr:
		if len(words) != NumberLine {
			return errors.New("Invaild Input")
		}
		if err = parseCity(words[0]); err == nil {
			commandSt.Pop()
			commandSt.Push(BranchesStr)
			return nil
		}

	case BranchStr:
		if num, err = parseBranch(words); err == nil {
			commandSt.Pop()
			for i := 0; i < num; i++ {
				commandSt.Push(PackageStr)
			}
			return nil
		}

	case PackageStr:
		if err = parsePackage(words); err == nil {
			commandSt.Pop()
			return nil
		}

	default:
		if len(words) != NumberLine {
			return errors.New("Invaild Input")
		}
		if CitiesStr == commandType || OperationsStr == commandType || BranchesStr == commandType {
			single := manyToOne(commandType)
			if num, err = strconv.Atoi(words[0]); err == nil && num >= 0 {
				commandSt.Pop()
				for i := 0; i < num; i++ {
					commandSt.Push(single)
				}
				return nil
			} else if num < 0 {
				return errors.New("Invaild Input")
			}
		}
	}

	return err
}

func main() {

	var commandSt Stack
	commandSt.Push(OperationsStr)
	commandSt.Push(CitiesStr)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Mail System")
	fmt.Println("Type exit to finish")

	for !commandSt.IsEmpty() {
		currCommand, _ := commandSt.Top()
		printLine := fmt.Sprintf("Please Enter %s", currCommand)
		fmt.Println(printLine)
		line, _ := reader.ReadString('\n')
		line = strings.Replace(line, "\r\n", "", -1)

		if line == "exit" {
			break
		}

		if line == "skip" {
			fmt.Println("skipping")
			commandSt.Pop()
			continue
		}

		err := handleLine(line, currCommand, &commandSt)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
