import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { OrgDomainValidationResponse, OrgDomainValidationType, OrgDomainView } from 'src/app/proto/generated/management_pb';
import { OrgService } from 'src/app/services/org.service';

@Component({
    selector: 'app-domain-verification',
    templateUrl: './domain-verification.component.html',
    styleUrls: ['./domain-verification.component.scss'],
})
export class DomainVerificationComponent {
    public domain!: OrgDomainView.AsObject;

    public OrgDomainValidationType: any = OrgDomainValidationType;

    public http!: OrgDomainValidationResponse.AsObject;
    public dns!: OrgDomainValidationResponse.AsObject;
    public copied: string = '';
    // @ViewChild('link') public link!: any;
    constructor(
        public dialogRef: MatDialogRef<DomainVerificationComponent>,
        @Inject(MAT_DIALOG_DATA) public data: any,
        private orgService: OrgService,
    ) {
        this.domain = data.domain;

        this.loadTokens().then(() => {
            const blob = new Blob([this.http.token], { type: 'octet/stream' });
            const link: any = document.getElementById('link');
            link.href = window.URL.createObjectURL(blob);
        });
    }

    async loadTokens(): Promise<void> {
        this.http = (await this.orgService.GenerateMyOrgDomainValidation(
            this.domain.domain,
            OrgDomainValidationType.ORGDOMAINVALIDATIONTYPE_HTTP)).toObject();
        this.dns = (await this.orgService.GenerateMyOrgDomainValidation(
            this.domain.domain,
            OrgDomainValidationType.ORGDOMAINVALIDATIONTYPE_DNS)).toObject();
    }

    public closeDialog(): void {
        this.dialogRef.close(false);
    }

    public validate(): void {
        this.orgService.ValidateMyOrgDomain(this.domain.domain).then(() => {
            this.dialogRef.close(false);
        });
    }
}
