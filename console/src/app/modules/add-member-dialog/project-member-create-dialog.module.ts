import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { TranslateModule } from '@ngx-translate/core';
import { SearchUserAutocompleteModule } from 'src/app/modules/search-user-autocomplete/search-user-autocomplete.module';

import {
    OrgMemberRolesAutocompleteModule,
} from '../../pages/orgs/org-member-roles-autocomplete/org-member-roles-autocomplete.module';
import { SearchRolesAutocompleteModule } from '../search-roles-autocomplete/search-roles-autocomplete.module';
import { ProjectMemberCreateDialogComponent } from './project-member-create-dialog.component';

@NgModule({
    declarations: [ProjectMemberCreateDialogComponent],
    imports: [
        CommonModule,
        MatDialogModule,
        MatButtonModule,
        TranslateModule,
        MatFormFieldModule,
        MatSelectModule,
        FormsModule,
        SearchUserAutocompleteModule,
        SearchRolesAutocompleteModule,
        OrgMemberRolesAutocompleteModule,
    ],
    entryComponents: [
        ProjectMemberCreateDialogComponent,
    ],
})
export class ProjectMemberCreateDialogModule { }
