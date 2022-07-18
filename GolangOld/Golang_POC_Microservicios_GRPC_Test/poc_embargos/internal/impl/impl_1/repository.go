package impl_1

import (
	"fmt"
	"context"
	"strconv"
	en "sizesure/internal/entity"
	grp "sizesure/cmd/sizesure/proto"

	"google.golang.org/grpc"
	
	
)

// SeizureRepository interface con los métodos implementados
type SeizureRepository interface {
	AddSeizures(ctx context.Context, seizure en.Seizure) (en.SeizureResp, error)
}

/* seizureRepository estructura del repository 
	en este caso la estructura es basada en el cliente grpc
*/
type seizureRepository struct {
	cc *grpc.ClientConn
}

// NewSeizureRepository constructor para repository
func NewSeizureRepository(cc *grpc.ClientConn) SeizureRepository {
	return &seizureRepository{
		cc: cc,
	}
}

// AddSeizures se agrega la operacion del repository
func (r *seizureRepository) AddSeizures(ctx context.Context, req en.Seizure) (en.SeizureResp, error) {
	// se inicializa el cliente grpc
	conn := grp.NewSeizureServiceClient(r.cc)
	// se invoca la operación del grpc con los parametros de entrada
	// en este caso se envia el context y el request en su entity
	// el request debe estar con & (apuntador de dirección de memoria) ya que grpc lo requiere de esta manera
	responsegrpc, err := conn.InsertSeizure(ctx, &grp.SeizureRequest{
		Id:       int32(0),
		Customer: &grp.Client{
			FirstName:   req.Customer.FirstName,
			MiddleName:  req.Customer.MiddleName,
			Surname:     req.Customer.Surname,
			LastSurname: req.Customer.LastSurname,
			Document: &grp.DocumentInfo{
				DocumentType:   req.Customer.Document.DocumentType,
				DocumentNumber: req.Customer.Document.DocumentNumber,
			},
		},
		SeizureInformation: &grp.SeizeInfo{
			Concept:         int32(req.SeizureInformation.Concept),
			SeizureType:     int32(req.SeizureInformation.SeizureType),
			DebitForm:       int32(req.SeizureInformation.DebitForm),
			ReceptionDate:   req.SeizureInformation.ReceptionDate,
			ProcessNumber:   req.SeizureInformation.ProcessNumber,
			Amount:          int32(req.SeizureInformation.Amount),
			City:            req.SeizureInformation.City,
			Address:         req.SeizureInformation.Address,
			JudicialAccount: req.SeizureInformation.JudicialAccount,
			Entity:          req.SeizureInformation.Entity,
			SeizureDate:     req.SeizureInformation.SeizureDate,
			Plaintiff: &grp.Client{
				FirstName:   req.SeizureInformation.Plaintiff.FirstName,
				MiddleName:  req.SeizureInformation.Plaintiff.MiddleName,
				Surname:     req.SeizureInformation.Plaintiff.Surname,
				LastSurname: req.SeizureInformation.Plaintiff.LastSurname,
				Document: &grp.DocumentInfo{
					DocumentType:   req.SeizureInformation.Plaintiff.Document.DocumentType,
					DocumentNumber: req.SeizureInformation.Plaintiff.Document.DocumentNumber,
				},
			},
		},
		Product: &grp.ProductName{
			ProductID:    req.Product.ProductID,
			ProductState: req.Product.ProductState,
		},
	})

	// se transforma la salida del grpc al formato entity del usecase
	respSeizure := en.SeizureResp{
		Codigo: strconv.FormatInt(int64(responsegrpc.GetId()), 10),
		Error: "1",
		Mensaje: "Prueba mensaje respuesta",
	}
	fmt.Println("Result: %d",responsegrpc.GetId())
	// retorna los parametros definidos en el metodo
	// en este caso retorna el response del entity del usecase y el error generado
	return  respSeizure, err
}