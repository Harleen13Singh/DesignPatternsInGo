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

func (s *SportsFactory) GetSportsProduct(sportType string,sportsCompany string) *SportsProducts {
	switch {
	case sportType == "cricket":
    if sportsCompany == "adidas"{
      return &SportsProducts{
        Bat: s.cricketFactory.adidasFactory.bat,
        Gloves: s.cricketFactory.adidasFactory.gloves,
      }
    }
    return &SportsProducts{
      Bat: s.cricketFactory.nikeFactory.bat,
      Gloves: s.cricketFactory.nikeFactory.gloves,
    }
	case sportType == "baseball":
		return &SportsProducts{
			Bat:    s.baseBallFactory.bat,
			Gloves: s.baseBallFactory.gloves,
		}
	}
  return nil
}
