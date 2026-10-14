package main

import "fmt"

func main() {

intNum := 5
intOther := 10
var sngNum float64 = -3

var bool1 = intNum > 5
var bool2 = intNum >= 5 && intOther < 11
var bool3 = sngNum != -1 || intOther < 0
var bool4 = !(intNum > 3) || intNum <= 5
var bool5 = !(intOther >= intNum)
var bool6 = 0 - sngNum > 0
var bool7 = 4 / 2 == intOther / intNum
var bool8 = intOther % 2 == 0
var bool9 = intOther + 2 * intNum != 30 || !(sngNum > 0)
var bool10 = intOther > 0 && intNum > 0 || sngNum > 0
var bool11 = sngNum > 0 || (intNum >= 0 && -1 * intOther == -10)
var bool12 = intNum == 5
var bool13 = intNum > 0 || (sngNum <= 0 && intOther == 13)
var bool14 = !(!(!(!(intNum > 0))))

fmt.Println(bool1)
fmt.Println(bool2)
fmt.Println(bool3)
fmt.Println(bool4)
fmt.Println(bool5)
fmt.Println(bool6)
fmt.Println(bool7)
fmt.Println(bool8)
fmt.Println(bool9)
fmt.Println(bool10)
fmt.Println(bool11)
fmt.Println(bool12)
fmt.Println(bool13)
fmt.Println(bool14)
}