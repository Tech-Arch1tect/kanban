import { useMutation, useQueryClient } from "@tanstack/react-query";
import { swimlanesApi } from "../../../lib/api";
import { ModelsSwimlane } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useDeleteSwimlane = () => {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (swimlane: ModelsSwimlane) =>
      swimlanesApi.apiV1SwimlanesDeletePost({
        request: swimlane,
      }),
    onSuccess: ({ swimlane }) => {
      queryClient.invalidateQueries({
        queryKey: ["boardData", String(swimlane?.boardId)],
      });
      toast.success("Swimlane deleted successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to delete swimlane.");
    },
  });

  const deleteSwimlane = (swimlane: ModelsSwimlane) => {
    if (window.confirm(`Are you sure you want to delete this swimlane?`)) {
      mutation.mutate(swimlane);
    }
  };

  return {
    deleteSwimlane,
    isDeleting: mutation.isPending,
    isError: mutation.isError,
    error: mutation.error,
    isSuccess: mutation.isSuccess,
  };
};
