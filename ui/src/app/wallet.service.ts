import { Injectable, EventEmitter } from '@angular/core';
import { Headers, RequestOptions, Http } from "@angular/http";
import { Router } from "@angular/router";
import { Observable } from "rxjs/Observable";
import "rxjs/add/operator/map";
import "rxjs/add/operator/do";

@Injectable()
export class WalletService {

    private host: string;
    private authenticated: EventEmitter<boolean> = new EventEmitter();
    public password: string;

    public constructor(private http: Http, private router: Router) {
        this.host = "http://localhost:12345/api";
        //this.host = "http://coin.local:12345/api";
        this.password = "";
    }

    public isAuthenticated() {
        return this.authenticated;
    }

    public authenticate(password: string) {
        let headers = new Headers({ "content-type": "application/json" });
        let options = new RequestOptions({ headers: headers });
        return this.http.post(this.host + "/authenticate?key=" + password, null, options)
            .map(result => result.json())
            .do(result => {
                if(!result.authenticated) {
                    throw new Error("The wallet does not exist or the password is incorrect");
                } else {
                    this.password = password;
                    this.authenticated.emit(true);
                }
            });
    }

    public disconnect() {
        this.password = "";
        this.authenticated.emit(false);
    }

    public create(password: string) {
        let headers = new Headers({ "content-type": "application/json" });
        let options = new RequestOptions({ headers: headers });
        return this.http.post(this.host + "/wallet?key=" + password, null, options)
            .map(result => result.json());
    }

    public dump() {
        return this.http.get(this.host + "/wallet?key=" + this.password)
            .map(result => result.json());
    }

    public getAddresses() {
        return this.http.get(this.host + "/addresses?key=" + this.password)
            .map(result => result.json());
    }

    public import(symbol: string, secret: string) {
        let headers = new Headers({ "content-type": "application/json" });
        let options = new RequestOptions({ headers: headers });
        if(secret != "") {
            return this.http.post(this.host + "/import-coin?key=" + this.password, JSON.stringify({ "symbol": symbol, "wif": secret }), options)
                .map(result => result.json());
        } else {
            return this.http.post(this.host + "/create-coin?key=" + this.password, JSON.stringify({ "symbol": symbol }), options)
                .map(result => result.json());
        }
    }

    public backup() {
        return this.http.get(this.host + "/backup")
            .map(result => result.text());
    }

    public createTransaction(sender: string, recipient: string, amount: number, txHash: string) {
        let headers = new Headers({ "content-type": "application/json" });
        let options = new RequestOptions({ headers: headers });
        return this.http.post(this.host + "/transaction?key=" + this.password, JSON.stringify({ "source_address": sender, "destination_address": recipient, "amount": amount, "txid": txHash }), options)
            .map(result => result.json());
    }

}
