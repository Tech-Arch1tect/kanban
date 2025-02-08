import { useMutation, useQueryClient } from "@tanstack/react-query";
import { notificationsApi } from "../../lib/api";
import { toast } from "react-toastify";
import { NotificationCreateNotificationRequest } from "../../typescript-fetch-client";

export const useCreateNotificationConfiguration = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (notification: NotificationCreateNotificationRequest) => {
      return await notificationsApi.apiV1NotificationsCreatePost({
        request: notification,
      });
    },
    onSuccess: () => {
      toast.success("Notification configuration created successfully!");
      queryClient.invalidateQueries({ queryKey: ["notifications"] });
    },
    onError: (error: any) => {
      toast.error(
        error.message || "Failed to create notification configuration."
      );
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
