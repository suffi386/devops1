import { COMMA, ENTER, SPACE } from '@angular/cdk/keycodes';
import { Location } from '@angular/common';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatChipInputEvent } from '@angular/material/chips';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import {
    Application,
    OIDCApplicationCreate,
    OIDCApplicationType,
    OIDCAuthMethodType,
    OIDCGrantType,
    OIDCResponseType,
} from 'src/app/proto/generated/management_pb';
import { ProjectService } from 'src/app/services/project.service';
import { ToastService } from 'src/app/services/toast.service';

import { AppSecretDialogComponent } from '../app-secret-dialog/app-secret-dialog.component';

@Component({
    selector: 'app-app-create',
    templateUrl: './app-create.component.html',
    styleUrls: ['./app-create.component.scss'],
})
export class AppCreateComponent implements OnInit, OnDestroy {
    private subscription?: Subscription;
    public projectId: string = '';
    public loading: boolean = false;
    public oidcApp: OIDCApplicationCreate.AsObject = new OIDCApplicationCreate().toObject();
    public oidcResponseTypes: { type: OIDCResponseType, checked: boolean; }[] = [
        { type: OIDCResponseType.OIDCRESPONSETYPE_CODE, checked: false },
        { type: OIDCResponseType.OIDCRESPONSETYPE_ID_TOKEN, checked: false },
        { type: OIDCResponseType.OIDCRESPONSETYPE_ID_TOKEN_TOKEN, checked: false },
    ];
    public oidcGrantTypes: {
        type: OIDCGrantType,
        checked: boolean,
        disabled: boolean,
    }[] = [
            { type: OIDCGrantType.OIDCGRANTTYPE_AUTHORIZATION_CODE, checked: true, disabled: false },
            { type: OIDCGrantType.OIDCGRANTTYPE_IMPLICIT, checked: false, disabled: true },
            { type: OIDCGrantType.OIDCGRANTTYPE_REFRESH_TOKEN, checked: false, disabled: true },
        ];
    public oidcAppTypes: OIDCApplicationType[] = [
        OIDCApplicationType.OIDCAPPLICATIONTYPE_WEB,
        OIDCApplicationType.OIDCAPPLICATIONTYPE_USER_AGENT,
        OIDCApplicationType.OIDCAPPLICATIONTYPE_NATIVE,
    ];
    public oidcAuthMethodType: { type: OIDCAuthMethodType, checked: boolean, disabled: boolean; }[] = [
        { type: OIDCAuthMethodType.OIDCAUTHMETHODTYPE_BASIC, checked: false, disabled: false },
        { type: OIDCAuthMethodType.OIDCAUTHMETHODTYPE_NONE, checked: false, disabled: false },
        { type: OIDCAuthMethodType.OIDCAUTHMETHODTYPE_POST, checked: false, disabled: false },
    ];

    firstFormGroup!: FormGroup;
    secondFormGroup!: FormGroup;
    thirdFormGroup!: FormGroup;

    public OIDCApplicationType: any = OIDCApplicationType;
    public OIDCGrantType: any = OIDCGrantType;
    public OIDCAuthMethodType: any = OIDCAuthMethodType;

    public postLogoutRedirectUrisList: string[] = [];

    public addOnBlur: boolean = true;
    public readonly separatorKeysCodes: number[] = [ENTER, COMMA, SPACE];

    constructor(
        private router: Router,
        private route: ActivatedRoute,
        private toast: ToastService,
        private dialog: MatDialog,
        private projectService: ProjectService,
        private fb: FormBuilder,
        private _location: Location,
    ) {
        this.firstFormGroup = this.fb.group({
            name: ['', [Validators.required]],
            applicationType: ['', [Validators.required]],
        });
        this.secondFormGroup = this.fb.group({
            authMethodType: [OIDCAuthMethodType.OIDCAUTHMETHODTYPE_BASIC, [Validators.required]],
        });

        this.oidcApp.authMethodType = OIDCAuthMethodType.OIDCAUTHMETHODTYPE_BASIC;

        this.firstFormGroup.valueChanges.subscribe(value => {
            if (this.applicationType?.value === OIDCApplicationType.OIDCAPPLICATIONTYPE_NATIVE) {
                this.oidcResponseTypes[0].checked = true;
                // this.oidcApp.responseTypesList = [this.oidcResponseTypes[0].type];

                this.authMethodType?.setValue(OIDCAuthMethodType.OIDCAUTHMETHODTYPE_NONE);
                this.oidcAuthMethodType[1].disabled = true;
            }
            if (this.applicationType?.value === OIDCApplicationType.OIDCAPPLICATIONTYPE_WEB) {
                this.oidcResponseTypes[0].checked = true;
                this.oidcApp.responseTypesList = [this.oidcResponseTypes[0].type];

                this.authMethodType?.setValue(OIDCAuthMethodType.OIDCAUTHMETHODTYPE_NONE);
                this.oidcAuthMethodType[0].disabled = true;
                this.oidcAuthMethodType[2].disabled = true;
            }
            if (this.applicationType?.value === OIDCApplicationType.OIDCAPPLICATIONTYPE_USER_AGENT) {
                this.oidcResponseTypes[0].checked = true;
                this.oidcApp.responseTypesList = [this.oidcResponseTypes[0].type];

                this.authMethodType?.setValue(OIDCAuthMethodType.OIDCAUTHMETHODTYPE_NONE);
                this.oidcAuthMethodType[0].disabled = true;
                this.oidcAuthMethodType[2].disabled = true;

                this.oidcGrantTypes[1].disabled = false;
            }

            this.changeResponseType();
            this.changeGrant();

            this.oidcApp.name = this.name?.value;
            this.oidcApp.applicationType = this.applicationType?.value;
        });

        this.secondFormGroup.valueChanges.subscribe(value => {
            this.oidcApp.authMethodType = this.authMethodType?.value;
        });
    }

    public ngOnInit(): void {
        this.subscription = this.route.params.subscribe(params => this.getData(params));
    }

    public ngOnDestroy(): void {
        this.subscription?.unsubscribe();
    }

    private async getData({ projectid }: Params): Promise<void> {
        this.projectId = projectid;
        this.oidcApp.projectId = projectid;
    }

    public close(): void {
        this._location.back();
    }

    public saveOIDCApp(): void {
        this.loading = true;
        this.projectService
            .CreateOIDCApp(this.oidcApp)
            .then((data: Application) => {
                this.loading = false;
                this.showSavedDialog(data.toObject());
            })
            .catch(error => {
                this.loading = false;
                this.toast.showError(error);
            });
    }

    public showSavedDialog(app: Application.AsObject): void {
        if (app.oidcConfig !== undefined) {
            const dialogRef = this.dialog.open(AppSecretDialogComponent, {
                data: app.oidcConfig,
            });

            dialogRef.afterClosed().subscribe(result => {
                this.router.navigate(['projects', this.projectId, 'apps', app.id]);
            });
        } else {
            this.router.navigate(['projects', this.projectId, 'apps', app.id]);
        }
    }

    public addUri(event: MatChipInputEvent, target: string): void {
        const input = event.input;
        const value = event.value.trim();

        if (value !== '') {
            if (target === 'REDIRECT') {
                this.oidcApp.redirectUrisList.push(value);
            } else if (target === 'POSTREDIRECT') {
                this.oidcApp.postLogoutRedirectUrisList.push(value);
            }
        }

        if (input) {
            input.value = '';
        }
    }

    public removeUri(uri: string, target: string): void {
        if (target === 'REDIRECT') {
            const index = this.oidcApp.redirectUrisList.indexOf(uri);

            if (index !== undefined && index >= 0) {
                this.oidcApp.redirectUrisList.splice(index, 1);
            }
        } else if (target === 'POSTREDIRECT') {
            const index = this.oidcApp.postLogoutRedirectUrisList.indexOf(uri);

            if (index !== undefined && index >= 0) {
                this.oidcApp.postLogoutRedirectUrisList.splice(index, 1);
            }
        }
    }

    changeGrant(): void {
        this.oidcApp.grantTypesList = this.oidcGrantTypes.filter(gt => gt.checked).map(gt => gt.type);
    }

    changeResponseType(): void {
        this.oidcApp.responseTypesList = this.oidcResponseTypes.filter(gt => gt.checked).map(gt => gt.type);
    }

    get name(): AbstractControl | null {
        return this.firstFormGroup.get('name');
    }

    get applicationType(): AbstractControl | null {
        return this.firstFormGroup.get('applicationType');
    }

    // get grantTypesList(): AbstractControl | null {
    //     return this.secondFormGroup.get('grantTypesList');
    // }

    // get getCheckedOidcGrantTypes(): OIDCGrantType[] {
    //     return this.oidcGrantTypes.filter(gt => gt.checked).map(gt => gt.type);
    // }

    public grantTypeChecked(type: OIDCGrantType): boolean {
        return this.oidcGrantTypes.filter(gt => gt.checked).map(gt => gt.type).findIndex(t => t === type) > -1;
    }

    get responseTypesList(): AbstractControl | null {
        return this.secondFormGroup.get('responseTypesList');
    }

    // get applicationType(): AbstractControl | null {
    //     return this.form.get('applicationType');
    // }

    get authMethodType(): AbstractControl | null {
        return this.secondFormGroup.get('authMethodType');
    }
}

