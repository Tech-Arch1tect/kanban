import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../../lib/api";
import { toast } from "react-toastify";
import { ModelsFile } from "../../../typescript-fetch-client";

export const useDeleteFile = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (file: ModelsFile) => {
      return await tasksApi
        .apiV1TasksDeleteFilePost({
          request: {
            fileId: file.id,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["task", file.taskId] });
        });
    },
    onSuccess: () => {
      toast.success("File deleted successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to delete file.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
