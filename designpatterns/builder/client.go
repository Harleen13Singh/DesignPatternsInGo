package builder

import "fmt"

func ClientCode() {
  // creation of bullet bike withoud director
  bulletBuilder := NewBikeBuilder("bullet")
  bulletBuilder.SetEngine()
  bulletBuilder.SetTyre()
  bulletBuilder.SetSilencer()
  bike := bulletBuilder.GetBike()
  fmt.Println(bike)

  // creation of yamaha bike with director
  yamahaBuilder := NewBikeBuilder("yamaha")
  bikeDirector := NewBikeCreationDirector(yamahaBuilder)
  fmt.Println(bikeDirector.GetBike())
}