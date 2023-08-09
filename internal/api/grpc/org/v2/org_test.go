package org

import (
	"testing"
	"time"

	"github.com/muhlemmer/gu"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zitadel/zitadel/internal/command"
	"github.com/zitadel/zitadel/internal/domain"
	caos_errs "github.com/zitadel/zitadel/internal/errors"
	object "github.com/zitadel/zitadel/pkg/grpc/object/v2alpha"
	org "github.com/zitadel/zitadel/pkg/grpc/org/v2beta"
	user "github.com/zitadel/zitadel/pkg/grpc/user/v2alpha"
)

func Test_addOrganisationRequestToCommand(t *testing.T) {
	type args struct {
		request *org.AddOrganisationRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *command.OrgSetup
		wantErr error
	}{
		{
			name: "nil user",
			args: args{
				request: &org.AddOrganisationRequest{
					Name: "name",
					Admins: []*org.AddOrganisationRequest_Admin{
						{},
					},
				},
			},
			wantErr: caos_errs.ThrowUnimplementedf(nil, "ORGv2-SD2r1", "userType oneOf %T in method AddOrganisation not implemented", nil),
		},
		{
			name: "user ID",
			args: args{
				request: &org.AddOrganisationRequest{
					Name: "name",
					Admins: []*org.AddOrganisationRequest_Admin{
						{
							UserType: &org.AddOrganisationRequest_Admin_UserId{
								UserId: "userID",
							},
							Roles: nil,
						},
					},
				},
			},
			want: &command.OrgSetup{
				Name:         "name",
				CustomDomain: "",
				Admins: []*command.OrgSetupAdmin{
					{
						ID: "userID",
					},
				},
			},
		},
		{
			name: "human user",
			args: args{
				request: &org.AddOrganisationRequest{
					Name: "name",
					Admins: []*org.AddOrganisationRequest_Admin{
						{
							UserType: &org.AddOrganisationRequest_Admin_Human{
								Human: &user.AddHumanUserRequest{
									Profile: &user.SetHumanProfile{
										FirstName: "firstname",
										LastName:  "lastname",
									},
									Email: &user.SetHumanEmail{
										Email: "email@test.com",
									},
								},
							},
							Roles: nil,
						},
					},
				},
			},
			want: &command.OrgSetup{
				Name:         "name",
				CustomDomain: "",
				Admins: []*command.OrgSetupAdmin{
					{
						Human: &command.AddHuman{
							Username:  "email@test.com",
							FirstName: "firstname",
							LastName:  "lastname",
							Email: command.Email{
								Address: "email@test.com",
							},
							Metadata: make([]*command.AddMetadataEntry, 0),
							Links:    make([]*command.AddLink, 0),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addOrganisationRequestToCommand(tt.args.request)
			require.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_createdOrganisationToPb(t *testing.T) {
	now := time.Now()
	type args struct {
		createdOrg *command.CreatedOrg
	}
	tests := []struct {
		name    string
		args    args
		want    *org.AddOrganisationResponse
		wantErr error
	}{
		{
			name: "human user with phone and email code",
			args: args{
				createdOrg: &command.CreatedOrg{
					ObjectDetails: &domain.ObjectDetails{
						Sequence:      1,
						EventDate:     now,
						ResourceOwner: "orgID",
					},
					CreatedAdmins: []*command.CreatedOrgAdmin{
						{
							ID:        "id",
							EmailCode: gu.Ptr("emailCode"),
							PhoneCode: gu.Ptr("phoneCode"),
						},
					},
				},
			},
			want: &org.AddOrganisationResponse{
				Details: &object.Details{
					Sequence:      1,
					ChangeDate:    timestamppb.New(now),
					ResourceOwner: "orgID",
				},
				OrganisationId: "orgID",
				CreatedAdmins: []*org.AddOrganisationResponse_CreatedAdmin{
					{
						UserId:    "id",
						EmailCode: gu.Ptr("emailCode"),
						PhoneCode: gu.Ptr("phoneCode"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createdOrganisationToPb(tt.args.createdOrg)
			require.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
