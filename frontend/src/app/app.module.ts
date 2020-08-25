import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import {FetcherService} from './services/fetcher/fetcher.service';
import { LoginComponent } from './auth/login/login.component';
import { RegistrationComponent } from './auth/registration/registration.component';
import { AuthComponent } from './auth/auth.component';
import { InputComponent } from './shared/input/input.component';
import {FormsModule} from '@angular/forms';
import { AdminComponent } from './admin/admin.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { NavigationComponent } from './admin/navigation/navigation.component';
import { RunTestsComponent } from './admin/run-tests/run-tests.component';
import { ObjectsListComponent } from './admin/objects-list/objects-list.component';
import { CreateObjectComponent } from './admin/create-object/create-object.component';
import { ButtonComponent } from './shared/button/button.component';
import { RouterButtonComponent } from './shared/router-button/router-button.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegistrationComponent,
    AuthComponent,
    InputComponent,
    AdminComponent,
    NotFoundComponent,
    NavigationComponent,
    RunTestsComponent,
    ObjectsListComponent,
    CreateObjectComponent,
    ButtonComponent,
    RouterButtonComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule
  ],
  providers: [
    {provide: 'Fetcher', useClass: FetcherService}
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
