// bat and gloves are 2 products that client will need
package abstractfactory

type Bat interface {
	GetBatType() string
}
type Gloves interface {
	GetGlovesType() string
}

// this is the final thing that the product needs, since they need in them in a group
// so added them in a struct
type SportsProducts struct{
  Bat Bat
  Gloves Gloves
}

