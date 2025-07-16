package entity

import (
	"time"
)

type Cliente struct {
	Chave 			string
	UnixBloqueio 	int64
	UnixRequest 	[]int64
}

func (c *Cliente) InsereNovaRequest(){

	c.UnixRequest = append(c.UnixRequest, time.Now().UnixMilli())
}

func (c *Cliente) CalculaNumeroRequestPorSegundo() int64{

	dataInicioCalculo := time.Now().Add(time.Second * -1).UnixMilli();

	var total int64 = 0
	for _, req := range c.UnixRequest {

		if req >= dataInicioCalculo {
			total++
		}
	}

	return total;

}