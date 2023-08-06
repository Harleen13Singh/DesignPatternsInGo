package abstractfactory

type BaseBallBat struct{}
func NewBaseBallBat() *BaseBallBat{
  return &BaseBallBat{}
}
type BaseBallGloves struct{}
func NewBaseBallGloves() *BaseBallGloves{
  return &BaseBallGloves{}
}
func (b *BaseBallBat) GetBatType() string {
	return "baseball_bat"
}
func (b *BaseBallGloves) GetGlovesType() string {
	return "baseball_Gloves"
}