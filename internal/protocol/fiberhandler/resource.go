package fiberhandler

import (
	"github.com/Arsfiqball/codec/internal/application/resource"
	"github.com/Arsfiqball/codec/internal/protocol/fiberhandler/formatter"
	"github.com/Arsfiqball/codec/internal/protocol/fiberhandler/parser"
	"github.com/Arsfiqball/codec/internal/value/domain"
	"github.com/Arsfiqball/codec/internal/value/user"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

type Resource struct {
	tracer  trace.Tracer
	service resource.IService
}

func NewResource(tracer trace.Tracer, svc resource.IService) *Resource {
	return &Resource{
		tracer:  tracer,
		service: svc,
	}
}

func (r *Resource) Create(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "github.com/Arsfiqball/codec/internal/protocol/fiberhandler/resource/Create")
	defer span.End()

	var (
		patch domain.Patch
		user  user.Identity
	)

	if err := parser.UserIdentityToken(c.Locals("user").(string), &user); err != nil {
		return FormatErrorWithCode(c, err, codeErrUnauthorized)
	}

	if err := parser.DomainPatchJSON(c.Body(), &patch); err != nil {
		return FormatErrorWithCode(c, err, codeErrBadRequest)
	}

	ent, err := r.service.Create(ctx, patch, user)
	if err != nil {
		return FormatError(c, err)
	}

	return FormatSuccess(c, formatter.FormatDomainEntityDTO(ent))
}

func (r *Resource) Update(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "github.com/Arsfiqball/codec/internal/protocol/fiberhandler/resource/Update")
	defer span.End()

	var (
		query domain.Query
		patch domain.Patch
		user  user.Identity
	)

	if err := parser.UserIdentityToken(c.Locals("user").(string), &user); err != nil {
		return FormatErrorWithCode(c, err, codeErrUnauthorized)
	}

	if err := parser.DomainQueryString(c.Request().URI().QueryString(), &query); err != nil {
		return FormatErrorWithCode(c, err, codeErrBadRequest)
	}

	if err := parser.DomainPatchJSON(c.Body(), &patch); err != nil {
		return FormatErrorWithCode(c, err, codeErrBadRequest)
	}

	ent, err := r.service.Update(ctx, query, patch, user)
	if err != nil {
		return FormatError(c, err)
	}

	return FormatSuccess(c, formatter.FormatDomainEntityDTO(ent))
}

func (r *Resource) Delete(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "github.com/Arsfiqball/codec/internal/protocol/fiberhandler/resource/Delete")
	defer span.End()

	var (
		query domain.Query
		user  user.Identity
	)

	if err := parser.UserIdentityToken(c.Locals("user").(string), &user); err != nil {
		return FormatErrorWithCode(c, err, codeErrUnauthorized)
	}

	if err := parser.DomainQueryString(c.Request().URI().QueryString(), &query); err != nil {
		return FormatErrorWithCode(c, err, codeErrBadRequest)
	}

	ent, err := r.service.Delete(ctx, query, user)
	if err != nil {
		return FormatError(c, err)
	}

	return FormatSuccess(c, formatter.FormatDomainEntityDTO(ent))
}

func (r *Resource) GetOne(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "github.com/Arsfiqball/codec/internal/protocol/fiberhandler/resource/GetOne")
	defer span.End()

	var (
		query domain.Query
		user  user.Identity
	)

	if err := parser.UserIdentityToken(c.Locals("user").(string), &user); err != nil {
		return FormatErrorWithCode(c, err, codeErrUnauthorized)
	}

	if err := parser.DomainQueryString(c.Request().URI().QueryString(), &query); err != nil {
		return FormatErrorWithCode(c, err, codeErrBadRequest)
	}

	ent, err := r.service.GetOne(ctx, query, user)
	if err != nil {
		return FormatError(c, err)
	}

	return FormatSuccess(c, formatter.FormatDomainEntityDTO(ent))
}

func (r *Resource) GetList(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "github.com/Arsfiqball/codec/internal/protocol/fiberhandler/resource/GetList")
	defer span.End()

	var (
		query domain.Query
		user  user.Identity
	)

	if err := parser.UserIdentityToken(c.Locals("user").(string), &user); err != nil {
		return FormatErrorWithCode(c, err, codeErrUnauthorized)
	}

	if err := parser.DomainQueryString(c.Request().URI().QueryString(), &query); err != nil {
		return FormatErrorWithCode(c, err, codeErrBadRequest)
	}

	list, err := r.service.GetList(ctx, query, user)
	if err != nil {
		return FormatError(c, err)
	}

	dtos := []formatter.DomainEntityDTO{}

	for _, ent := range list {
		dtos = append(dtos, formatter.FormatDomainEntityDTO(ent))
	}

	return FormatSuccess(c, dtos)
}

func (r *Resource) GetStat(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "github.com/Arsfiqball/codec/internal/protocol/fiberhandler/resource/GetStat")
	defer span.End()

	var (
		query domain.Query
		user  user.Identity
	)

	if err := parser.UserIdentityToken(c.Locals("user").(string), &user); err != nil {
		return FormatErrorWithCode(c, err, codeErrUnauthorized)
	}

	if err := parser.DomainQueryString(c.Request().URI().QueryString(), &query); err != nil {
		return FormatErrorWithCode(c, err, codeErrBadRequest)
	}

	stat, err := r.service.GetStat(ctx, query, user)
	if err != nil {
		return FormatError(c, err)
	}

	dtos := []formatter.DomainStatDTO{}

	for _, ent := range stat {
		dtos = append(dtos, formatter.FormatDomainStatDTO(ent))
	}

	return FormatSuccess(c, dtos)
}

func (r *Resource) BulkOps(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "github.com/Arsfiqball/codec/internal/protocol/fiberhandler/resource/BulkOps")
	defer span.End()

	var (
		ops  []resource.Ops
		user user.Identity
	)

	if err := parser.UserIdentityToken(c.Locals("user").(string), &user); err != nil {
		return FormatErrorWithCode(c, err, codeErrUnauthorized)
	}

	if err := parser.DomainBulkOpsJSON(c.Body(), &ops); err != nil {
		return FormatErrorWithCode(c, err, codeErrBadRequest)
	}

	list, err := r.service.BulkOps(ctx, ops, user)
	if err != nil {
		return FormatError(c, err)
	}

	return FormatSuccess(c, list)
}
