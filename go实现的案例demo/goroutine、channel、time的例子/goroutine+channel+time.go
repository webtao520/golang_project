package main 

import (
	"time"
	"fmt"
)

func timenow(ch chan string) {
	tn := time.Now().Format("2006年01月02日 15点04分05秒.0000000 时区-0700")
    ch <-tn
}

func main (){
	ch:=make(chan string)
	for i:=0;i<10;i++ {
		go timenow(ch)
		fmt.Println(<-ch)
	    time.Sleep(500*time.Millisecond)
	}
}

/*
   ################ 运行结果 ##################
   PS D:\golang\github\golang_project\go实现的案例demo\goroutine、channel、time的例子> go run .\goroutine+channel+time.go
		2020年08月10日 10点19分27秒.2609770 时区+0800
		2020年08月10日 10点19分27秒.7786030 时区+0800
		2020年08月10日 10点19分28秒.2792588 时区+0800
		2020年08月10日 10点19分28秒.7794479 时区+0800
		2020年08月10日 10点19分29秒.2805396 时区+0800
		2020年08月10日 10点19分29秒.7809675 时区+0800
		2020年08月10日 10点19分30秒.2814328 时区+0800
		2020年08月10日 10点19分30秒.7820869 时区+0800
		2020年08月10日 10点19分31秒.2829391 时区+0800
		2020年08月10日 10点19分31秒.7834190 时区+0800
*/