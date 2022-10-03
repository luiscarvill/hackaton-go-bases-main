package file

import (
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func (f *File) definePath() {
	pwd, _ := os.Getwd()
	f.path = pwd + "/tickets.csv"

}
func (f *File) Read() ([]service.Ticket, error) {
	var tickts []service.Ticket

	f.definePath()

	data, err := os.ReadFile(f.path)
	if err != nil {
		panic("Error al leer el archivo")
	}
	splitDt := strings.Split(string(data), "\n")

	for i := 0; i < len(splitDt)-1; i++ {
		tick := strings.Split(splitDt[i], ",")
		id, _ := strconv.Atoi(tick[0])
		price, _ := strconv.Atoi(tick[5])
		dt := service.Ticket{
			Id:          id,
			Names:       tick[1],
			Email:       tick[2],
			Destination: tick[3],
			Date:        tick[4],
			Price:       price,
		}
		tickts = append(tickts, dt)
	}
	return tickts, nil
}

func (f *File) Write(s []service.Ticket) error {

	f.definePath()
	strDt, err := OrderData(s)
	if err != nil {
		panic("Error al ordenar la informacion para generar el archivo")
	}
	os.WriteFile(f.path, []byte(strDt), 644)

	return nil
}

func OrderData(s []service.Ticket) (string, error) {
	var str string

	for _, v := range s {
		spr := ","
		str += strconv.Itoa(v.Id) + spr + v.Names + spr + v.Email +
			spr + v.Destination + spr + v.Date + spr + strconv.Itoa(v.Price) + "\n"
	}
	return str, nil
}
