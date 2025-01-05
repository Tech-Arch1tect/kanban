import { useMutation } from "@tanstack/react-query";
import { authApi } from "../../lib/api";

export const useRegister = () => {
  return useMutation({
    mutationFn: async ({
      email,
      password,
      username,
    }: {
      email: string;
      password: string;
      username: string;
    }) => {
      return await authApi.apiV1AuthRegisterPost({
        user: { email, password, username },
      });
    },
  });
};
