import axios, { AxiosRequestConfig } from "axios";
import { LoginInput } from "../pages/login";
import { AuthToken } from "./auth_token";

type errorMessage = string;

const baseConfig: AxiosRequestConfig = {
  baseURL: "http://localhost:3001",
};

const post = (url: string, data: URLSearchParams) => {
  return axios.post(url, data, baseConfig);
};

export const postLogin = async (
  inputs: LoginInput
): Promise<errorMessage | void> => {
  const data = new URLSearchParams(inputs);
  const res: any = await post("/api/login", data);
  if (res.error) {
    return res.error;
  }
  if (res.data && res.data.token) {
    await AuthToken.storeToken(res.data.token);
    return;
  }
  return "Something unexected happened";
};
