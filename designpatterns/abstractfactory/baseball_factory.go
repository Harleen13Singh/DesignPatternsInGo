package abstractfactory

type BaseBallFactory struct{
  bat *BaseBallBat
  gloves *BaseBallGloves
}

func NewBaseBallFactory(bat *BaseBallBat,gloves *BaseBallGloves) *BaseBallFactory{
  return &BaseBallFactory{
    bat: bat,
    gloves: gloves,
  }
}