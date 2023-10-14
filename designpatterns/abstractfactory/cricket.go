package abstractfactory

type AdidasCricketFactory struct {
  gloves *AdidasCricketGloves
  bat *AdidasCricketBat
}

func NewAdidasCricketFactory( gloves *AdidasCricketGloves, bat *AdidasCricketBat) *AdidasCricketFactory{
  return &AdidasCricketFactory{
    gloves: gloves,
    bat: bat,
  }
}

type NikeCricketFactory struct{
  gloves *NikeCricketGloves
  bat *NikeCricketBat
}

func NewNikeCricketFactory( gloves *NikeCricketGloves, bat *NikeCricketBat) *NikeCricketFactory{
  return &NikeCricketFactory{
    gloves: gloves,
    bat: bat,
  }
}



type AdidasCricketBat struct{}

func NewAdidasCricketBat() *AdidasCricketBat{
  return &AdidasCricketBat{}
}
type AdidasCricketGloves struct{}
func NewAdidasCricketGloves() *AdidasCricketGloves{
  return &AdidasCricketGloves{}
}
func (c *AdidasCricketBat) GetBatType() string {
	return "adidas_cricket_bat"
}
func (c *AdidasCricketGloves) GetGlovesType() string {
	return "adidas_cricket_Gloves"
}

type NikeCricketBat struct{}
func NewNikeCricketBat() *NikeCricketBat{
  return &NikeCricketBat{}
}
type NikeCricketGloves struct{}
func NewNikeCricketGloves() *NikeCricketGloves{
  return &NikeCricketGloves{}
}
func (c *NikeCricketBat) GetBatType() string {
  return "nike_cricket_bat"
}
func (c *NikeCricketGloves) GetGlovesType() string {
  return "nike_cricket_Gloves"
}

