import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from './public/login/login.component';
import {RegistrationComponent} from './public/registration/registration.component';
import {AdminComponent} from './admin/admin.component';
import {NotFoundComponent} from './public/not-found/not-found.component';
import {SessionGuard} from './session.guard';
import {ObjectsListComponent} from './admin/objects-list/objects-list.component';
import {RunTestsComponent} from './admin/run-tests/run-tests.component';
import {CreateObjectComponent} from './admin/create-object/create-object.component';
import {EditObjectComponent} from './admin/edit-object/edit-object.component';
import {CreateCommandComponent} from './admin/command/create-command/create-command.component';
import {EditCommandComponent} from './admin/command/edit-command/edit-command.component';
import {BaseUrlsComponent} from './admin/general-settings/base-urls/base-urls.component';
import {TimeoutsComponent} from './admin/general-settings/timeouts/timeouts.component';
import {HeadersComponent} from './admin/general-settings/headers/headers.component';
import {CookiesComponent} from './admin/general-settings/cookies/cookies.component';
import {FaqComponent} from './admin/faq/faq.component';
import {PublicComponent} from './public/public.component';
import {AboutComponent} from './public/about/about.component';


const routes: Routes = [
  {path: '', redirectTo: 'public', pathMatch: 'full'},
  {path: 'admin', component: AdminComponent, canActivate: [SessionGuard], children: [
    {path: '', redirectTo: 'objects-list', pathMatch: 'full'},
    {path: 'objects-list', component: ObjectsListComponent},
    {path: 'create-object', component: CreateObjectComponent},
    {path: 'edit-object/:object_hash', component: EditObjectComponent},
    {path: 'create-command/:object_hash', component: CreateCommandComponent},
    {path: 'edit-command/:command_hash', component: EditCommandComponent},
    {path: 'run-tests', component: RunTestsComponent},
    {path: 'general-base-urls', component: BaseUrlsComponent},
    {path: 'general-timeouts', component: TimeoutsComponent},
    {path: 'general-headers', component: HeadersComponent},
    {path: 'general-cookies', component: CookiesComponent},
    {path: 'faq', component: FaqComponent},
    {path: 'about', component: AboutComponent},
    {path: '**', component: NotFoundComponent},
  ]},
  {path: 'public', component: PublicComponent, children: [
    {path: '', redirectTo: 'sign-in', pathMatch: 'full'},
    {path: 'sign-in', component: LoginComponent},
    {path: 'sign-up', component: RegistrationComponent},
    {path: 'about', component: AboutComponent},
    {path: '**', component: NotFoundComponent},
  ]},
  {path: '**', redirectTo: 'public/not-found'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
