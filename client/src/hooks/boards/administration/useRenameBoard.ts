import { useMutation, useQueryClient } from "@tanstack/react-query";
import { boardsApi } from "../../../lib/api";
import { BoardRenameBoardRequest } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useRenameBoard = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (board: BoardRenameBoardRequest) =>
      boardsApi
        .apiV1BoardsRenamePost({
          request: board,
        })
        .then((response) => {
          queryClient.invalidateQueries({
            queryKey: ["boardData", String(response?.board?.id)],
          });
        }),
    onSuccess: () => {
      toast.success("Board renamed successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to rename board.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
