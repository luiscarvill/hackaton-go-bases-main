package service

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	_, err := FindIndex(b.Tickets, t.Id)
	if err == true {
		panic("El ticket a crear ya existe")
	}
	//TODO error cuando uno de los datos no exista. (Ticket)
	b.Tickets = append(b.Tickets, t)

	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, v := range b.Tickets {
		if v.Id == id {
			return v, nil
		}
	}

	//TODO error" no se escontro el ticket con el id enviado
	panic("no se encontro el id del ticket enviado")

}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for i, v := range b.Tickets {
		if v.Id == id {
			b.Tickets[i].Date = t.Date
			b.Tickets[i].Destination = t.Destination
			b.Tickets[i].Email = t.Email
			b.Tickets[i].Names = t.Names
			b.Tickets[i].Price = t.Price
			return v, nil
		}
	}
	panic("no se encontro el id del ticket enviado")
}

func (b *bookings) Delete(id int) (int, error) {
	index, err := FindIndex(b.Tickets, id)
	if err == true {
		panic("No se pudo encontrar el ticket a borrar")
	}
	b.Tickets = append(b.Tickets[0:index],
		b.Tickets[index+1:len(b.Tickets)]...)

	return id, nil
}

func FindIndex(t []Ticket, id int) (int, bool) {
	for idTick, v := range t {
		if v.Id == id {
			return idTick, false
		}
	}
	return 0, true
}
