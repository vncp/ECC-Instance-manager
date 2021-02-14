import jwtDecode from "jwt-decode";

export type DecodedToken = {
  readonly netid: string;
  readonly exp: number;
};

export class AuthToken {
  readonly decodedToken: DecodedToken;
  constructor(readonly token?: string) {
    //Default expired token
    this.decodedToken = { netid: "", exp: 0 };
    //Then decode using jwt-decode
    try {
      if (token) this.decodedToken = jwtDecode(token);
    } catch (e) {
      console.error("Could not decode token!");
    }
  }
}
