package main
import (
    "fmt"
    "math/rand"
    "time"
    "bufio"
    "strings"
    "log"
    "os"

)


func main() {

   data_matrix := initializeMatrix()
   lookUp := createLookUp()

   for row := 0; row < 20; row++ {
   for column := 0; column < 10; column++{
       fmt.Print(data_matrix[row][column], " ")
   }
   fmt.Print("\n")
   }

   fmt.Println("input text:")
   reader := bufio.NewReader(os.Stdin)
   line, err := reader.ReadString('\n')
   if err != nil {
   log.Fatal(err)
   }

   split_input := strings.Split(line, ",")

   if contains(split_input, "OR") || contains(split_input, "AND") {
     //This is input type 1
     //Let's check if the input format is correct: Starts with a keyword(OR,AND) and keyword only appears once.
    count := 0
  for _, element := range split_input {
    if (element == "OR") || (element == "AND"){
      count++
    }
      }
    if (count != 1) || (!isOperator(split_input[0])){
       fmt.Println("Error: Invalid Input Type 1!")
         }

    len_input := len(split_input)
    operator_input := split_input[0]
    //There can be maximum of 4 combinations. Possible Options a user can enter are:
    //2, in this case len_input must be 3 because there is an operator at the beginning as well.
    //3
    //4

    if len_input == 3{
      result3 := 0
      segment1 := split_input[1]
      segment2 := split_input[2]
      segmentLetter1 := segment1[7:len(segment1)] //Extracting the segment value(A,B,C,...AA,AB,....)
      segmentLetter2 := segment2[7:len(segment2)-1]

      segmentInt1 := lookUp[segmentLetter1]
      segmentInt2 := lookUp[segmentLetter2]



      if operator_input == "AND" {
        for i:=0; i < len(data_matrix); i++ {
          if (data_matrix[i][segmentInt1] == 1) && (data_matrix[i][segmentInt2] == 1) {
            result3 ++
          }}
          fmt.Println(result3)
        }
        if operator_input == "OR" {
           for i:=0; i < len(data_matrix); i++ {
             if ((data_matrix[i][segmentInt1] == 1) || (data_matrix[i][segmentInt2] == 1)) {
               result3 ++
            }}
fmt.Println(result3)}}
    if len_input == 4{
      result4 := 0
      segment1 := split_input[1]
      segment2 := split_input[2]
      segment3 := split_input[3]
      segmentLetter1 := segment1[7:len(segment1)] //Extracting the segment value(A,B,C,...AA,AB,....)
      segmentLetter2 := segment2[7:len(segment2)]
      segmentLetter3 := segment3[7:len(segment3)-1]

      segmentInt1 := lookUp[segmentLetter1]
      segmentInt2 := lookUp[segmentLetter2]
      segmentInt3 := lookUp[segmentLetter3]
      if operator_input == "AND" {
        for i:=0; i < len(data_matrix); i++ {
          if data_matrix[i][segmentInt1] == 1 {
            if data_matrix[i][segmentInt2] == 1 {
              if data_matrix[i][segmentInt3] == 1{
                result4 = result4 + 1
              }}}}
              fmt.Println(result4)
            }
        if operator_input == "OR" {
           for i:=0; i < len(data_matrix); i++ {
             if ((data_matrix[i][segmentInt1] == 1) || (data_matrix[i][segmentInt2] == 1) || (data_matrix[i][segmentInt3] == 1)){
               result4 ++
            }
          }
          fmt.Println(result4)
          }}
    if len_input == 5{
      result5 := 0
      segment1 := split_input[1]
      segment2 := split_input[2]
      segment3 := split_input[3]
      segment4 := split_input[4]
      segmentLetter1 := segment1[7:len(segment1)] //Extracting the segment value(A,B,C,...AA,AB,....)
      segmentLetter2 := segment2[7:len(segment2)]
      segmentLetter3 := segment3[7:len(segment3)]
      segmentLetter4 := segment4[7:len(segment4)-1]

      segmentInt1 := lookUp[segmentLetter1]
      segmentInt2 := lookUp[segmentLetter2]
      segmentInt3 := lookUp[segmentLetter3]
      segmentInt4 := lookUp[segmentLetter4]

      if operator_input == "AND" {
        for i:=0; i < len(data_matrix); i++ {
          if data_matrix[i][segmentInt1] == 1 {
            if data_matrix[i][segmentInt2] == 1 {
              if data_matrix[i][segmentInt3] == 1{
                if data_matrix[i][segmentInt4] == 1{
                  result5 = result5 + 1
                }}}}}
                fmt.Println(result5)
              }

        if operator_input == "OR" {
           for i:=0; i < len(data_matrix); i++ {
             if ((data_matrix[i][segmentInt1] == 1) || (data_matrix[i][segmentInt2] == 1) || (data_matrix[i][segmentInt3] == 1) || (data_matrix[i][segmentInt4] == 1)){
               result5 ++
            }}
            fmt.Println(result5)
            }}
     }else{
      //This is input type 2
      //Let's check if the input format is correct: It should only have one string in the input eg. SegmentA
      if len(split_input) != 1 {
        fmt.Println("Error: Invalid Input Type 2!")
      }
     result := 0
     element := split_input[0]
     segmentLetter := element[7:len(element)-1] //Extracting the segment value(A,B,C,...AA,AB,....)
     var segmentInt int
     segmentInt = lookUp[segmentLetter]

     for i:=0; i < len(data_matrix); i++ {
       if data_matrix[i][segmentInt] == 1{
         result ++
       }
     }
    fmt.Println("Result:" ,result)}
  /*
   INPUT TYPES:
   1)AND/OR OPERATION FOLLOWED BY SEGMENTS UP TO 4: eg. AND, SegmentA, SegmentB
   2)One Segment Specification: eg. SegmentA

   ASSUMPTIONS:
   COLUMNS CONSECUTIVELY COME FROM THE DATABASE AS FOLLOWS:
   COLUMN1: Segment1(Corresponds to SegmentA)
   COLUMN2: Segment2(Corresponds to SegmentB)
   COLUMN3: Segment3(Corresponds to SegmentC)
        .
        .
        .
        .
   COLUMN27: Segment27(Corresponds to SegmentAA)
   COLUMN28: Segment27(Corresponds to SegmentAB)
        .
        .
        .
        .
   COLUMN676: Segment676(Corresponds to SegmentZZ)
   COLUMN677: Segment27(Corresponds to SegmentAAA)
        .
        .
        .
   COLUMN999: Segment999(Corresponds to SegmentAKK)

   -Based on the documentation, We are assuming there are only 10 columns(segments).
   -These segments are SegmentA through SegmentJ(Segment1 through Segment10)
   -Based on the input, we can directly go find the index of Segment1 in our data_matrix
   which is the 0th index. (like Segment2's index in data_matrix is 1 etc.)
   -However, this would work only for data which have 10 segments. I will impelement this. Furthermore,
   I will propose an algorithm which would handle all the 999 columns if such a large data come
   from the database.



 */


}

func isOperator(str string) bool {
  if str == "OR"{
    return true
  }
  if str == "AND" {
    return true
  }
  return false
}


func contains(s []string, str string) bool {  //It checks whether a string exists in a given slice
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}



func gRV() int { //stands for Generate Random Cell Value
    rand.Seed(time.Now().UnixNano())
    min := 0
    max := 1
    return rand.Intn(max - min + 1) + min

}



func initializeMatrix()[][]int {
  values := [][]int{}


  var user1 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()} //Creating 10 columns randomly between 0-1
  var user2 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user3 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user4 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user5 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user6 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user7 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user8 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user9 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user10 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user11 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user12 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user13 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user14 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user15 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user16 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user17 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user18 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user19 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}
  var user20 = []int{gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV(),gRV()}

  values = append(values, user1)
  values = append(values, user2)
  values = append(values, user3)
  values = append(values, user4)
  values = append(values, user5)
  values = append(values, user6)
  values = append(values, user7)
  values = append(values, user8)
  values = append(values, user9)
  values = append(values, user10)
  values = append(values, user11)
  values = append(values, user12)
  values = append(values, user13)
  values = append(values, user14)
  values = append(values, user15)
  values = append(values, user16)
  values = append(values, user17)
  values = append(values, user18)
  values = append(values, user19)
  values = append(values, user20)


return values
}

func createLookUp() map[string]int {
   m := make(map[string]int)


     m["A"] = 0
     m["B"] = 1
     m["C"] = 2
     m["D"] = 3
     m["E"] = 4
     m["F"] = 5
     m["G"] = 6
     m["H"] = 7
     m["I"] = 8
     m["J"] = 9
     m["K"] = 10
     m["L"] = 11
     m["M"] = 12
     m["N"] = 13
     m["O"] = 14
     m["P"] = 15
     m["Q"] = 16
     m["R"] = 17
     m["S"] = 18
     m["T"] = 19
     m["U"] = 20
     m["V"] = 21
     m["W"] = 22
     m["X"] = 23
     m["Y"] = 24
     m["Z"] = 25

     return m
 }
