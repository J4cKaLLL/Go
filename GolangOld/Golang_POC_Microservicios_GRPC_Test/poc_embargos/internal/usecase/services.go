package usecase

import "context"
import sizesures "sizesure/internal/entity"
import reposito "sizesure/internal/impl/impl_1"

// SizesuresUseCase Inteerfaz Caso de Uso
type SizesuresUseCase interface {
    AddSizesures(sizesures.Seizure) (sizesures.SeizureResp, error)
}

/* sizesuresUseCase define la estructura de usecase 
    en este caso la estructura es basada en el repository
*/
type sizesuresUseCase struct {
    repo reposito.SeizureRepository
}

// NewSizesuresUseCase inyeccion de depndecias y retorno de interfaz y estructura
func NewSizesuresUseCase(repo reposito.SeizureRepository) SizesuresUseCase {
    return &sizesuresUseCase{
        repo: repo,
    }
}

// AddSizesures se agrega la operacion del usecase
func (d *sizesuresUseCase) AddSizesures(req sizesures.Seizure) (sizesures.SeizureResp, error){
    ctx := context.Background()
	return d.repo.AddSeizures(ctx, req)
}
