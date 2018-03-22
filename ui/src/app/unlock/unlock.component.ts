import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import { WalletService } from "../wallet.service";

@Component({
    selector: 'app-unlock',
    templateUrl: './unlock.component.html',
    styleUrls: ['./unlock.component.css']
})
export class UnlockComponent implements OnInit {

    public input: any;

    public constructor(private router: Router, private wallet: WalletService) {
        this.input = {
            password: ""
        };
    }

    public ngOnInit() { }

    public unlock() {
        if(this.input.password == "") {
            return console.log("A password is required");
        }
        this.wallet.authenticate(this.input.password)
            .subscribe(result => {
                this.router.navigate(["/dashboard"]);
            }, error => {
                console.error(error.message);
            });
    }

}
