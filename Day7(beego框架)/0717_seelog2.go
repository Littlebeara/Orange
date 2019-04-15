package main

import (
	"fmt"
	seelog "github.com/cihub/seelog"
)

var logger seelog.LoggerInterface

func loadAppConfig(){
	appConfig := `
<seelog minlevel="warn">
    <outputs formatid="common">
        <rollingfile type="size" filename="/data/logs/roll.log" maxsize="100000" maxrolls="5"/>
        <filter levels="critical">
            <file path="/data/logs/critical.log" formatid="critical"/>
            <smtp formatid="criticalemail" senderaddress="astaxie@gmail.com" sendername="ShortUrl API" hostname="smtp.gmail.com" hostport="587" username="mailusername" password="mailpassword">
                <recipient address="xiemengjun@gmail.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date/%Time [%LEV] %Msg%n" />
        <format id="critical" format="%File %FullPath %Func %Msg%n" />
        <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil{
		fmt.Println(err)
		return
	}
	
	UserLogger(logger)
	
	
}

func init(){
	Disablelog()
	loadAppConfig()
}

func Disablelog(){
	Logger = seelog.Disabled
}

func Userlogger(newlogger seelog.LoggerInterface){
	Logger = newlogger
}











