import {
  BaseAPI,
  Configuration,
  AuthApi,
  AdminApi,
  MiscApi,
  AuthControllerGetCSRFTokenResponse,
  BoardsApi,
  SwimlanesApi,
  TasksApi,
  CommentsApi,
  ColumnsApi,
  SampleDataApi,
} from "../typescript-fetch-client";

// @ts-ignore
const basePath = import.meta.env.VITE_API_BASE_PATH || "http://localhost:8090";

const configuration = new Configuration({
  basePath: basePath,
  credentials: "include",
  fetchApi: csrfFetch,
});

export const api = new BaseAPI(configuration);
export const authApi = new AuthApi(configuration);
export const adminApi = new AdminApi(configuration);
export const miscApi = new MiscApi(configuration);
export const boardsApi = new BoardsApi(configuration);
export const swimlanesApi = new SwimlanesApi(configuration);
export const tasksApi = new TasksApi(configuration);
export const commentsApi = new CommentsApi(configuration);
export const columnsApi = new ColumnsApi(configuration);
export const sampleDataApi = new SampleDataApi(configuration);

async function GetCSRFToken(): Promise<string | undefined> {
  const response: AuthControllerGetCSRFTokenResponse =
    await authApi.apiV1AuthCsrfTokenGet();
  return response.csrfToken;
}

async function csrfFetch(
  input: RequestInfo,
  init?: RequestInit
): Promise<Response> {
  const methodsRequiringCSRF = ["POST", "PUT", "DELETE", "PATCH"];
  const method = (init?.method || "GET").toUpperCase();

  if (
    methodsRequiringCSRF.includes(method) &&
    !input.toString().includes("api/v1/auth/register") &&
    !input.toString().includes("api/v1/auth/login") &&
    !input.toString().includes("api/v1/auth/totp/confirm")
  ) {
    const csrfToken = await GetCSRFToken();
    if (csrfToken) {
      init = {
        ...init,
        headers: {
          ...(init?.headers || {}),
          "X-CSRF-Token": csrfToken,
        },
      };
    } else {
      throw new Error("Unable to obtain CSRF token.");
    }
  }

  return fetch(input, init);
}
