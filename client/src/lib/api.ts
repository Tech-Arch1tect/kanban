import {
  BaseAPI,
  Configuration,
  AuthApi,
  AdminApi,
  MiscApi,
} from "../typescript-fetch-client";

// @ts-ignore
const basePath = import.meta.env.VITE_API_BASE_PATH || "http://localhost:8090";

const configuration = new Configuration({
  basePath: basePath,
  credentials: "include",
});

export const api = new BaseAPI(configuration);
export const authApi = new AuthApi(configuration);
export const adminApi = new AdminApi(configuration);
export const miscApi = new MiscApi(configuration);
