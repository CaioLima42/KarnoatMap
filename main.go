package main

import (	
	"fmt"
	"slices"
)

func CheckDashesAlign(min1, min2 string)bool{
	size := len(min1)
	for i := 0; i < size; i++{
		if (min1[i] == '-' && min2[i] != '-') || (min1[i] != '-' && min2[i] == '-') {
			return false
		}
	}
	return true
}

func Convert2Binary(s string)byte{
	var binary byte
	for _,v := range s{
		binary = binary << 1
		binary |= byte(v - '0')	
	}
	return binary
}

func CheckMintermDifference(min1, min2 string)bool{
	min1b := Convert2Binary(min1)
	min2b := Convert2Binary(min2)

	xorValue := min1b ^ min2b

	if xorValue & (xorValue - 1) == 0{
		return true
	} 
	return false

}

func MergeMinterms(min1, min2 string)string{
	var newString string
	for i := range min1{
		if min1[i] != min2[i]{
			newString += "-"
		}else{
			newString += string(min1[i])
		}
	}
	return newString
}


func getPrimeImplicants(min []string)[]string{
	l := len(min)
	numberOfMerge := 0
	var PrimeImplicants []string
	merges := make([]bool, l)
	for i := 0; i < l; i++{
		for j := i+1; j < l; j++{
			min1, min2 := min[i], min[j]
			if CheckDashesAlign(min1, min2) && CheckMintermDifference(min1, min2){
				mergMin := MergeMinterms(min1, min2)
				if !slices.Contains(PrimeImplicants, mergMin){
					PrimeImplicants = append(PrimeImplicants, mergMin)
				}
				numberOfMerge++
				merges[i] = true
				merges[j] = true
			}
		}
	} 
	for i := 0; i < l; i++{
		if merges[i] == false && !slices.Contains(PrimeImplicants, min[i]){
			PrimeImplicants = append(PrimeImplicants, min[i])
		}
	}
	if numberOfMerge == 0{
		return PrimeImplicants
	}
	return getPrimeImplicants(PrimeImplicants)
}

func generateCorectString(s []string)[]string{
	if len(s) == 0{
		return []string{"1"}
	}
	newSlice := make([]string, len(s))
	for _, v := range s{
		strAux := ""
		for i, v2 := range v{
			if v2 == '0'{
				strAux += string('a' + i)
				strAux += "/"
			}else if v2 == '1'{
				strAux += string('a' + i)
			}
		}
		newSlice = append(newSlice, strAux)
	}
	return newSlice
}

func main(){
	fmt.Println("Rodando")
	value := []string{"1110","1100"}
	fmt.Println(generateCorectString(getPrimeImplicants(value)))
}


