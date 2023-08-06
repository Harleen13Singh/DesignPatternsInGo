package abstractfactory

type CricketBat struct{}
func NewCricketBat() *CricketBat{
  return &CricketBat{}
}
type CricketGloves struct{}
func NewCricketGloves() *CricketGloves{
  return &CricketGloves{}
}
func (c *CricketBat) GetBatType() string {
	return "cricket_bat"
}
func (c *CricketGloves) GetGlovesType() string {
	return "cricket_Gloves"
}