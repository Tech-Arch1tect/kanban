import { useMutation, useQueryClient } from '@tanstack/react-query';
import { authApi } from '../../lib/api';

export const useConfirmTOTP = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (code: string) => {
      return await authApi.apiV1AuthTotpConfirmPost({ request: { code } });
    },
    onSuccess: (data) => {
      if (data.message === 'totp_confirmed') {
        queryClient.invalidateQueries({ queryKey: ['userProfile'] });
      }
    },
  });
};
