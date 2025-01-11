import { useMutation } from "@tanstack/react-query";
import { tasksApi } from "../../../lib/api";
import { toast } from "react-toastify";

export const useGetImage = () => {
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (id: number) => {
      return await tasksApi.apiV1TasksGetImageFileIdGet({
        fileId: id,
      });
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to download image.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
