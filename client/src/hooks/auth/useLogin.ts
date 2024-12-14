import { useMutation, useQueryClient } from '@tanstack/react-query';
import { authApi } from '../../lib/api';

export const useLogin = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async ({ email, password }: { email: string; password: string }) => {
      return await authApi.apiV1AuthLoginPost({
        user: { email, password },
      });
    },
    onSuccess: (data) => {
      // If TOTP is not required, invalidate user profile
      if (data.message !== 'totp_required') {
        queryClient.invalidateQueries({ queryKey: ['userProfile'] });
      }
    },
  });
};
