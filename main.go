package main

import (
	"fmt"
	"time"
)

func Proceso(id uint64, cT chan string, cC chan bool) {
	var i uint64 = 0
	for {
		select {
		case <-cC:
			return
		default:
			cT <- fmt.Sprint(id) + ": " + fmt.Sprint(i)
			i += 1
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func Printer(cT chan string, cC chan bool) {
	p := false
	for {
		select {
		case <-cC:
			p = !p
		case msg := <-cT:
			if p {
				fmt.Println(msg)
			}
		}
	}
}

func main() {
	var opc int
	var nPcs uint64 = 0
	c1 := make(chan string)
	c2 := make(chan bool)
	cPs := make(map[uint64]chan bool)
	go Printer(c1, c2)
	for opc != 4 {
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar procesos")
		fmt.Println("3) Eliminar proceso")
		fmt.Println("4) Salir")
		fmt.Scanln(&opc)
		switch opc {
		case 1:
			nC := make(chan bool)
			cPs[nPcs] = nC
			go Proceso(nPcs, c1, cPs[nPcs])
			nPcs += 1
		case 2:
			c2 <- true
		case 3:
			var dlt uint64
			fmt.Scanln(&dlt)
			if dlt < nPcs && dlt >= 0 {
				cPs[dlt] <- true
				close(cPs[dlt])
				delete(cPs, dlt)
			}
		}

	}
}
