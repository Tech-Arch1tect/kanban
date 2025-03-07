import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { TaskDeleteTaskExternalLinkRequest } from "../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useDeleteTaskExternalLink = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (link: TaskDeleteTaskExternalLinkRequest) => {
      return await tasksApi.apiV1TasksDeleteExternalLinkPost({
        request: {
          id: link.id,
        },
      });
    },
    onSuccess: (response) => {
      queryClient.invalidateQueries({ queryKey: ["task", response.taskId] });
      toast.success("Task external link deleted successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to delete task external link.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
