package main

import (
        "github.com/vonwenm/goserial"
        "log"
        "time"
)

func main() {

        log.Print("CONNECTING")
    
        c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600 }
            s, err := serial.OpenPort(c)
            if err != nil {
                    log.Fatal(err)
        }
      
        log.Print("CONNECTED")
             
        test := 10
        
        for (test == 10) {
        
        //======================
        

                var msg01 string  = "m1 on"
                
		n, err := s.Write([]byte( msg01 ))
		if err != nil {
			log.Fatal(err)
		}

                log.Print( "SND:  ", msg01 )

		buf := make([]byte, 128)
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

	        sx0 := string( buf[:n] )
	        
	        log.Print("REC: ", sx0 )
        
       //==================
        
        	time.Sleep(time.Second * 6 )
        
       //==================
          
                var msg02 string  = "m1 off"
	        
	        n1, err1 := s.Write([]byte( msg02  ))
	        if err1 != nil {
	                log.Fatal(err1)
	        }
	
                log.Print( "SND:  ", msg02 )
	        
	        buf1 := make([]byte, 128)
	        n1, err1 = s.Read(buf1)
	        if err1 != nil {
	                log.Fatal(err1)
	        }
	        
	        sx1 := string( buf1[:n1] )
	        
	        log.Print("REC: ", sx1 )
	        
       //==================
          
       		time.Sleep(time.Second * 6  )
        
        
        }
}
