package builder

type BikeCreationDirector struct {
  bikeBuilder BuildBike
}

func NewBikeCreationDirector (builder BuildBike) *BikeCreationDirector{
  return &BikeCreationDirector{
    bikeBuilder: builder,
  }
}

func (bd *BikeCreationDirector) GetBike() *Bike{
  bd.bikeBuilder.SetEngine()
  bd.bikeBuilder.SetTyre()
  bd.bikeBuilder.SetSilencer()
  return bd.bikeBuilder.GetBike()
}