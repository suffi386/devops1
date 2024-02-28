//go:build integration

package admin_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zitadel/zitadel/internal/integration"
	admin_pb "github.com/zitadel/zitadel/pkg/grpc/admin"
	"github.com/zitadel/zitadel/pkg/grpc/member"
	"github.com/zitadel/zitadel/pkg/grpc/object"
)

var iamRoles = []string{
	"IAM_OWNER",
	"IAM_OWNER_VIEWER",
	"IAM_ORG_MANAGER",
	"IAM_USER_MANAGER",
	"IAM_ADMIN_IMPERSONATOR",
	"IAM_END_USER_IMPERSONATOR",
}

func TestServer_ListIAMMemberRoles(t *testing.T) {
	got, err := Client.ListIAMMemberRoles(AdminCTX, &admin_pb.ListIAMMemberRolesRequest{})
	require.NoError(t, err)
	assert.ElementsMatch(t, iamRoles, got.GetRoles())
}

func TestServer_ListIAMMembers(t *testing.T) {
	user := Tester.CreateHumanUserVerified(AdminCTX, Tester.Organisation.ID, gofakeit.Email())
	_, err := Client.AddIAMMember(AdminCTX, &admin_pb.AddIAMMemberRequest{
		UserId: user.GetUserId(),
		Roles:  iamRoles,
	})
	require.NoError(t, err)
	type args struct {
		ctx context.Context
		req *admin_pb.ListIAMMembersRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *admin_pb.ListIAMMembersResponse
		wantErr bool
	}{
		{
			name: "permission error",
			args: args{
				ctx: Tester.WithAuthorization(CTX, integration.OrgOwner),
				req: &admin_pb.ListIAMMembersRequest{
					Query: &object.ListQuery{},
					Queries: []*member.SearchQuery{{
						Query: &member.SearchQuery_UserIdQuery{
							UserIdQuery: &member.UserIDQuery{
								UserId: user.GetUserId(),
							},
						},
					}},
				},
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx: AdminCTX,
				req: &admin_pb.ListIAMMembersRequest{
					Query: &object.ListQuery{},
					Queries: []*member.SearchQuery{{
						Query: &member.SearchQuery_UserIdQuery{
							UserIdQuery: &member.UserIDQuery{
								UserId: user.GetUserId(),
							},
						},
					}},
				},
			},
			want: &admin_pb.ListIAMMembersResponse{
				Result: []*member.Member{{
					UserId: user.GetUserId(),
					Roles:  iamRoles,
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Client.ListIAMMembers(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			wantResult := tt.want.GetResult()
			gotResult := got.GetResult()

			require.Len(t, gotResult, len(wantResult))
			for i, want := range wantResult {
				assert.Equal(t, want.GetUserId(), gotResult[i].GetUserId())
				assert.ElementsMatch(t, want.GetRoles(), gotResult[i].GetRoles())
			}
		})
	}
}

func TestServer_AddIAMMember(t *testing.T) {
	user := Tester.CreateHumanUserVerified(AdminCTX, Tester.Organisation.ID, gofakeit.Email())
	type args struct {
		ctx context.Context
		req *admin_pb.AddIAMMemberRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *admin_pb.AddIAMMemberResponse
		wantErr bool
	}{
		{
			name: "permission error",
			args: args{
				ctx: Tester.WithAuthorization(CTX, integration.OrgOwner),
				req: &admin_pb.AddIAMMemberRequest{
					UserId: user.GetUserId(),
					Roles:  iamRoles,
				},
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx: AdminCTX,
				req: &admin_pb.AddIAMMemberRequest{
					UserId: user.GetUserId(),
					Roles:  iamRoles,
				},
			},
			want: &admin_pb.AddIAMMemberResponse{
				Details: &object.ObjectDetails{
					ResourceOwner: Tester.Instance.InstanceID(),
				},
			},
		},
		{
			name: "unknown roles error",
			args: args{
				ctx: AdminCTX,
				req: &admin_pb.AddIAMMemberRequest{
					UserId: user.GetUserId(),
					Roles:  []string{"FOO", "BAR"},
				},
			},
			wantErr: true,
		},
		{
			name: "org role error",
			args: args{
				ctx: AdminCTX,
				req: &admin_pb.AddIAMMemberRequest{
					UserId: user.GetUserId(),
					Roles:  []string{"ORG_OWNER"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Client.AddIAMMember(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			integration.AssertDetails(t, tt.want, got)
		})
	}
}

func TestServer_UpdateIAMMember(t *testing.T) {
	user := Tester.CreateHumanUserVerified(AdminCTX, Tester.Organisation.ID, gofakeit.Email())
	_, err := Client.AddIAMMember(AdminCTX, &admin_pb.AddIAMMemberRequest{
		UserId: user.GetUserId(),
		Roles:  []string{"IAM_OWNER"},
	})
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		req *admin_pb.UpdateIAMMemberRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *admin_pb.UpdateIAMMemberResponse
		wantErr bool
	}{
		{
			name: "permission error",
			args: args{
				ctx: Tester.WithAuthorization(CTX, integration.OrgOwner),
				req: &admin_pb.UpdateIAMMemberRequest{
					UserId: user.GetUserId(),
					Roles:  iamRoles,
				},
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx: AdminCTX,
				req: &admin_pb.UpdateIAMMemberRequest{
					UserId: user.GetUserId(),
					Roles:  iamRoles,
				},
			},
			want: &admin_pb.UpdateIAMMemberResponse{
				Details: &object.ObjectDetails{
					ResourceOwner: Tester.Instance.InstanceID(),
					ChangeDate:    timestamppb.Now(),
				},
			},
		},
		{
			name: "unknown roles error",
			args: args{
				ctx: AdminCTX,
				req: &admin_pb.UpdateIAMMemberRequest{
					UserId: user.GetUserId(),
					Roles:  []string{"FOO", "BAR"},
				},
			},
			wantErr: true,
		},
		{
			name: "org role error",
			args: args{
				ctx: AdminCTX,
				req: &admin_pb.UpdateIAMMemberRequest{
					UserId: user.GetUserId(),
					Roles:  []string{"ORG_OWNER"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Client.UpdateIAMMember(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			integration.AssertDetails(t, tt.want, got)
		})
	}
}

func TestServer_RemoveIAMMember(t *testing.T) {
	user := Tester.CreateHumanUserVerified(AdminCTX, Tester.Organisation.ID, gofakeit.Email())
	_, err := Client.AddIAMMember(AdminCTX, &admin_pb.AddIAMMemberRequest{
		UserId: user.GetUserId(),
		Roles:  []string{"IAM_OWNER"},
	})
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		req *admin_pb.RemoveIAMMemberRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *admin_pb.RemoveIAMMemberResponse
		wantErr bool
	}{
		{
			name: "permission error",
			args: args{
				ctx: Tester.WithAuthorization(CTX, integration.OrgOwner),
				req: &admin_pb.RemoveIAMMemberRequest{
					UserId: user.GetUserId(),
				},
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx: AdminCTX,
				req: &admin_pb.RemoveIAMMemberRequest{
					UserId: user.GetUserId(),
				},
			},
			want: &admin_pb.RemoveIAMMemberResponse{
				Details: &object.ObjectDetails{
					ResourceOwner: Tester.Instance.InstanceID(),
					ChangeDate:    timestamppb.Now(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Client.RemoveIAMMember(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			integration.AssertDetails(t, tt.want, got)
		})
	}
}
