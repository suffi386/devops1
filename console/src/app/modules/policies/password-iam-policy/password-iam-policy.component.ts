import { Component, Injector, Input, OnDestroy, Type } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import { OrgIamPolicyView as AdminOrgIamPolicyView } from 'src/app/proto/generated/admin_pb';
import { Org } from 'src/app/proto/generated/auth_pb';
import { OrgIamPolicyView as MgmtOrgIamPolicyView } from 'src/app/proto/generated/management_pb';
import { AdminService } from 'src/app/services/admin.service';
import { ManagementService } from 'src/app/services/mgmt.service';
import { StorageService } from 'src/app/services/storage.service';
import { ToastService } from 'src/app/services/toast.service';

import { PolicyComponentServiceType } from '../policy-component-types.enum';

@Component({
    selector: 'app-password-iam-policy',
    templateUrl: './password-iam-policy.component.html',
    styleUrls: ['./password-iam-policy.component.scss'],
})
export class PasswordIamPolicyComponent implements OnDestroy {
    @Input() service!: AdminService;
    private managementService!: ManagementService;
    public serviceType: PolicyComponentServiceType = PolicyComponentServiceType.MGMT;

    public iamData!: AdminOrgIamPolicyView.AsObject | MgmtOrgIamPolicyView.AsObject;

    private sub: Subscription = new Subscription();
    private org!: Org.AsObject;

    public PolicyComponentServiceType: any = PolicyComponentServiceType;

    constructor(
        private route: ActivatedRoute,
        private toast: ToastService,
        private sessionStorage: StorageService,
        private injector: Injector,
        private adminService: AdminService,
    ) {
        const temporg = this.sessionStorage.getItem('organization') as Org.AsObject;
        if (temporg) {
            this.org = temporg;
        }
        this.sub = this.route.data.pipe(switchMap(data => {
            this.serviceType = data.serviceType;
            console.log(data.serviceType);
            if (this.serviceType === PolicyComponentServiceType.MGMT) {
                this.managementService = this.injector.get(ManagementService as Type<ManagementService>);
            }
            return this.route.params;
        })).subscribe(_ => {
            this.getData().then(data => {
                if (data) {
                    console.log(data.toObject());
                    this.iamData = data.toObject();
                }
            });
        });
    }

    public ngOnDestroy(): void {
        this.sub.unsubscribe();
    }

    private async getData(): Promise<AdminOrgIamPolicyView | MgmtOrgIamPolicyView | undefined> {
        switch (this.serviceType) {
            case PolicyComponentServiceType.MGMT:
                return this.managementService.GetMyOrgIamPolicy();
            case PolicyComponentServiceType.ADMIN:
                if (this.org?.id) {
                    return this.adminService.GetOrgIamPolicy(this.org.id);
                }
                break;
        }
    }

    public savePolicy(): void {
        switch (this.serviceType) {
            case PolicyComponentServiceType.MGMT:
                if ((this.iamData as MgmtOrgIamPolicyView.AsObject).pb_default) {
                    this.adminService.CreateOrgIamPolicy(
                        this.org.id,
                        this.iamData.userLoginMustBeDomain,
                    ).then(() => {
                        this.toast.showInfo('ORG.POLICY.TOAST.SET', true);
                    }).catch(error => {
                        this.toast.showError(error);
                    });
                    break;
                } else {
                    this.adminService.UpdateOrgIamPolicy(
                        this.org.id,
                        this.iamData.userLoginMustBeDomain,
                    ).then(() => {
                        this.toast.showInfo('ORG.POLICY.TOAST.SET', true);
                    }).catch(error => {
                        this.toast.showError(error);
                    });
                    break;
                }
            case PolicyComponentServiceType.ADMIN:
                // update Default org iam policy?
                this.adminService.UpdateOrgIamPolicy(
                    this.org.id,
                    this.iamData.userLoginMustBeDomain,
                ).then(() => {
                    this.toast.showInfo('ORG.POLICY.TOAST.SET', true);
                }).catch(error => {
                    this.toast.showError(error);
                });
                break;
        }
    }

    // public removePolicy(): void {
    //     if (this.service instanceof ManagementService) {
    //         this.service.removeIamPolicy().then(() => {
    //             this.toast.showInfo('Successfully deleted');
    //         }).catch(error => {
    //             this.toast.showError(error);
    //         });
    //     }
    // }

    public get isDefault(): boolean {
        if (this.iamData && this.serviceType === PolicyComponentServiceType.MGMT) {
            return (this.iamData as MgmtOrgIamPolicyView.AsObject).pb_default;
        } else {
            return false;
        }
    }
}
