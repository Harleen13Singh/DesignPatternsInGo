package builder

// BuildBike is the builder interface responsible for building the BuildBike
// Each represenation of the bike will implement this interface for its creation
type BuildBike interface{
  SetEngine()
  SetTyre()
  SetSilencer()
  GetBike() *Bike
}

func NewBikeBuilder(bikeName string) BuildBike {
  switch bikeName {
    case "bullet":
    return NewBulletBuilder()
    
    case "yamaha":
    return NewYamahaBuilder()
    
    default:
    return nil
  }
}