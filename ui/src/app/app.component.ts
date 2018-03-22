import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import { saveAs } from "file-saver/FileSaver";
import { WalletService } from "./wallet.service";

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

    public authenticated: boolean;

    public constructor(private router: Router, private wallet: WalletService) {
        this.authenticated = false;
    }

    public ngOnInit() {
        if(!this.authenticated) {
            this.router.navigate(["/unlock"], { replaceUrl: true });
        }
        this.wallet.isAuthenticated().subscribe(authenticated => {
            this.authenticated = authenticated;
            if(!authenticated) {
                this.router.navigate(["/unlock"], { replaceUrl: true });
            }
        });
    }

    public backup() {
        this.wallet.backup()
            .subscribe(result => {
                var blob = new Blob([result], { type: "text/plain" });
                saveAs(blob, "wallet.dat")
            })
    }

    public disconnect() {
        this.wallet.disconnect();
        this.router.navigate(["/unlock"]);
    }

}