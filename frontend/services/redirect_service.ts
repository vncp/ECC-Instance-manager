import { ServerResponse } from "http";
import Router from "next/router";

export const redirectToLogin = (server?: ServerResponse) => {
  const loginURL = "/?redirected=true";
  if (server) {
    server.writeHead(302, { Location: loginURL });
    server.end();
  } else {
    Router.push(loginURL);
  }
};

export const redirectToRequestForm = (server?: ServerResponse) => {
  const requestURL = "/form?redirected=true";
  if (server) {
    server.writeHead(302, { Location: requestURL });
    server.end();
  } else {
    Router.push(requestURL);
  }
};
