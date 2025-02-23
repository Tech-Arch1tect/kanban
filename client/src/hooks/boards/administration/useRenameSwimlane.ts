import { useMutation, useQueryClient } from "@tanstack/react-query";
import { swimlanesApi } from "../../../lib/api";
import { ModelsSwimlane } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useRenameSwimlane = () => {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (swimlane: ModelsSwimlane) =>
      swimlanesApi.apiV1SwimlanesRenamePost({
        request: swimlane,
      }),
    onSuccess: ({ swimlane }) => {
      queryClient.invalidateQueries({
        queryKey: ["boardData", String(swimlane?.boardId)],
      });
      toast.success("Swimlane renamed successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to rename swimlane.");
    },
  });

  return {
    mutate: mutation.mutate,
    isRenaming: mutation.isPending,
    isError: mutation.isError,
    error: mutation.error,
    isSuccess: mutation.isSuccess,
  };
};
