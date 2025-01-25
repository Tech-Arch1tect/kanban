import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { TaskCreateTaskLinkRequest } from "../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useCreateTaskLink = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (link: TaskCreateTaskLinkRequest) => {
      return await tasksApi
        .apiV1TasksCreateLinkPost({
          request: {
            srcTaskId: link.srcTaskId,
            dstTaskId: link.dstTaskId,
            linkType: link.linkType,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["task", link.srcTaskId] });
          queryClient.invalidateQueries({ queryKey: ["task", link.dstTaskId] });
        });
    },
    onSuccess: () => {
      toast.success("Task link created successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to create task link.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
