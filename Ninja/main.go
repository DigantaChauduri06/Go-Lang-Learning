package main

import "fmt"

func main() {
	// fmt.Print("Hoa")
	/*
	
	// Data type
	var name string = "Diganta"
	var age = 100;
	// var isVoted bool
	// isVoted = false
	// nationalty:="Indian"

	// fmt.Print(name, age, isVoted, nationalty)

	// Format specifier
	fmt.Printf("my name is %v and my age is %v \n",name, age)
	fmt.Printf("my name is type of %T and my age is type of %T \n",name, age)
	fmt.Printf("my name is string formatted %q \n",name)
	fmt.Printf("To print float we use '%0.3f' \n",12.33)

	// To save the formatted string we use Sprintf
	var str = fmt.Sprintf("my name is %v and my age is %v \n",name, age)
	fmt.Print(str)
	*/
/*

	//? Arrays

	var ages [4]int = [4]int {10,12,45,64}
	var ages2 = [4]int {1000,22,66,44}
	ages3 := [4]int {50,20,30,40}
	fmt.Println(ages,ages2, ages3, len(ages))


	//? Slices (uses array under the hood) (Dynamic array)
	var nums = []int {100,200,300,400}
	nums = append(nums, 500)
	fmt.Println(nums, len(nums))

	/*
		* Slicing is Done same as python using [:] 
		* Which has 2 params [start:end-1]
	
	range1 := nums[1:2]
	range2 := nums[2:]
	fmt.Println(range1, range2)

	*/

	//? Strings
	/*
		* Important Methods of Strings

		1. strings.Contains(str, f_str)
		2. strings.Split(str, split_str)
		3. strings.ToUpper(str)
		4. strings.ToLower(str)
		5. strings.Index(str, f_str)
		6. strings.ReplaceAll(str, f_str)
		6. strings.Replace(str, f_str)

		greeting := "Hello There"
	
		fmt.Println(strings.Contains(greeting, "The"))
	*/
	/*

	// Sort fn only sorts Slices(Dynamic Arrays)
	ages3 := []int {50,20,30,40}
	ages4 := []int {50,20,30,40}
	sort.Ints(ages3)
	fmt.Println(ages3)
	// sort.SearchInts(arr,num) Doesn't change the order only finds the matching value if it finds then retuns the index other wise it will return len(nums) + 1

	fmt.Println(sort.SearchInts(ages4,100)) 
	fmt.Println(ages4)

	*/
	// LOOPS

	i := 0
	for i < 20 {
		fmt.Print(i, " ")
		i++;
	}
	fmt.Println()
	for i:=0; i < 20;i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	names := []string {"Diganta", "Ankur", "Apsara"}

	for _,value := range names {
		fmt.Printf("Name = %v \n", value)
	}
	
}