package main


import (
	"io"
	"log"
	"net"
	"sync"
//	"strings"
//	"strconv"
//	"fmt"
//	"github.com/natb0412/sem03simpleserver/conv"
	"github.com/natb0412/is105sem03/mycrypt"
)

func main() {

	var wg sync.WaitGroup

	server, err := net.Listen("tcp", "172.17.0.3:16")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bundet til %s", server.Addr().String())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Println("før server.Accept() kallet")
			conn, err := server.Accept()
			if err != nil {
				return
			}

		  	

			go func(c net.Conn) {
				defer c.Close()
				for {
					buf := make([]byte, 1024)
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return // fra for løkke
					}
					dekryptertMelding := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
					kryptertMelding := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)+5)
					log.Println("Dekrypter melding: ", string(dekryptertMelding))
					log.Println("Kryptert melding: ", string(kryptertMelding))

					//krypterMelding := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03) 5)
//                   			 if strings.HasPrefix(string(dekryptertMelding), "Kjevik") {
//                        		fields := strings.Split(string(dekryptertMelding), ";")
//                        		if len(fields) >= 4 {
//                            		celsius, err := strconv.ParseFloat(fields[3], 64)
//                            		if err != nil {
//                               		 log.Println(err)
//                                	continue
//                            		}
//                            		fahrenheit := conv.CelsiusToFahrenheit(celsius)
//                            		x = fmt.Sprintf("%s;%s;%s;%.1f", fields[0], fields[1], fields[2], fahrenheit)
//                            		if err != nil {
//                                	log.Println(err)
//                                	return // from for loop
//                            		}
//                        		} else {
//                            			log.Println("Invalid input:", string(dekryptertMelding))
//                        			}
//                    } 				else {
//                       					 x = dekryptertMelding
//							
//                    }



					switch msg := string(dekryptertMelding); msg {
  				        case "ping":
						_, err = c.Write([]byte("pong"))
					default:
						_, err = c.Write([]byte(string(kryptertMelding)))
					}
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return // fra for løkke
					}
				}
			}(conn)
		}
	}()
	wg.Wait()
}

