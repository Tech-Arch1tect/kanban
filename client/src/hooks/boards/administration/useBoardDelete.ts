import { useMutation, useQueryClient } from "@tanstack/react-query";
import { boardsApi } from "../../../lib/api";
import { toast } from "react-toastify";

export const useBoardDelete = () => {
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: (boardId: number) =>
      boardsApi.apiV1BoardsDeletePost({ request: { id: boardId } }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["boards"] });
      toast.success("Board deleted successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to delete board.");
    },
  });

  const deleteBoard = (boardId: number) => {
    if (window.confirm(`Are you sure you want to delete this board?`)) {
      mutation.mutate(boardId);
    }
  };

  return {
    deleteBoard,
    isDeleting: mutation.isPending,
    isError: mutation.isError,
    error: mutation.error,
    isSuccess: mutation.isSuccess,
  };
};
