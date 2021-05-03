package main

import (
	"fmt"

	mycountry "github.com/GoesToTwentyOne/10DaysChallengesDrillAgain/01_First_day_challenge/02_packages/02_icomefromlasvagas_stringusefull_main/icomefromlasvagas"
	stringusefull "github.com/GoesToTwentyOne/10DaysChallengesDrillAgain/01_First_day_challenge/02_packages/02_icomefromlasvagas_stringusefull_main/stringusefull"
)

func main() {
	fmt.Println(mycountry.MyName)
	fmt.Println(stringusefull.Reverse(mycountry.MyName))
	fmt.Println(stringusefull.Reverse(stringusefull.MyName))
	fmt.Println(stringusefull.Reverse("olleH"))

}
