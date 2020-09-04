import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {AuthComponent} from './auth/auth.component';
import {LoginComponent} from './auth/login/login.component';
import {RegistrationComponent} from './auth/registration/registration.component';
import {AdminComponent} from './admin/admin.component';
import {NotFoundComponent} from './not-found/not-found.component';
import {SessionGuard} from './session.guard';
import {ObjectsListComponent} from './admin/objects-list/objects-list.component';
import {RunTestsComponent} from './admin/run-tests/run-tests.component';
import {CreateObjectComponent} from './admin/create-object/create-object.component';
import {EditObjectComponent} from './admin/edit-object/edit-object.component';
import {CreateCommandComponent} from './admin/command/create-command/create-command.component';
import {EditCommandComponent} from './admin/command/edit-command/edit-command.component';


const routes: Routes = [
  {path: '', redirectTo: 'authorization', pathMatch: 'full'},
  {path: 'authorization', component: AuthComponent, children: [
    {path: '', redirectTo: 'login', pathMatch: 'full'},
    {path: 'login', component: LoginComponent},
    {path: 'registration', component: RegistrationComponent},
    {path: '**', redirectTo: 'login'},
  ]},
  {path: 'admin', component: AdminComponent, canActivate: [SessionGuard], children: [
    {path: '', redirectTo: 'objects-list', pathMatch: 'full'},
    {path: 'objects-list', component: ObjectsListComponent},
    {path: 'create-object', component: CreateObjectComponent},
    {path: 'edit-object/:object_hash', component: EditObjectComponent},
    {path: 'create-command/:object_hash', component: CreateCommandComponent},
    {path: 'edit-command/:command_hash', component: EditCommandComponent},
    {path: 'run-tests', component: RunTestsComponent},
    {path: '**', redirectTo: 'objects-list'},
  ]},
  {path: '**', component: NotFoundComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
