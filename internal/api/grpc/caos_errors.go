package grpc

import (
	"context"

	"github.com/caos/logging"
	"github.com/caos/zitadel/pkg/message"

	caos_errs "github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/i18n"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CaosToGRPCError(err error, ctx context.Context, translator *i18n.Translator) error {
	if err == nil {
		return nil
	}
	code, key, id, ok := ExtractCaosError(err)
	if !ok {
		return status.Convert(err).Err()
	}
	msg := key
	if translator != nil {
		msg = translator.LocalizeFromCtx(ctx, key, nil)
	}
	s, err := status.New(code, key).WithDetails(&message.ErrorDetail{Id: id, Message: msg})
	if err != nil {
		logging.Log("GRPC-gIeRw").WithError(err).Debug("unable to add detail")
		return status.New(code, key).Err()
	}

	return s.Err()
}

func ExtractCaosError(err error) (c codes.Code, msg, id string, ok bool) {
	switch caosErr := err.(type) {
	case *caos_errs.AlreadyExistsError:
		return codes.AlreadyExists, caosErr.GetMessage(), caosErr.GetID(), true
	case *caos_errs.DeadlineExceededError:
		return codes.DeadlineExceeded, caosErr.GetMessage(), caosErr.GetID(), true
	case caos_errs.InternalError:
		return codes.Internal, caosErr.GetMessage(), caosErr.GetID(), true
	case *caos_errs.InvalidArgumentError:
		return codes.InvalidArgument, caosErr.GetMessage(), caosErr.GetID(), true
	case *caos_errs.NotFoundError:
		return codes.NotFound, caosErr.GetMessage(), caosErr.GetID(), true
	case *caos_errs.PermissionDeniedError:
		return codes.PermissionDenied, caosErr.GetMessage(), caosErr.GetID(), true
	case *caos_errs.PreconditionFailedError:
		return codes.FailedPrecondition, caosErr.GetMessage(), caosErr.GetID(), true
	case *caos_errs.UnauthenticatedError:
		return codes.Unauthenticated, caosErr.GetMessage(), caosErr.GetID(), true
	case *caos_errs.UnavailableError:
		return codes.Unavailable, caosErr.GetMessage(), caosErr.GetID(), true
	case *caos_errs.UnimplementedError:
		return codes.Unimplemented, caosErr.GetMessage(), caosErr.GetID(), true
	default:
		return codes.Unknown, err.Error(), "", false
	}
}
