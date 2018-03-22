import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule } from "@angular/router";
import { FormsModule } from "@angular/forms";
import { HttpModule } from "@angular/http";

import { AppComponent } from './app.component';
import { UnlockComponent } from './unlock/unlock.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { GenerateComponent } from './generate/generate.component';
import { WalletService } from "./wallet.service";
import { ImportComponent } from './import/import.component';
import { TransactionComponent } from './transaction/transaction.component';

const routes = [
    { path: "", redirectTo: "/unlock", pathMatch: "full" },
    { path: "unlock", component: UnlockComponent },
    { path: "generate", component: GenerateComponent },
    { path: "dashboard", component: DashboardComponent },
    { path: "import", component: ImportComponent },
    { path: "transaction/:source", component: TransactionComponent },
];

@NgModule({
  declarations: [
    AppComponent,
    UnlockComponent,
    DashboardComponent,
    GenerateComponent,
    ImportComponent,
    TransactionComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    RouterModule,
    RouterModule.forRoot(routes, { useHash: true })
  ],
  providers: [WalletService],
  bootstrap: [AppComponent]
})
export class AppModule { }
