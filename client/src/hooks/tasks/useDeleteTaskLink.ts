import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { TaskDeleteTaskLinkRequest } from "../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useDeleteTaskLink = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (link: TaskDeleteTaskLinkRequest) => {
      return await tasksApi
        .apiV1TasksDeleteLinkPost({
          request: {
            linkId: link.linkId,
          },
        })
        .then((response) => {
          queryClient.invalidateQueries({
            queryKey: ["task", response.link?.srcTaskId],
          });
          queryClient.invalidateQueries({
            queryKey: ["task", response.link?.dstTaskId],
          });
        });
    },
    onSuccess: () => {
      toast.success("Task link deleted successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to delete task link.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
