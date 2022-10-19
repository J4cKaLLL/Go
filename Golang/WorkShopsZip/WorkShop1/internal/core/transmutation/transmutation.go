package transmutation

import (
	"github.com/jinzhu/copier"

	entity "corp/fif/inte/customers/internal/core/entity"
	pbFind "corp/fif/inte/customers/internal/core/repository/find_customer/proto"
	pb "corp/fif/inte/customers/internal/entrypoint/api/grpc/proto"
)

// Transmutation ...
type Transmutation interface {
	EncodeEntityPB(*entity.Customer) (*pb.Customer, error)
	EncodePBEntityFind(*pbFind.Customer) (*entity.Customer, error)
}

type transmutation struct {
}

// NewTransmutation ...
func NewTransmutation() Transmutation {
	return &transmutation{}
}

// EncodeEntityPB ...
func (t *transmutation) EncodeEntityPB(cu *entity.Customer) (*pb.Customer, error) {
	from := &cu
	to := new(pb.Customer)

	err := copier.Copy(to, from)
	if err != nil {
		return nil, err
	}
	return to, nil
}

// EncodePBEntityAR ...
func (t *transmutation) EncodePBEntityFind(cu *pbFind.Customer) (*entity.Customer, error) {
	from := &cu
	to := new(entity.Customer)

	err := copier.Copy(to, from)
	if err != nil {
		return nil, err
	}
	return to, nil
}
