package builder

type YamahaBuilder struct {
  Engine string
  Tyre string
  Silencer string
}

var _ BuildBike = &YamahaBuilder{}

func (y *YamahaBuilder) SetEngine(){
  y.Engine = "normal engine"
}
func (y *YamahaBuilder) SetTyre(){
  y.Tyre = "Spokes tyre"
}
func (y *YamahaBuilder) SetSilencer(){
  y.Silencer = "drum-drum"
}

func (y *YamahaBuilder) GetBike() *Bike {
  return &Bike{
    Engine: y.Engine,
    Tyre: y.Tyre,
    Silencer: y.Silencer,
  }
}


func NewYamahaBuilder() *YamahaBuilder {
  return &YamahaBuilder{}
}