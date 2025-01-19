import { useMutation, useQueryClient } from "@tanstack/react-query";
import { settingsApi } from "../../lib/api";
import { toast } from "react-toastify";
import { ModelsSettings } from "../../typescript-fetch-client/models/ModelsSettings";

export const useUpdateSettings = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (settings: ModelsSettings) => {
      return await settingsApi
        .apiV1SettingsUpdatePost({
          request: {
            settings: settings,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["settings"] });
        });
    },
    onSuccess: () => {
      toast.success("Settings updated successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to update settings.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
