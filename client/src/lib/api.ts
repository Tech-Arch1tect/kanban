import {
  BaseAPI,
  Configuration,
  AuthApi,
  AdminApi,
  MiscApi,
  AuthAuthGetCSRFTokenResponse,
  SwimlanesApi,
  ColumnsApi,
  BoardsApi,
  CommentsApi,
  TasksApi,
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
export const swimlanesApi = new SwimlanesApi(configuration);
export const columnsApi = new ColumnsApi(configuration);
export const boardsApi = new BoardsApi(configuration);
export const tasksApi = new TasksApi(configuration);
export const commentsApi = new CommentsApi(configuration);

async function GetCSRFToken(): Promise<string | undefined> {
  const response: AuthAuthGetCSRFTokenResponse =
    await authApi.apiV1AuthCsrfTokenGet();
  return response.csrfToken;
}

async function csrfFetch(
  input: RequestInfo,
  init?: RequestInit
): Promise<Response> {
  const methodsRequiringCSRF = ["POST", "PUT", "DELETE", "PATCH"];
  const method = (init?.method || "GET").toUpperCase();

  if (methodsRequiringCSRF.includes(method)) {
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
