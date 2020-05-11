import { SelectionModel } from '@angular/cdk/collections';
import { AfterViewInit, Component, EventEmitter, Input, OnInit, Output, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatTable } from '@angular/material/table';
import { tap } from 'rxjs/operators';
import { ProjectRole } from 'src/app/proto/generated/management_pb';
import { ProjectService } from 'src/app/services/project.service';
import { ToastService } from 'src/app/services/toast.service';

import { ProjectRolesDataSource } from './project-roles-datasource';


@Component({
    selector: 'app-project-roles',
    templateUrl: './project-roles.component.html',
    styleUrls: ['./project-roles.component.scss'],
})
export class ProjectRolesComponent implements AfterViewInit, OnInit {
    @Input() public projectId: string = '';
    @Input() public disabled: boolean = false;
    @Input() public actionsVisible: boolean = false;
    @ViewChild(MatPaginator) public paginator!: MatPaginator;
    @ViewChild(MatTable) public table!: MatTable<ProjectRole.AsObject>;
    public dataSource!: ProjectRolesDataSource;
    public selection: SelectionModel<ProjectRole.AsObject> = new SelectionModel<ProjectRole.AsObject>(true, []);
    @Output() public changedSelection: EventEmitter<Array<ProjectRole.AsObject>> = new EventEmitter();

    /** Columns displayed in the table. Columns IDs can be added, removed, or reordered. */
    public displayedColumns: string[] = ['select', 'name', 'displayname', 'group'];

    constructor(private projectService: ProjectService, private toast: ToastService) { }

    public ngOnInit(): void {
        this.dataSource = new ProjectRolesDataSource(this.projectService);
        this.dataSource.loadRoles(this.projectId, 0, 25, 'asc');

        this.selection.changed.subscribe(() => {
            this.changedSelection.emit(this.selection.selected);
        });
    }

    public ngAfterViewInit(): void {
        this.paginator.page
            .pipe(
                tap(() => this.loadRolesPage()),
            )
            .subscribe();
    }

    public selectAllOfGroup(group: string): void {
        const groupRoles: ProjectRole.AsObject[] = this.dataSource.rolesSubject.getValue()
            .filter(role => role.group === group);
        this.selection.select(...groupRoles);
    }

    private loadRolesPage(): void {
        this.dataSource.loadRoles(
            this.projectId,
            this.paginator.pageIndex,
            this.paginator.pageSize,
        );
    }

    public isAllSelected(): boolean {
        const numSelected = this.selection.selected.length;
        const numRows = this.dataSource.rolesSubject.value.length;
        return numSelected === numRows;
    }

    public masterToggle(): void {
        this.isAllSelected() ?
            this.selection.clear() :
            this.dataSource.rolesSubject.value.forEach((row: ProjectRole.AsObject) => this.selection.select(row));
    }

    public deleteSelectedRoles(): Promise<any> {
        const oldState = this.dataSource.rolesSubject.value;
        const indexes = this.selection.selected.map(sel => {
            return oldState.findIndex(iter => iter.name === sel.name);
        });

        return Promise.all(this.selection.selected.map(role => {
            return this.projectService.RemoveProjectRole(role.projectId, role.name);
        })).then(() => {
            this.toast.showInfo('Deleted');
            indexes.forEach(index => {
                if (index > -1) {
                    oldState.splice(index, 1);
                    this.dataSource.rolesSubject.next(this.dataSource.rolesSubject.value);
                }
            });
            this.selection.clear();
        }).catch(error => {
            this.toast.showError(error?.message || 'Error');
        });
    }

    public removeRole(role: ProjectRole.AsObject, index: number): void {
        this.projectService
            .RemoveProjectRole(role.projectId, role.name)
            .then(() => {
                this.toast.showInfo('Role removed');
                this.dataSource.rolesSubject.value.splice(index, 1);
                this.dataSource.rolesSubject.next(this.dataSource.rolesSubject.value);
            })
            .catch(data => {
                this.toast.showError(data.message);
            });
    }
}
