import { useMutation } from '@tanstack/react-query';
import { authApi } from '../../lib/api';

export const useRegister = () => {
  return useMutation({
    mutationFn: async ({ email, password }: { email: string; password: string }) => {
      return await authApi.apiV1AuthRegisterPost({
        user: { email, password },
      });
    },
  });
};
