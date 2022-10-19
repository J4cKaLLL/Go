package beers

// BeerUseCase es una interfaz que tienen las firmas de los métodos para esta workshop
type BeerUseCase interface {
	SearchBeers() (*[]Beer, error)
	AddBeers(Beer) error
}

type beerUseCase struct {
}

// NewBeerUseCase inyecta las dependencias y retorna la interfaz implementada en la estructura beerUseCase
func NewBeerUseCase() BeerUseCase {
	return &beerUseCase{}
}

func (b *beerUseCase) SearchBeers() (*[]Beer, error) {
	return nil, nil
}

func (b *beerUseCase) AddBeers(Beer) error {
	return nil
}
