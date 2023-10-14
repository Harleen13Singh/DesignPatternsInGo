package abstractfactory

type CricketFactory struct{
  adidasFactory *AdidasCricketFactory
  nikeFactory *NikeCricketFactory
}

func NewCricketFactory( adidasFactory *AdidasCricketFactory, nikeFactory *NikeCricketFactory) *CricketFactory{
  return &CricketFactory{
   adidasFactory: adidasFactory,
    nikeFactory: nikeFactory,
  }
}