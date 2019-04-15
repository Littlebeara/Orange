/*Fix the regex pattern in line A, to replace all consecutive 0s to one 0 when 0 is the beginning of the line.
The expected output is like:
100
0
1
0
1
*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := `100
00
0
1
0

0
1`
	pattern := regexp.MustCompile("0(0|\n)*0") //A   |:两项中间选一个匹配   \n匹配换行符
	//* :  匹配前面的子表达式零次或多次   
	s = pattern.ReplaceAllString(s, "0")
	fmt.Println(s)
	fmt.Fscanf()
}

                        