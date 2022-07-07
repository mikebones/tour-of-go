package main

// Write any import statements here

func getWrongAnswers(N int32, C string) string {
  // 1 to N questions
  // try/catch
  // check if N is greater or equal to 1
  // we know it is A or B only answers
  // in place function destructive of input
  runes := []rune(C)
  if int32(len(runes)) == N {
    for i := 0; i < len(runes); i++ {
    //   if string(runes[i]) == "A" {
    //     runes[i] = rune('B')
    //   }
    //   if string(runes[i]) == "B" {
    //     runes[i] = rune('A')
    //   }
	  switch string(runes[i]) {
		case "A":
			runes[i] = rune('B')
		case "B":	
			runes[i] = rune('A')
	  }
    }
    return string(runes)
  }
  return ""
}
