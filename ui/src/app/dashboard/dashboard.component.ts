import { Component, OnInit } from '@angular/core';
import { Http } from "@angular/http";
import { Observable } from "rxjs/Observable";
import { WalletService } from "../wallet.service";

@Component({
    selector: 'app-dashboard',
    templateUrl: './dashboard.component.html',
    styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

    public coins: any;

    public constructor(private http: Http, private wallet: WalletService) {
        this.coins = [];
    }

    public ngOnInit() {
        this.wallet.getAddresses()
            .subscribe(result => {
                this.coins = result.coins;
            }, error => {
                console.error(error);
            });
    }

}