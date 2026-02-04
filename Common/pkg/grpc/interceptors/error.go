package interceptors

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
)

// ServerErrorInterceptor returns a new unary server interceptor that handles error mapping.
func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, mapErrorToGRPCStatus(err)
		}
		return resp, nil
	}
}

// mapErrorToGRPCStatus maps app errors to gRPC status errors.
func mapErrorToGRPCStatus(err error) error {
	var appErr *apperr.AppError
	if errors.As(err, &appErr) {
		code := codes.Internal

		// Map HTTP status codes/AppErr codes to gRPC codes
		// Using HTTPStatus as primary mapping source since it's well defined in AppErr
		switch appErr.HTTPStatus {
		case http.StatusBadRequest:
			code = codes.InvalidArgument
		case http.StatusUnauthorized:
			code = codes.Unauthenticated
		case http.StatusForbidden:
			code = codes.PermissionDenied
		case http.StatusNotFound:
			code = codes.NotFound
		case http.StatusConflict:
			code = codes.AlreadyExists // Or codes.Aborted depending on context
		case http.StatusTooManyRequests:
			code = codes.ResourceExhausted
		case http.StatusInternalServerError:
			code = codes.Internal
		case http.StatusNotImplemented:
			code = codes.Unimplemented
		case http.StatusServiceUnavailable:
			code = codes.Unavailable
		case http.StatusGatewayTimeout:
			code = codes.DeadlineExceeded
		default:
			// Fallback to mapping based on custom AppErr codes if needed,
			// currently mapping common ones.
			if appErr.Code == response.CodeParamInvalid {
				code = codes.InvalidArgument
			} else if appErr.Code == response.CodeUnauthorized {
				code = codes.Unauthenticated
			} else if appErr.Code == response.CodeNotFound {
				code = codes.NotFound
			} else if appErr.Code == response.CodeConflict {
				code = codes.AlreadyExists
			}
		}

		return status.Error(code, appErr.Message)
	}

	// If it's already a gRPC status error, return it as is
	if _, ok := status.FromError(err); ok {
		return err
	}

	// For unknown errors, return Internal
	return status.Error(codes.Internal, err.Error())
}

// FromGRPCStatus converts a gRPC error to an AppError.
func FromGRPCStatus(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		// Not a gRPC error, return as is
		return err
	}

	var httpCode int
	var errCode int

	switch st.Code() {
	case codes.OK:
		return nil
	case codes.InvalidArgument:
		httpCode = http.StatusBadRequest
		errCode = response.CodeParamInvalid
	case codes.Unauthenticated:
		httpCode = http.StatusUnauthorized
		errCode = response.CodeUnauthorized
	case codes.PermissionDenied:
		httpCode = http.StatusForbidden
		errCode = response.CodeForbidden
	case codes.NotFound:
		httpCode = http.StatusNotFound
		errCode = response.CodeNotFound
	case codes.AlreadyExists:
		httpCode = http.StatusConflict
		errCode = response.CodeConflict
	case codes.ResourceExhausted:
		httpCode = http.StatusTooManyRequests
		errCode = response.CodeTooManyRequests
	case codes.Unimplemented:
		httpCode = http.StatusNotImplemented
		errCode = response.CodeInternalError
	case codes.Unavailable:
		httpCode = http.StatusServiceUnavailable
		errCode = response.CodeInternalError
	case codes.DeadlineExceeded:
		httpCode = http.StatusGatewayTimeout
		errCode = response.CodeInternalError
	default:
		httpCode = http.StatusInternalServerError
		errCode = response.CodeInternalError
	}

	return apperr.NewError(
		"gRPC Client",
		errCode,
		st.Message(),
		httpCode,
		nil,
	)
}
