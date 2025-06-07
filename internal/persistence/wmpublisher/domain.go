package wmpublisher

import (
	"context"
	"encoding/json"

	"github.com/Arsfiqball/codec/internal/value/domain"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Domain struct {
	tracer trace.Tracer
	pub    message.Publisher
}

func NewDomain(
	tracer trace.Tracer,
	pub message.Publisher,
) *Domain {
	return &Domain{
		tracer: tracer,
		pub:    pub,
	}
}

func (d *Domain) Created(ctx context.Context, ent domain.Entity) error {
	_, span := d.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/persistence/wmpublisher/domain/Created")
	defer span.End()

	data, err := json.Marshal(domainEntityToDTO(ent))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to marshal domain entity")

		return err
	}

	msg := message.NewMessage(uuid.NewString(), data)

	return d.pub.Publish("domain_created", msg)
}

func (d *Domain) Updated(ctx context.Context, oldEnt domain.Entity, newEnt domain.Entity) error {
	_, span := d.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/persistence/wmpublisher/domain/Updated")
	defer span.End()

	data, err := json.Marshal(struct {
		Old domainDTO `json:"old"`
		New domainDTO `json:"new"`
	}{
		Old: domainEntityToDTO(oldEnt),
		New: domainEntityToDTO(newEnt),
	})

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to marshal domain entity")

		return err
	}

	msg := message.NewMessage(uuid.NewString(), data)

	return d.pub.Publish("domain_updated", msg)
}

func (d *Domain) Deleted(ctx context.Context, ent domain.Entity) error {
	_, span := d.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/persistence/wmpublisher/domain/Deleted")
	defer span.End()

	data, err := json.Marshal(domainEntityToDTO(ent))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to marshal domain entity")

		return err
	}

	msg := message.NewMessage(uuid.NewString(), data)

	return d.pub.Publish("domain_deleted", msg)
}

type domainDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func domainEntityToDTO(ent domain.Entity) domainDTO {
	return domainDTO{
		ID:       ent.ID(),
		Name:     ent.Name(),
		Email:    ent.Email(),
		Password: ent.Password(),
	}
}
