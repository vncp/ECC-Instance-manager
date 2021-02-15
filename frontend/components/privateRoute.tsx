import ServerCookie from "next-cookies";
import { NextPageContext } from "next";
import React, { Component } from "react";
import { TOKEN_STORAGE_KEY, AuthToken } from "../services/auth_token";
import { redirectToLogin } from "../services/redirect_service";

export type AuthProps = {
  token: string;
};

export function privateRoute(WrappedComponent: any) {
  return class extends Component<AuthProps> {
    state = {
      auth: new AuthToken(this.props.token),
    };

    static async getInitialProps(ctx: NextPageContext) {
      const token = ServerCookie(ctx)[TOKEN_STORAGE_KEY];
      const auth = new AuthToken(token);
      const initialProps = { auth };

      if (auth.isExpired) redirectToLogin(ctx.res);
      if (WrappedComponent.getInitialProps) {
        const wrappedProps = await WrappedComponent.Component.getInitialProps(
          initialProps
        );
        return { ...wrappedProps, auth };
      }
      return initialProps;
    }

    componentDidMount(): void {
      this.setState({ auth: new AuthToken(this.props.token) });
    }

    render() {
      const { auth, ...propsWithoutAuth } = this.props;
      return <WrappedComponent auth={this.props.auth} {...propsWithoutAuth} />;
    }
  };
}
