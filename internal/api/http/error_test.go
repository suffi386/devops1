package http

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	caos_errors "github.com/zitadel/zitadel/internal/errors"
)

func TestZitadelErrorToHTTPStatusCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantOk         bool
	}{
		{
			name: "no error",
			args: args{
				err: nil,
			},
			wantStatusCode: http.StatusOK,
			wantOk:         true,
		},
		{
			name: "wrapped already exists",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowAlreadyExists(nil, "id", "message")),
			},
			wantStatusCode: http.StatusConflict,
			wantOk:         true,
		},
		{
			name: "wrapped deadline exceeded",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowDeadlineExceeded(nil, "id", "message")),
			},
			wantStatusCode: http.StatusGatewayTimeout,
			wantOk:         true,
		},
		{
			name: "wrapped internal",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowInternal(nil, "id", "message")),
			},
			wantStatusCode: http.StatusInternalServerError,
			wantOk:         true,
		},
		{
			name: "wrapped invalid argument",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowInvalidArgument(nil, "id", "message")),
			},
			wantStatusCode: http.StatusBadRequest,
			wantOk:         true,
		},
		{
			name: "wrapped not found",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowNotFound(nil, "id", "message")),
			},
			wantStatusCode: http.StatusNotFound,
			wantOk:         true,
		},
		{
			name: "wrapped permission denied",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowPermissionDenied(nil, "id", "message")),
			},
			wantStatusCode: http.StatusForbidden,
			wantOk:         true,
		},
		{
			name: "wrapped precondition failed",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowPreconditionFailed(nil, "id", "message")),
			},
			wantStatusCode: http.StatusBadRequest,
			wantOk:         true,
		},
		{
			name: "wrapped unauthenticated",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowUnauthenticated(nil, "id", "message")),
			},
			wantStatusCode: http.StatusUnauthorized,
			wantOk:         true,
		},
		{
			name: "wrapped unavailable",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowUnavailable(nil, "id", "message")),
			},
			wantStatusCode: http.StatusServiceUnavailable,
			wantOk:         true,
		},
		{
			name: "wrapped unimplemented",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowUnimplemented(nil, "id", "message")),
			},
			wantStatusCode: http.StatusNotImplemented,
			wantOk:         true,
		},
		{
			name: "wrapped resource exhausted",
			args: args{
				err: fmt.Errorf("wrapped %w", caos_errors.ThrowResourceExhausted(nil, "id", "message")),
			},
			wantStatusCode: http.StatusTooManyRequests,
			wantOk:         true,
		},
		{
			name: "wrapped resource exhausted",
			args: args{
				err: errors.New("error"),
			},
			wantStatusCode: http.StatusInternalServerError,
			wantOk:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatusCode, gotOk := ZitadelErrorToHTTPStatusCode(tt.args.err)
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("ZitadelErrorToHTTPStatusCode() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
			if gotOk != tt.wantOk {
				t.Errorf("ZitadelErrorToHTTPStatusCode() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
