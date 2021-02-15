import axios, { AxiosRequestConfig } from "axios";
import { AuthToken } from "./auth_token";

type errorMessage = string;

const get = (url: string, authString: string) => {
  const baseConfig: AxiosRequestConfig = {
    headers: {
      Authorization: authString,
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*",
    },
    baseURL: "http://localhost:3001",
  };
  return axios.get(url, baseConfig);
};

const getInstances = async (
  authString: string
): Promise<errorMessage | void> => {
  const res: any = await get("/api/instances", authString);
  if (res.error) {
    return res.error;
  }
  if (res.data) {
    console.log(res.data);
  }
};

export default getInstances;
