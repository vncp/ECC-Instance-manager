import axios, { AxiosRequestConfig } from "axios";
import { AuthToken } from "./auth_token";
import Router from "next/router";

const getReq = (url: string, authString: string) => {
  const baseConfig: AxiosRequestConfig = {
    headers: {
      Authorization: authString,
      "Content-Type": "application/json",
    },
    baseURL: "http://127.0.0.1:3001",
  };
  return axios.get(url, baseConfig);
};

export const getRequests = async (authString: string) => {
  let res: any = await getReq("/api/requests", authString);
  if (res.error) {
    return res.error;
  }
  if (res.data == "Unauthorized User Header") {
    console.log("Unauthorized user");
    Router.push("/login");
  }
  return res.data;
};

export const getInstances = async (authString: string) => {
  let res: any = await getReq("/api/instances", authString);
  if (res.error) {
    return res.error;
  }
  if (res.data == "Unauthorized User Header") {
    console.log("Unauthorized User");
    Router.push("/form");
  }
  return res.data;
};
