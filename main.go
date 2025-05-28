package main

import (
	"fmt"
	"naztransport"
)

func main() {
	nazTransport := naztransport.NewNazTransport(
		"Автомобиль", 1500, 1000000, "Легковой", 200, 
		"Пассажирский", "Седан", "Бензиновый", 8,
	)

	nazTransport.Vkluchit()
	nazTransport.ZavestiDvigatel()
	nazTransport.UvelichitSkorost(50)
	nazTransport.PomenyatKuzov("Хэтчбек")
	nazTransport.ProveritRashod()
	nazTransport.ZalitToplivo(30)
	nazTransport.SetCena(1200000)

	fmt.Printf("Новая цена: %d руб.\n", nazTransport.GetCena())

	nazTransport.SetSkorost(180)
	fmt.Printf("Новая скорость: %d км/ч.\n", nazTransport.GetSkorost())

	nazTransport.SetTipDvigatelia("Дизельный")
	fmt.Printf("Новый тип двигателя: %s.\n", nazTransport.GetTipDvigatelia())
}
