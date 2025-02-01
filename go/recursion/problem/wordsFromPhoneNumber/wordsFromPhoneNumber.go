//Problem
   
//Words From Phone Number

//Given a seven-digit phone number, return all the character combinations that can be generated according to the following mapping:

//Graph

//Return the combinations in the lexicographical order.

//Example One

//{
//"phone_number": "1234567"
//}
//Output:

//[
//"adgjmp",
//"adgjmq",
//"adgjmr",
//"adgjms",
//"adgjnp",
//...
//"cfilns",
//"cfilop",
//"cfiloq",
//"cfilor",
//"cfilos"
//]
//First string "adgjmp" in the first line comes from the first characters mapped to digits 2, 3, 4, 5, 6 and 7 respectively. Since digit 1 maps //to nothing, nothing is appended before 'a'. Similarly, the fifth string "adgjnp" generated from first characters of 2, 3, 4, 5 second characte//r of 6 and first character of 7. All combinations generated in such a way must be returned in the lexicographical order.

//Example Two

//{
//"phone_number": "1010101"
//}
//Output:

//[""]
//Notes

//Return an array of the generated string combinations in the lexicographical order. If nothing can be generated, return a list with an empty st//ring "".
//Digits 0 and 1 map to nothing. Other digits map to either three or four different characters each.
//Constraints:

// Input string is 7 characters long; each charac//ter is a digit.
package wordsFromPhoneNumber

// get_words_from_phone_number returns all the permutations of the letters in the given phone number
func get_words_from_phone_number(phone_number string) [] string {
    numMap := make(map[byte]string, 0)
    numMap['1'] = ""
    numMap['2'] = "abc"
    numMap['3'] = "def"
    numMap['4'] = "ghi"
    numMap['5'] = "jkl"
    numMap['6'] = "mno"
    numMap['7'] = "pqrs"
    numMap['8'] = "tuv"
    numMap['9'] = "wxyz"
    numMap['0'] = ""
    results := make([]string, 0)
    helper(phone_number, 0, []byte{}, numMap, &results)
    return results
}

func helper(phoneNumber string, ndx int, slate []byte, numMap map[byte]string, results *[]string) {
    if ndx == len(phoneNumber) {
        *results = append(*results, string(slate))
				return
    }

		if len(numMap[phoneNumber[ndx]]) == 0 {
			helper(phoneNumber, ndx+1, slate, numMap, results)
			return
		}
    
    for i := 0; i < len(numMap[phoneNumber[ndx]]); i++ {
        helper(phoneNumber, ndx+1, append(slate, numMap[phoneNumber[ndx]][i]), numMap, results)
    }
}