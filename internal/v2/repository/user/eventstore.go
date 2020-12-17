package user

import (
	"github.com/caos/zitadel/internal/eventstore/v2"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(UserV1AddedType, HumanAddedEventMapper).
		RegisterFilterEventMapper(UserV1RegisteredType, HumanRegisteredEventMapper).
		RegisterFilterEventMapper(UserV1InitialCodeAddedType, HumanInitialCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1InitialCodeSentType, HumanInitialCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1InitializedCheckSucceededType, HumanInitializedCheckSucceededEventMapper).
		RegisterFilterEventMapper(UserV1InitializedCheckFailedType, HumanInitializedCheckFailedEventMapper).
		RegisterFilterEventMapper(UserV1SignedOutType, HumanSignedOutEventMapper).
		RegisterFilterEventMapper(UserV1PasswordChangedType, HumanPasswordChangedEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCodeAddedType, HumanPasswordCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCodeSentType, HumanPasswordCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCheckSucceededType, HumanPasswordCheckSucceededEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCheckFailedType, HumanPasswordCheckFailedEventMapper).
		RegisterFilterEventMapper(UserV1EmailChangedType, HumanEmailChangedEventMapper).
		RegisterFilterEventMapper(UserV1EmailVerifiedType, HumanEmailVerifiedEventMapper).
		RegisterFilterEventMapper(UserV1EmailVerificationFailedType, HumanEmailVerificationFailedEventMapper).
		RegisterFilterEventMapper(UserV1EmailCodeAddedType, HumanEmailCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1EmailCodeSentType, HumanEmailCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1PhoneChangedType, HumanPhoneChangedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneRemovedType, HumanPhoneRemovedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneVerifiedType, HumanPhoneVerifiedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneVerificationFailedType, HumanPhoneVerificationFailedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneCodeAddedType, HumanPhoneCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneCodeSentType, HumanPhoneCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1ProfileChangedType, HumanProfileChangedEventMapper).
		RegisterFilterEventMapper(UserV1AddressChangedType, HumanAddressChangedEventMapper).
		RegisterFilterEventMapper(UserV1MFAInitSkippedType, HumanMFAInitSkippedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPAddedType, HumanOTPAddedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPVerifiedType, HumanOTPVerifiedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPRemovedType, HumanOTPRemovedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPCheckSucceededType, HumanOTPCheckSucceededEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPCheckFailedType, HumanOTPCheckFailedEventMapper).
		RegisterFilterEventMapper(UserLockedType, UserLockedEventMapper).
		RegisterFilterEventMapper(UserUnlockedType, UserLockedEventMapper).
		RegisterFilterEventMapper(UserDeactivatedType, UserDeactivatedEventMapper).
		RegisterFilterEventMapper(UserReactivatedType, UserReactivatedEventMapper).
		RegisterFilterEventMapper(UserRemovedType, UserRemovedEventMapper).
		RegisterFilterEventMapper(UserTokenAddedType, UserTokenAddedEventMapper).
		RegisterFilterEventMapper(UserDomainClaimedType, DomainClaimedEventMapper).
		RegisterFilterEventMapper(UserDomainClaimedSentType, DomainClaimedEventMapper).
		RegisterFilterEventMapper(UserUserNameChangedType, UsernameChangedEventMapper).
		RegisterFilterEventMapper(HumanAddedType, HumanAddedEventMapper).
		RegisterFilterEventMapper(HumanRegisteredType, HumanRegisteredEventMapper).
		RegisterFilterEventMapper(HumanInitialCodeAddedType, HumanInitialCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanInitialCodeSentType, HumanInitialCodeSentEventMapper).
		RegisterFilterEventMapper(HumanInitializedCheckSucceededType, HumanInitializedCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanInitializedCheckFailedType, HumanInitializedCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanSignedOutType, HumanSignedOutEventMapper).
		RegisterFilterEventMapper(HumanPasswordChangedType, HumanPasswordChangedEventMapper).
		RegisterFilterEventMapper(HumanPasswordCodeAddedType, HumanPasswordCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanPasswordCodeSentType, HumanPasswordCodeSentEventMapper).
		RegisterFilterEventMapper(HumanPasswordCheckSucceededType, HumanPasswordCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanPasswordCheckFailedType, HumanPasswordCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanExternalIDPAddedType, HumanExternalIDPAddedEventMapper).
		RegisterFilterEventMapper(HumanExternalIDPRemovedType, HumanExternalIDPRemovedEventMapper).
		RegisterFilterEventMapper(HumanExternalIDPCascadeRemovedType, HumanExternalIDPCascadeRemovedEventMapper).
		RegisterFilterEventMapper(HumanExternalLoginCheckSucceededType, HumanExternalIDPCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanEmailChangedType, HumanEmailChangedEventMapper).
		RegisterFilterEventMapper(HumanEmailVerifiedType, HumanEmailVerifiedEventMapper).
		RegisterFilterEventMapper(HumanEmailVerificationFailedType, HumanEmailVerificationFailedEventMapper).
		RegisterFilterEventMapper(HumanEmailCodeAddedType, HumanEmailCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanEmailCodeSentType, HumanEmailCodeSentEventMapper).
		RegisterFilterEventMapper(HumanPhoneChangedType, HumanPhoneChangedEventMapper).
		RegisterFilterEventMapper(HumanPhoneRemovedType, HumanPhoneRemovedEventMapper).
		RegisterFilterEventMapper(HumanPhoneVerifiedType, HumanPhoneVerifiedEventMapper).
		RegisterFilterEventMapper(HumanPhoneVerificationFailedType, HumanPhoneVerificationFailedEventMapper).
		RegisterFilterEventMapper(HumanPhoneCodeAddedType, HumanPhoneCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanPhoneCodeSentType, HumanPhoneCodeSentEventMapper).
		RegisterFilterEventMapper(HumanProfileChangedType, HumanProfileChangedEventMapper).
		RegisterFilterEventMapper(HumanAddressChangedType, HumanAddressChangedEventMapper).
		RegisterFilterEventMapper(HumanMFAInitSkippedType, HumanMFAInitSkippedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPAddedType, HumanOTPAddedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPVerifiedType, HumanOTPVerifiedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPRemovedType, HumanOTPRemovedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPCheckSucceededType, HumanOTPCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPCheckFailedType, HumanOTPCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenAddedType, WebAuthNAddedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenVerifiedType, HumanWebAuthNVerifiedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenSignCountChangedType, HumanWebAuthNSignCountChangedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenRemovedType, HumanWebAuthNRemovedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenBeginLoginType, HumanWebAuthNBeginLoginEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenCheckSucceededType, HumanWebAuthNCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenCheckFailedType, HumanWebAuthNCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenAddedType, WebAuthNAddedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenVerifiedType, HumanWebAuthNVerifiedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenSignCountChangedType, HumanWebAuthNSignCountChangedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenRemovedType, HumanWebAuthNRemovedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenBeginLoginType, HumanWebAuthNBeginLoginEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenCheckSucceededType, HumanWebAuthNCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenCheckFailedType, HumanWebAuthNCheckFailedEventMapper).
		RegisterFilterEventMapper(MachineAddedEventType, MachineAddedEventMapper).
		RegisterFilterEventMapper(MachineChangedEventType, MachineChangedEventMapper).
		RegisterFilterEventMapper(MachineKeyAddedEventType, MachineKeyAddedEventMapper).
		RegisterFilterEventMapper(MachineKeyRemovedEventType, MachineKeyRemovedEventMapper)
}
