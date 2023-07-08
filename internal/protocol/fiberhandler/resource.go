package fiberhandler

import (
	"feature/internal/application/resource"
	"feature/internal/protocol/fiberhandler/parser"
	"feature/internal/value/domain"
	"feature/internal/value/user"

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
	ctx, span := r.tracer.Start(c.UserContext(), "feature/internal/protocol/fiberhandler/resource/Create")
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

	return FormatSuccess(c, ent)
}

func (r *Resource) Update(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "feature/internal/protocol/fiberhandler/resource/Update")
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

	return FormatSuccess(c, ent)
}

func (r *Resource) Delete(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "feature/internal/protocol/fiberhandler/resource/Delete")
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

	return FormatSuccess(c, ent)
}

func (r *Resource) GetOne(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "feature/internal/protocol/fiberhandler/resource/GetOne")
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

	return FormatSuccess(c, ent)
}

func (r *Resource) GetList(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "feature/internal/protocol/fiberhandler/resource/GetList")
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

	return FormatSuccess(c, list)
}

func (r *Resource) GetStat(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "feature/internal/protocol/fiberhandler/resource/GetStat")
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

	return FormatSuccess(c, stat)
}

func (r *Resource) BulkOps(c *fiber.Ctx) error {
	ctx, span := r.tracer.Start(c.UserContext(), "feature/internal/protocol/fiberhandler/resource/BulkOps")
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
