package seizures

// SeizureUseCase Interfaz Caso de uso
type SeizureUseCase interface {
	CreateSeizure(Seizure) error
}

//2. retorna la estructura
type seizureUseCase struct {
}

// NewSeizureUseCase crear funcion para inyectar independencias y retornar la interfaz
func NewSeizureUseCase() SeizureUseCase {
	return &seizureUseCase{}
}

//3. se crea los metodos de la interfaz
func (s *seizureUseCase) CreateSeizure(Seizure) error {
	return nil
}
