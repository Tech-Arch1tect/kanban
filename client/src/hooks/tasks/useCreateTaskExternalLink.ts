import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { TaskCreateTaskExternalLinkRequest } from "../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useCreateTaskExternalLink = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (link: TaskCreateTaskExternalLinkRequest) => {
      return await tasksApi
        .apiV1TasksCreateExternalLinkPost({
          request: {
            taskId: link.taskId,
            title: link.title,
            url: link.url,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["task", link.taskId] });
        });
    },
    onSuccess: () => {
      toast.success("Task external link created successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to create task external link.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
