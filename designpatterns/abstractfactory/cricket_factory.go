package abstractfactory

type CricketFactory struct{
  bat *CricketBat
  gloves *CricketGloves
}

func NewCricketFactory(bat *CricketBat,gloves *CricketGloves) *CricketFactory{
  return &CricketFactory{
    bat: bat,
    gloves: gloves,
  }
}