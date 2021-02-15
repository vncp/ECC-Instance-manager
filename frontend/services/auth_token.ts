import jwtDecode from "jwt-decode";
import Router from "next/router";
import Cookie from "js-cookie";
import { redirectToLogin } from "../services/redirect_service";

export type DecodedToken = {
  readonly netid: string;
  readonly expiry: number;
};

export const TOKEN_STORAGE_KEY = "linuxremote.authToken";

export class AuthToken {
  readonly decodedToken: DecodedToken;

  constructor(readonly token?: string) {
    //Default expired token
    this.decodedToken = { netid: "", expiry: 0 };
    //Then decode using jwt-decode
    console.log(token);
    try {
      if (token) {
        this.decodedToken = jwtDecode(token);
      }
    } catch (e) {
      console.error("Could not decode token!");
    }
  }

  get authorizationString() {
    return `Bearer ${this.token}`;
  }

  get expiresAt(): Date {
    return new Date(this.decodedToken.expiry * 1000);
  }

  get isExpired(): boolean {
    return new Date() > this.expiresAt;
  }

  get isValid(): boolean {
    return !this.isExpired;
  }

  static async storeToken(token: string) {
    Cookie.set(TOKEN_STORAGE_KEY, token);
    await Router.push("/dashboard");
  }

  static async logout() {
    Cookie.remove(TOKEN_STORAGE_KEY);
    await redirectToLogin();
  }

  logout = AuthToken.logout;
}
