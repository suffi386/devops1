import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { PolicyComponentAction } from '../policy-component-action.enum';
import { PasswordIamPolicyComponent } from './password-iam-policy.component';

const routes: Routes = [
    {
        path: '',
        component: PasswordIamPolicyComponent,
        data: {
            animation: 'DetailPage',
            action: PolicyComponentAction.MODIFY,
        },
    },
    {
        path: 'create',
        component: PasswordIamPolicyComponent,
        data: {
            animation: 'DetailPage',
            action: PolicyComponentAction.CREATE,
        },
    },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class PasswordIamPolicyRoutingModule { }
