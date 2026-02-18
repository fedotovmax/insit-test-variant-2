package ports

import "github.com/fedotovmax/insit-test-variant-2/text_receiver/internal/domain"

type EventPublisher interface {
	Send(newOp *domain.ToProcessOperation) error
}
