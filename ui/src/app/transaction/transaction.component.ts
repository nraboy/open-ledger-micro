import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from "@angular/router";
import { WalletService } from "../wallet.service";

@Component({
    selector: 'app-transaction',
    templateUrl: './transaction.component.html',
    styleUrls: ['./transaction.component.css']
})
export class TransactionComponent implements OnInit {

    public input: any;
    public signedTx: string;

    public constructor(private router: Router, private route: ActivatedRoute, private wallet: WalletService) {
        this.input = {
            "sourceAddress": "",
            "destinationAddress": "",
            "amount": "1000",
            "txId": ""
        }
        this.signedTx = "";
    }

    public ngOnInit() {
        this.route.params.subscribe(params => {
            this.input.sourceAddress = params["source"];
        });
    }

    public create() {
        if(this.input.sourceAddress != "" && this.input.destinationAddress != "" && this.input.txId != "" && this.input.amount != "") {
            this.wallet.createTransaction(this.input.sourceAddress, this.input.destinationAddress, parseInt(this.input.amount), this.input.txId)
                .subscribe(result => {
                    this.signedTx = result.signedtx;
                });
        } else {
            console.log("All fields are required");
        }
    }

}
