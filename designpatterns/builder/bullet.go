package builder

type BulletBuilder struct {
  Engine string
  Tyre string
  Silencer string
}

var _ BuildBike = &BulletBuilder{}

func (b *BulletBuilder) SetEngine(){
  b.Engine = "EFI 350cc"
}
func (b *BulletBuilder) SetTyre(){
  b.Tyre = "Spokes tyre"
}
func (b *BulletBuilder) SetSilencer(){
  b.Silencer = "dug-dug"
}

func (b *BulletBuilder) GetBike() *Bike {
  return &Bike{
    Engine: b.Engine,
    Tyre: b.Tyre,
    Silencer: b.Silencer,
  }
}

func NewBulletBuilder() *BulletBuilder{
  return &BulletBuilder{}
}

