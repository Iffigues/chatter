package main

import(
	"os"
)





func writeError(err error)(vrai bool){
	vrai = false
	if err != nil{
		vrai = true
      text := err.Error()+"\n"
        f, err := os.OpenFile("./logs", os.O_APPEND|os.O_WRONLY, 0600)
           if err != nil {
	                   panic(err)
			           }

        defer f.Close()

        if _, err = f.WriteString(text); err != nil {
	                panic(err)
			        }
	}
	return
}
