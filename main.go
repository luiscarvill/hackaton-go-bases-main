package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	fil := file.File{}
	// 1. Se lee el archivo para sacar los registros
	tickets, err := fil.Read()
	if err != nil {
		//TODO retornar error en caso de error lectura
	}
	// Funcion para obtener tickets del archivo csv
	dataBook := service.NewBookings(tickets)

	//fmt.Println(dataBook)

	//=============  TEST ==================

	//=============  READ FILE ==================
	read1, _ := dataBook.Read(9)

	fmt.Println(read1)
	//=============  UPDATE FILE ==================
	update1, _ := dataBook.Update(1,
		service.Ticket{
			Date:        "1:44",
			Names:       "Luis Carvajal",
			Email:       "luiscarvajal@correo.com.co",
			Destination: "Colombia",
			Price:       530_000,
		})

	//read2, errRead := dataBook.Read(1)
	fmt.Println("Return update ==>", update1)
	//=============  DELETE FILE ==================
	deleteDt, _ := dataBook.Delete(2)
	fmt.Println("Return delete==>", deleteDt)
	//=============  CREATE FILE ==================
	dataBook.Create(
		service.Ticket{
			Id:          11,
			Date:        "1:44",
			Names:       "Luis H Carvajal",
			Email:       "luiscarvajal22@correo.com.co",
			Destination: "Colombia",
			Price:       532_000,
		})
	//read2, errRead := dataBook.Read(1)
	//fmt.Println(read2)

	//============= END TEST =============
	// WRITE FILE
	fil.Write(tickets)
	//fmt.Println(dataBook)
}
