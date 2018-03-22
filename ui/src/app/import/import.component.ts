import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import { WalletService } from "../wallet.service";

@Component({
    selector: 'app-import',
    templateUrl: './import.component.html',
    styleUrls: ['./import.component.css']
})
export class ImportComponent implements OnInit {


    public input: any;
    public types: Array<any>;

    public constructor(private router: Router, private wallet: WalletService) {
        this.input = {
            "type": "",
            "secret": ""
        }
    }

    public ngOnInit() { }

    public import() {
        if(this.input.type != "") {
            this.wallet.import(this.input.type, this.input.secret)
                .subscribe(result => {
                    this.router.navigate(["/dashboard"]);
                }, error => {
                    console.error(error);
                });
        }
    }

}