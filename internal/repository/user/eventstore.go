package user

import (
	"github.com/zitadel/zitadel/internal/eventstore"
)

func init() {
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1AddedType, HumanAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1RegisteredType, HumanRegisteredEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1InitialCodeAddedType, HumanInitialCodeAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1InitialCodeSentType, HumanInitialCodeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1InitializedCheckSucceededType, HumanInitializedCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1InitializedCheckFailedType, HumanInitializedCheckFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1SignedOutType, HumanSignedOutEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PasswordChangedType, HumanPasswordChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PasswordCodeAddedType, HumanPasswordCodeAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PasswordCodeSentType, HumanPasswordCodeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PasswordCheckSucceededType, HumanPasswordCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PasswordCheckFailedType, HumanPasswordCheckFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1EmailChangedType, HumanEmailChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1EmailVerifiedType, HumanEmailVerifiedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1EmailVerificationFailedType, HumanEmailVerificationFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1EmailCodeAddedType, HumanEmailCodeAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1EmailCodeSentType, HumanEmailCodeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PhoneChangedType, HumanPhoneChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PhoneRemovedType, HumanPhoneRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PhoneVerifiedType, HumanPhoneVerifiedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PhoneVerificationFailedType, HumanPhoneVerificationFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PhoneCodeAddedType, HumanPhoneCodeAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1PhoneCodeSentType, HumanPhoneCodeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1ProfileChangedType, HumanProfileChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1AddressChangedType, HumanAddressChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1MFAInitSkippedType, HumanMFAInitSkippedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1MFAOTPAddedType, HumanOTPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1MFAOTPVerifiedType, HumanOTPVerifiedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1MFAOTPRemovedType, HumanOTPRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1MFAOTPCheckSucceededType, HumanOTPCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserV1MFAOTPCheckFailedType, HumanOTPCheckFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserLockedType, UserLockedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserUnlockedType, UserUnlockedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserDeactivatedType, UserDeactivatedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserReactivatedType, UserReactivatedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserRemovedType, UserRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserTokenAddedType, UserTokenAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserTokenRemovedType, UserTokenRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserDomainClaimedType, DomainClaimedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserDomainClaimedSentType, DomainClaimedSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserUserNameChangedType, UsernameChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MetadataSetType, MetadataSetEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MetadataRemovedType, MetadataRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MetadataRemovedAllType, MetadataRemovedAllEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanAddedType, HumanAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanRegisteredType, HumanRegisteredEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanInitialCodeAddedType, HumanInitialCodeAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanInitialCodeSentType, HumanInitialCodeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanInitializedCheckSucceededType, HumanInitializedCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanInitializedCheckFailedType, HumanInitializedCheckFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanSignedOutType, HumanSignedOutEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordChangedType, HumanPasswordChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordCodeAddedType, HumanPasswordCodeAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordCodeSentType, HumanPasswordCodeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordChangeSentType, HumanPasswordChangeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordCheckSucceededType, HumanPasswordCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordCheckFailedType, HumanPasswordCheckFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordHashUpdatedType, eventstore.GenericEventMapper[HumanPasswordHashUpdatedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, UserIDPLinkAddedType, UserIDPLinkAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserIDPLinkRemovedType, UserIDPLinkRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserIDPLinkCascadeRemovedType, UserIDPLinkCascadeRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserIDPLoginCheckSucceededType, UserIDPCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, UserIDPExternalIDMigratedType, eventstore.GenericEventMapper[UserIDPExternalIDMigratedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, UserIDPExternalUsernameChangedType, eventstore.GenericEventMapper[UserIDPExternalUsernameEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanEmailChangedType, HumanEmailChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanEmailVerifiedType, HumanEmailVerifiedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanEmailVerificationFailedType, HumanEmailVerificationFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanEmailCodeAddedType, HumanEmailCodeAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanEmailCodeSentType, HumanEmailCodeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPhoneChangedType, HumanPhoneChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPhoneRemovedType, HumanPhoneRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPhoneVerifiedType, HumanPhoneVerifiedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPhoneVerificationFailedType, HumanPhoneVerificationFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPhoneCodeAddedType, HumanPhoneCodeAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPhoneCodeSentType, HumanPhoneCodeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanProfileChangedType, HumanProfileChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanAvatarAddedType, HumanAvatarAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanAvatarRemovedType, HumanAvatarRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanAddressChangedType, HumanAddressChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanMFAInitSkippedType, HumanMFAInitSkippedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanMFAOTPAddedType, HumanOTPAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanMFAOTPVerifiedType, HumanOTPVerifiedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanMFAOTPRemovedType, HumanOTPRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanMFAOTPCheckSucceededType, HumanOTPCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanMFAOTPCheckFailedType, HumanOTPCheckFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPSMSAddedType, eventstore.GenericEventMapper[HumanOTPSMSAddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPSMSRemovedType, eventstore.GenericEventMapper[HumanOTPSMSRemovedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPSMSCodeAddedType, eventstore.GenericEventMapper[HumanOTPSMSCodeAddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPSMSCodeSentType, eventstore.GenericEventMapper[HumanOTPSMSCodeSentEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPSMSCheckSucceededType, eventstore.GenericEventMapper[HumanOTPSMSCheckSucceededEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPSMSCheckFailedType, eventstore.GenericEventMapper[HumanOTPSMSCheckFailedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPEmailAddedType, eventstore.GenericEventMapper[HumanOTPEmailAddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPEmailRemovedType, eventstore.GenericEventMapper[HumanOTPEmailRemovedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPEmailCodeAddedType, eventstore.GenericEventMapper[HumanOTPEmailCodeAddedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPEmailCodeSentType, eventstore.GenericEventMapper[HumanOTPEmailCodeSentEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPEmailCheckSucceededType, eventstore.GenericEventMapper[HumanOTPEmailCheckSucceededEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanOTPEmailCheckFailedType, eventstore.GenericEventMapper[HumanOTPEmailCheckFailedEvent])
	eventstore.RegisterFilterEventMapper(AggregateType, HumanU2FTokenAddedType, HumanU2FAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanU2FTokenVerifiedType, HumanU2FVerifiedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanU2FTokenSignCountChangedType, HumanU2FSignCountChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanU2FTokenRemovedType, HumanU2FRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanU2FTokenBeginLoginType, HumanU2FBeginLoginEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanU2FTokenCheckSucceededType, HumanU2FCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanU2FTokenCheckFailedType, HumanU2FCheckFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessTokenAddedType, HumanPasswordlessAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessTokenVerifiedType, HumanPasswordlessVerifiedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessTokenSignCountChangedType, HumanPasswordlessSignCountChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessTokenRemovedType, HumanPasswordlessRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessTokenBeginLoginType, HumanPasswordlessBeginLoginEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessTokenCheckSucceededType, HumanPasswordlessCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessTokenCheckFailedType, HumanPasswordlessCheckFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessInitCodeAddedType, HumanPasswordlessInitCodeAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessInitCodeRequestedType, HumanPasswordlessInitCodeRequestedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessInitCodeSentType, HumanPasswordlessInitCodeSentEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessInitCodeCheckFailedType, HumanPasswordlessInitCodeCodeCheckFailedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanPasswordlessInitCodeCheckSucceededType, HumanPasswordlessInitCodeCodeCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanRefreshTokenAddedType, HumanRefreshTokenAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanRefreshTokenRenewedType, HumanRefreshTokenRenewedEventEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, HumanRefreshTokenRemovedType, HumanRefreshTokenRemovedEventEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MachineAddedEventType, MachineAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MachineChangedEventType, MachineChangedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MachineKeyAddedEventType, MachineKeyAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MachineKeyRemovedEventType, MachineKeyRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, PersonalAccessTokenAddedType, PersonalAccessTokenAddedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, PersonalAccessTokenRemovedType, PersonalAccessTokenRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MachineSecretSetType, MachineSecretSetEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MachineSecretRemovedType, MachineSecretRemovedEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MachineSecretCheckSucceededType, MachineSecretCheckSucceededEventMapper)
	eventstore.RegisterFilterEventMapper(AggregateType, MachineSecretCheckFailedType, MachineSecretCheckFailedEventMapper)
}
