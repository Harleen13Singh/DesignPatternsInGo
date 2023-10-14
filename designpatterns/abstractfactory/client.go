package abstractfactory

import "fmt"

func ClientCode() {
  adidasCricketBat := NewAdidasCricketBat()
  AdidasCricketGloves := NewAdidasCricketGloves()

  adidasCricketFactory := NewAdidasCricketFactory(AdidasCricketGloves,adidasCricketBat)

  nikeCricketBat := NewNikeCricketBat()
  nikeCricketGloves := NewNikeCricketGloves()

  nikeFactory := NewNikeCricketFactory(nikeCricketGloves,nikeCricketBat)
  
	cricfactory := NewCricketFactory(adidasCricketFactory, nikeFactory)

baseBat := NewBaseBallBat()
  baseGloves := NewBaseBallGloves()
  
	baseFactory := NewBaseBallFactory(baseBat, baseGloves)

	sportsFactory := NewSportsFactory(cricfactory, baseFactory)

	sportsProducts := sportsFactory.GetSportsProduct("cricket","adidas")
	fmt.Println(sportsProducts.Bat.GetBatType())
	fmt.Println(sportsProducts.Gloves.GetGlovesType())

  sportsProducts = sportsFactory.GetSportsProduct("cricket","nike")
  fmt.Println(sportsProducts.Bat.GetBatType())
  fmt.Println(sportsProducts.Gloves.GetGlovesType())

	sportsProducts = sportsFactory.GetSportsProduct("baseball","nike")
	fmt.Println(sportsProducts.Bat.GetBatType())
	fmt.Println(sportsProducts.Gloves.GetGlovesType())
}