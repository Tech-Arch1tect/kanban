import { useMutation, useQueryClient } from "@tanstack/react-query";
import { boardsApi } from "../../../lib/api";

export const useBoardDelete = () => {
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: (boardId: number) =>
      boardsApi.apiV1BoardsDeletePost({ request: { id: boardId } }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["boards"] });
    },
    onError: (error: any) => {
      console.error("Error deleting board:", error);
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
