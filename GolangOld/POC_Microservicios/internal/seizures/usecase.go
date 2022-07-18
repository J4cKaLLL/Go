package seizures

// SeizureUseCase Interfaz Caso de uso
type SeizureUseCase interface {
	SeizureCreate(Seizure) error
}
