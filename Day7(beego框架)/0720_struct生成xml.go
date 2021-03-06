package main

import (
	"os"
	"fmt"
	"encoding/xml"
)

type Servers struct{
	XMLName 	xml.Name  `xml:"servers"`
	Version		string	  `xml:"version,attr"`
	Sev			[]server  `xml:"server"`		
}

type server struct{
	ServerName	string  `xml:serverName`
	ServerIP	string  `xml:serverIP`
}

func main(){
	v := &Servers{Version:"1"}
	v.Sev = append(v.Sev,server{"shanghai","127.0.0.1"})
	v.Sev = append(v.Sev,server{"beijing","127.0.0.2"})
	output, err := xml.MarshalIndent(v,"  ","	")
	if err != nil{
		fmt.Printf("error:%v",err)
	}
	
	os.Stdout.Write([]byte(xml.Header))//输出的信息都是不带XML头的，为了生成正确的xml文件，
	//我们使用了xml包预定义的Header变量。
	os.Stdout.Write(output)
	
	
	
}


/*XMLName不会被输出

tag中含有"-"的字段不会输出

tag中含有"name,attr"，会以name作为属性名，字段值作为值输出为这个XML元素的属性，如上version字段所描述

tag中含有",attr"，会以这个struct的字段名作为属性名输出为XML元素的属性，类似上一条，只是这个name默认是字段名了。

tag中含有",chardata"，输出为xml的 character data而非element。

tag中含有",innerxml"，将会被原样输出，而不会进行常规的编码过程

tag中含有",comment"，将被当作xml注释来输出，而不会进行常规的编码过程，字段值中不能含有"--"字符串

tag中含有"omitempty",如果该字段的值为空值那么该字段就不会被输出到XML，空值包括：
false、0、nil指针或nil接口，任何长度为0的array, slice, map或者string

tag中含有"a>b>c"，那么就会循环输出三个元素a包含b，b包含c，例如如下代码就会输出

FirstName string   `xml:"name>first"`
LastName  string   `xml:"name>last"`

<name>
<first>Asta</first>
<last>Xie</last>
</name>

*/