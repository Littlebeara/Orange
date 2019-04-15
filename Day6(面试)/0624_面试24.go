/*Complete the code in line A, to generate content with text template 
for at most 3 tiers. The expected output is as follows:
A
B
C
D
*/
package main

import (
	"log"
	"os"
	"text/template"//template包实现了数据驱动的用于生成文本输出的模板。
)

func main() {
	const tpl = `

` // A

	type Tree struct {
		Name     string
		Children []Tree
	}
	root := Tree{
		"A", []Tree{
			Tree{"B", []Tree{
				Tree{"C", []Tree{}},
			}},
			Tree{"D", []Tree{}},
		},
	}

	t := template.Must(template.New("tree").Parse(tpl))

	err := t.Execute(os.Stdout, root)
	if err != nil {
		log.Fatalf("executing template:", err)
	}

}
