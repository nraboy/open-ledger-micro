import { Component, OnInit } from '@angular/core';
import { Headers, RequestOptions, Http } from "@angular/http";
import { Router } from "@angular/router";
import { WalletService } from "../wallet.service";

@Component({
    selector: 'app-generate',
    templateUrl: './generate.component.html',
    styleUrls: ['./generate.component.css']
})
export class GenerateComponent implements OnInit {

    public input: any;

    public constructor(private http: Http, private router: Router, private wallet: WalletService) {
        this.input = {
            "password": "",
            "confirmPassword": ""
        };
    }

    public ngOnInit() { }

    public create() {
        if(this.input.password != "" && this.input.confirmPassword != "") {
            if(this.input.password == this.input.confirmPassword) {
                this.wallet.create(this.input.password)
                    .subscribe(result => {
                        this.router.navigate(["/unlock"]);
                    });
            } else {
                console.error("The passwords to not match");
            }
        } else {
            console.error("All fields must be filled");
        }
    }

}