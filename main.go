package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	fil := file.File{}
	// 1. Se lee el archivo para sacar los registros
	tickets, err := fil.Read()
	if err != nil {
		panic(err)
	}
	// Funcion para obtener tickets del archivo csv
	dataBook := service.NewBookings(tickets)

	//=============  TEST ==================
	//=============  READ FILE ==================
	read1, errRead := dataBook.Read(9)
	if errRead != nil {
		panic(errRead)
	}
	fmt.Println("Return Read ==>", read1)
	//=============  UPDATE FILE ==================
	update1, errUpd := dataBook.Update(1,
		service.Ticket{
			Date:        "1:44",
			Names:       "Luis Carvajal",
			Email:       "luiscarvajal@correo.com.co",
			Destination: "Colombia",
			Price:       530_000,
		})
	if errUpd != nil {
		panic(errUpd)
	}

	//read2, errRead := dataBook.Read(1)
	fmt.Println("Return update ==>", update1)
	//=============  DELETE FILE ==================
	deleteDt, errDt := dataBook.Delete(2)
	if errDt != nil {
		panic(errDt)
	}
	fmt.Println("Return delete==>", deleteDt)
	//=============  CREATE FILE ==================
	dataCreat, errCt := dataBook.Create(
		service.Ticket{
			Id:          11,
			Date:        "1:44",
			Names:       "Luis H Carvajal",
			Email:       "luiscarvajal22@correo.com.co",
			Destination: "Colombia",
			Price:       532_000,
		})
	if errCt != nil {
		panic(errCt)
	}
	fmt.Println("Return create ==>", dataCreat)
	//============= END TEST =============
	// WRITE FILE
	fil.Write(tickets)
}
