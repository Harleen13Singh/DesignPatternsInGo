package abstractfactory

type SportsFactory struct {
	cricketFactory  *CricketFactory
	baseBallFactory *BaseBallFactory
}

func NewSportsFactory(cricketFactory *CricketFactory, baseBallFactory *BaseBallFactory) *SportsFactory {
	return &SportsFactory{
		cricketFactory:  cricketFactory,
		baseBallFactory: baseBallFactory,
	}
}

func (s *SportsFactory) GetSportsProduct(sportType string) *SportsProducts {
	switch {
	case sportType == "cricket":
		return &SportsProducts{
			Bat:    s.cricketFactory.bat,
			Gloves: s.cricketFactory.gloves,
		}
	case sportType == "baseball":
		return &SportsProducts{
			Bat:    s.baseBallFactory.bat,
			Gloves: s.baseBallFactory.gloves,
		}
	}
  return nil
}
