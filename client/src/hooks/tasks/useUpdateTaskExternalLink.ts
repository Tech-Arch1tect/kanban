import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { TaskUpdateTaskExternalLinkRequest } from "../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useUpdateTaskExternalLink = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (link: TaskUpdateTaskExternalLinkRequest) => {
      return await tasksApi.apiV1TasksUpdateExternalLinkPost({
        request: {
          id: link.id,
          title: link.title,
          url: link.url,
        },
      });
    },
    onSuccess: (response) => {
      queryClient.invalidateQueries({
        queryKey: ["task", response.link?.taskId],
      });
      toast.success("Task external link updated successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to update task external link.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
