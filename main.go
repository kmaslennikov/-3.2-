package main

import "fmt"

func main() {
  fmt.Println("Прибываем на остановку Площадь Победы. В салоне пассажиров?")
  var total int
  var passenger int
  fmt.Scan(&passenger)
  total += passenger

  fmt.Println("Сколько пассажиров вышло на остановке?")
  var passengerComeOut int
  fmt.Scan(&passengerComeOut)

  fmt.Println("Сколько пассажиров вошло на остановке?")
  var passengerEntered int
  fmt.Scan(&passengerEntered)
  total += passengerEntered

  passenger = passenger - passengerComeOut + passengerEntered

  fmt.Println("Отправляемся с остановки Площадь Победы. В салоне пассажиров:" ,passenger)
  fmt.Println("-------Едем-------")

  fmt.Println("Прибываем на остановку Кинотеатр Центральный. В салоне пассажиров?", passenger)

  fmt.Println("Сколько пассажиров вышло на остановке?")
  fmt.Scan(&passengerComeOut)

  fmt.Println("Сколько пассажиров вошло на остановке?")
  fmt.Scan(&passengerEntered)
  total += passengerEntered

  passenger = passenger - passengerComeOut + passengerEntered

  fmt.Println("Отправляемся с остановки Кинотеатр Центральный. В салоне пассажиров:" ,passenger)
  fmt.Println("-------Едем-------")

  fmt.Println("Прибываем на остановку Сквер медсестер. В салоне пассажиров?", passenger)

  fmt.Println("Сколько пассажиров вышло на остановке?")
  fmt.Scan(&passengerComeOut)

  fmt.Println("Сколько пассажиров вошло на остановке?")
  fmt.Scan(&passengerEntered)
  total += passengerEntered

  passenger = passenger - passengerComeOut + passengerEntered

  fmt.Println("Отправляемся с остановки Сквер медсестер. В салоне пассажиров:" ,passenger)
  fmt.Println("-------Едем-------")

  fmt.Println("Прибываем на остановку Сквер Мира. В салоне пассажиров?", passenger)

  fmt.Println("Сколько пассажиров вышло на остановке?")
  fmt.Scan(&passengerComeOut)

  fmt.Println("Сколько пассажиров вошло на остановке?")
  fmt.Scan(&passengerEntered)
  total += passengerEntered

  passenger = passenger - passengerComeOut + passengerEntered

  fmt.Println("Отправляемся с остановки Сквер Мира. В салоне пассажиров:" ,passenger)
  ticketPrice := 20
  var money int
  var driverSalary int
  var gas int
  var taxes int
  var carRepairs int
  var income int
  money = total * ticketPrice
  driverSalary = money / 4
  gas = money / 5
  taxes = money / 5
  carRepairs = money / 5
  income = money - driverSalary - gas - taxes - carRepairs
  fmt.Println("Всего заработали:" , money , "руб")
  fmt.Println("Зарплата водителя: " , driverSalary , "руб")
  fmt.Println("Расходы на топливо:" , gas , "руб")
  fmt.Println("Расходы на налоги:" , taxes , "руб")
  fmt.Println("Расходы на ремонт машины:" , carRepairs , "руб")
  fmt.Println("Итого доход:" , income , "руб")



  


}