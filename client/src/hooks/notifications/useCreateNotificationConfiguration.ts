import { useMutation } from "@tanstack/react-query";
import { notificationsApi } from "../../lib/api";
import { toast } from "react-toastify";
import { NotificationCreateNotificationRequest } from "../../typescript-fetch-client";

export const useCreateNotificationConfiguration = () => {
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (notification: NotificationCreateNotificationRequest) => {
      return await notificationsApi.apiV1NotificationsCreatePost({
        request: notification,
      });
    },
    onSuccess: () => {
      toast.success("Notification configuration created successfully!");
    },
    onError: (error: any) => {
      toast.error(
        error.message || "Failed to create notification configuration."
      );
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
