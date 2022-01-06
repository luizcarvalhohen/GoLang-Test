package main

import (
    "fmt"
    "strings"
    "strconv"
    "os"
)

/* ---------------------------------------------------------- utils --------------------------------------------------*/
/*
function to chunk a slice
return chunk slice
*/
func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

/*
function to get the type of a string
return boolean
*/
func testType(stringIn string) bool{
    _, err := strconv.Atoi(stringIn)
    if err != nil {
        return false
    } else {
        return true
    }
}

/*
function to check if a string is empty
return boolean
*/
func isEmpty(stringIn string) bool{
if len(stringIn) != 0 {
    return false
}else{
    return true
}

}

/*
function to check if a slice of string contains a specific string
return boolean
*/
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

/*
function to check the longest length of a string in slice
return string
*/
func longestWord(s string) string {
    best, length := "", 0
    for _, word := range strings.Split(s, " ") {
        if len(word) > length {
            best, length = word, len(word)
        }
    }
    return best
}

/*
function to check the shortest length of a string in slice
return string
*/
func shortestWord(s string) string {
    best, length := "", 1000000000000
    for _, word := range strings.Split(s, " ") {
        if len(word) < length {
            best, length = word, len(word)
        }
    }
    return best
}

/* ------------------------------------------------------ end utils --------------------------------------------------*/
// ---------------------------------------------------------------------------------------------------------------------
/*
function that takes a string as an input, and returns boolean flag `true` if the given string complies with the format {num-str}
or `false` if the string does not comply
*/
func testValidity(inputStr string) bool{
    booleanList := []string{}
    inString := strings.Split(inputStr, "-")
    slicedString := chunkSlice(inString,2)

    for i := range slicedString {
        strInt := slicedString[i][0]
        strStr := slicedString[i][1]
        if testType(strInt) && ! testType(strStr)  && ! isEmpty(strInt) && ! isEmpty(strStr){
            booleanList = append(booleanList, "true")
        }else{
            booleanList = append(booleanList, "false")
        }
}

return ! contains(booleanList,"false")

}
// ---------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------
/*
function that takes the string, and returns the average number from all the numbers
*/
func averageNumber(inputStr string) int64{
    intList := []int64{}
    inString := strings.Split(inputStr, "-")
    slicedString := chunkSlice(inString,2)

    for i := range slicedString {
        strInt := slicedString[i][0]
        if testType(strInt) && ! isEmpty(strInt){
        strToInt,_ := strconv.ParseInt(strInt,0, 64)
            intList = append(intList, strToInt)
        }
}
    // size of the array
    n := len(intList)

    // declaring a variable
    // to store the sum
    sum := int64(0)

    // traversing through the
    // array using for loop
    for i := 0; i < n; i++ {

        // adding the values of
        // array to the variable sum
        sum += (intList[i])
    }

    // declaring a variable
    // avg to find the average
    avg := (int64(sum)) / (int64(n))

return avg

}
// ---------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------
/*
function that takes the string, and returns a text that is composed from all the text words separated by spaces,
 e.g. `story` called for the string `1-hello-2-world` should return text: `"hello world"
*/
func wholeStory(inputStr string) string {
    strList := []string{}
    inString := strings.Split(inputStr, "-")
    slicedString := chunkSlice(inString,2)

    for i := range slicedString {
        strStr := slicedString[i][1]
        if ! testType(strStr) && ! isEmpty(strStr){
            strList = append(strList, strStr)
        }
}

return strings.Join(strList[:], " ")

}
// ---------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------
/*
function that returns four things:
   * the shortest word
   * the longest word
   * the average word length
   * the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
*/
func storyStats(inputStr string) (string,string,int64,[]string){
    theWholeStory := wholeStory(inputStr)
    intList := []int64{}
    sameAsTheAverage := []string{}
    inString := strings.Split(theWholeStory, " ")
    // the shortest word
    shortestWordBest := shortestWord(theWholeStory)
    // the longest word
    longestWordBest := longestWord(theWholeStory)

    // the average word length
    for i := range inString {
        strStr := inString[i]
        if ! testType(strStr) && ! isEmpty(strStr){
            intList = append(intList, int64(len(strStr)))
        }
}
    // size of the array
    n := len(intList)

    // declaring a variable
    // to store the sum
    sum := int64(0)

    // traversing through the
    // array using for loop
    for i := 0; i < n; i++ {

        // adding the values of
        // array to the variable sum
        sum += (intList[i])
    }

    // declaring a variable
    // avg to find the average
avg := (int64(sum)) / (int64(n))

// the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
for i := range inString {
        strStr := inString[i]
        if ! testType(strStr) && ! isEmpty(strStr){
        if int64(len(strStr)) == avg {
        sameAsTheAverage = append(sameAsTheAverage, strStr)
        }
        }
}

return shortestWordBest,longestWordBest,avg,sameAsTheAverage
}
// ---------------------------------------------------------------------------------------------------------------------


func main() {
    if len(os.Args) > 1 {
    inputStr := os.Args[1]
    if testValidity(inputStr) {
    // testValidity
    fmt.Println("test validity :", testValidity(inputStr))
    // averageNumber
    fmt.Println("average number :", averageNumber(inputStr))
    // wholeStory
    fmt.Println("whole story :", wholeStory(inputStr))
    // storyStats
    shortestWordBest,longestWordBest,avg,sameAsTheAverage := storyStats(inputStr)
    fmt.Println("history stats :\n\t shortest word:", shortestWordBest, "\n\t longest word:" ,longestWordBest, "\n\t average word:", avg, "\n\t same as average:" , sameAsTheAverage)
    }else{
    fmt.Println("invalid sequence ->", inputStr, "\n\t eg: <sequence_number_0>-<sequence_string_0>-<sequence_number_1>-<sequence_string_1-<sequence_number_N>-<sequence_string_N>")
    }
    } else {
    fmt.Println("usage : go run logic.go <string to work with>")
}
}