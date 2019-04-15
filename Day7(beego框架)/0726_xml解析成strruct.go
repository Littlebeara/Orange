package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/xml"
)
//为了正确解析，go语言的xml包要求struct定义中的所有字段必须是可导出的（即首字母大写）

/*func Unmarshal(data []byte, v interface{}) error
第一个是XML数据流，第二个是存储的对应类型，目前支持struct、slice和string，
XML包内部采用了反射来进行数据的映射，所以v里面的字段必须是导出的。
有一个优先级读取流程的，首先会读取struct tag，如果没有，那么就会对应字段名
。必须注意一点的是解析的时候tag、字段名、XML元素都是大小写敏感的，所以必须一一对应字段。*/

type Recurlyservers struct {
	XMLName		xml.Name `xml:"servers"`
	//如果struct中有一个叫做XMLName，且类型为xml.Name字段，
	//那么在解析的时候就会保存这个element的名字到该字段,如上面例子中的servers。
	
	Version		string	 `xml:"version,attr"`
	//如果某个struct字段的tag定义了中含有",attr"，
	//那么解析的时候就会将该结构所对应的element的与字段同名的属性的值赋值给该字段，如上version定义。
	
	Svs			[]server `xml:"server"`
	
	Description string	 `xml:",innerxml"`
	//如果struct的一个字段是string或者[]byte类型且它的tag含有",innerxml"，
	//Unmarshal将会将此字段所对应的元素内所有内嵌的原始xml累加到此字段上
	
	
/*如果某个struct字段的tag定义 型如"a>b>c",则解析的时候，会将xml结构a下面的b下面的c元素的值赋值给该字段。

如果某个struct字段的tag定义了"-",那么不会为该字段解析匹配任何xml数据。

如果struct字段后面的tag定义了",any"，如果他的子元素在不满足其他的规则的时候就会匹配到这个字段。

如果某个XML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有",comments"的字段上，
这个字段的类型可能是[]byte或string,如果没有这样的字段存在，那么注释将会被抛弃。*/
}

type server struct{
	XMLName		xml.Name  `xml:"server"`
	ServeName   string	  `xml:"serverName`
	ServerIP    string	  `xml:serverIP`
}
func main(){
	file, err := os.Open("servers.xml")
	if err != nil{
		fmt.Printf("error : %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)//ReadAll 读取 file 中的所有数据
	if err != nil{
		fmt.Printf("error : %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil{
		fmt.Printf("error : %v", err)
		return
	}
	fmt.Println(v)
	
}












