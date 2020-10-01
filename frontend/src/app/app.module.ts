import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FetcherService } from './services/fetcher/fetcher.service';
import { LoginComponent } from './public/login/login.component';
import { RegistrationComponent } from './public/registration/registration.component';
import { InputComponent } from './shared/input/input.component';
import { FormsModule } from '@angular/forms';
import { AdminComponent } from './admin/admin.component';
import { NotFoundComponent } from './public/not-found/not-found.component';
import { NavigationComponent } from './admin/navigation/navigation.component';
import { RunTestsComponent } from './admin/run-tests/run-tests.component';
import { ObjectsListComponent } from './admin/objects-list/objects-list.component';
import { CreateObjectComponent } from './admin/create-object/create-object.component';
import { ButtonComponent } from './shared/button/button.component';
import { EditObjectComponent } from './admin/edit-object/edit-object.component';
import { ButtonModalTriggerComponent } from './shared/button-modal-trigger/button-modal-trigger.component';
import { ButtonWithIconComponent } from './shared/button-with-icon/button-with-icon.component';
import { CreateCommandComponent } from './admin/command/create-command/create-command.component';
import { CommandFieldsComponent } from './admin/command/shared/command-fields/command-fields.component';
import { EditCommandComponent } from './admin/command/edit-command/edit-command.component';
import { FloatingIconComponent } from './shared/floating-icon/floating-icon.component';
import { InputFileComponent } from './shared/input-file/input-file.component';
import { InputNumberComponent } from './shared/input-number/input-number.component';
import { BaseUrlsComponent } from './admin/general-settings/base-urls/base-urls.component';
import { TimeoutsComponent } from './admin/general-settings/timeouts/timeouts.component';
import { HeadersComponent } from './admin/general-settings/headers/headers.component';
import { CookiesComponent } from './admin/general-settings/cookies/cookies.component';
import { CommandsTableComponent } from './admin/general-settings/shared/commands-table/commands-table.component';
import { InputCheckboxComponent } from './shared/input-checkbox/input-checkbox.component';
import { HeadersAddingComponent } from './admin/shared/headers-adding/headers-adding.component';
import { CookiesAddingComponent } from './admin/shared/cookies-adding/cookies-adding.component';
import { PublicComponent } from './public/public.component';
import { FaqComponent } from './admin/faq/faq.component';
import { PublicNavigationComponent } from './public/public-navigation/public-navigation.component';
import { AboutComponent } from './public/about/about.component';
import { SidenavLinksComponent } from './shared/sidenav-links/sidenav-links.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegistrationComponent,
    InputComponent,
    AdminComponent,
    NotFoundComponent,
    NavigationComponent,
    RunTestsComponent,
    ObjectsListComponent,
    CreateObjectComponent,
    ButtonComponent,
    EditObjectComponent,
    ButtonModalTriggerComponent,
    ButtonWithIconComponent,
    CreateCommandComponent,
    CommandFieldsComponent,
    EditCommandComponent,
    FloatingIconComponent,
    InputFileComponent,
    InputNumberComponent,
    BaseUrlsComponent,
    TimeoutsComponent,
    HeadersComponent,
    CookiesComponent,
    CommandsTableComponent,
    InputCheckboxComponent,
    HeadersAddingComponent,
    CookiesAddingComponent,
    PublicComponent,
    FaqComponent,
    PublicNavigationComponent,
    AboutComponent,
    SidenavLinksComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
  ],
  providers: [
    {provide: 'Fetcher', useClass: FetcherService},
    {provide: 'FileSender', useClass: FetcherService},
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
