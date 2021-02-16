import axios, { AxiosRequestConfig } from "axios";
import { AuthToken } from "./auth_token";
import Router from "next/router";

const get = (url: string, authString: string) => {
  const baseConfig: AxiosRequestConfig = {
    headers: {
      Authorization: authString,
      "Content-Type": "application/json",
    },
    baseURL: "http://127.0.0.1:3001",
  };
  return axios.get(url, baseConfig);
};

const getInstances = async (authString: string) => {
  let res: any = await get("/api/instances", authString);
  if (res.error) {
    return res.error;
  }
  if (res.data == "Unauthorized User Header") {
    console.log("Unauthorized User");
    Router.push("/form");
  }
  return res.data;
};

export default getInstances;
