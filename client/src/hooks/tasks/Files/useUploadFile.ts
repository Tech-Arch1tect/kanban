import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../../lib/api";
import { toast } from "react-toastify";
import { TaskUploadFileRequest } from "../../../typescript-fetch-client";

export const useUploadFile = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (file: TaskUploadFileRequest) => {
      return await tasksApi
        .apiV1TasksUploadPost({
          request: {
            file: file.file,
            name: file.name,
            taskId: file.taskId,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["task", file.taskId] });
        });
    },
    onSuccess: () => {
      toast.success("File uploaded successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to upload file.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
