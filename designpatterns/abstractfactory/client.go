package abstractfactory

import "fmt"

func ClientCode() {
  cricBat := NewCricketBat()
	cricGloves := NewCricketGloves()

	baseBat := NewBaseBallBat()
	baseGloves := NewBaseBallGloves()

	cricfactory := NewCricketFactory(cricBat, cricGloves)
	baseFactory := NewBaseBallFactory(baseBat, baseGloves)

	sportsFactory := NewSportsFactory(cricfactory, baseFactory)

	sportsProducts := sportsFactory.GetSportsProduct("cricket")
	fmt.Println(sportsProducts.Bat.GetBatType())
	fmt.Println(sportsProducts.Gloves.GetGlovesType())

	sportsProducts = sportsFactory.GetSportsProduct("baseball")
	fmt.Println(sportsProducts.Bat.GetBatType())
	fmt.Println(sportsProducts.Gloves.GetGlovesType())
}