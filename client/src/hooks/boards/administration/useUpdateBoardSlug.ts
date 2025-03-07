import { useMutation, useQueryClient } from "@tanstack/react-query";
import { boardsApi } from "../../../lib/api";
import { BoardUpdateBoardSlugRequest } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useUpdateBoardSlug = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (board: BoardUpdateBoardSlugRequest) =>
      boardsApi
        .apiV1BoardsUpdateSlugPost({
          request: board,
        })
        .then((response) => {
          queryClient.invalidateQueries({
            queryKey: ["boardData", String(response?.board?.id)],
          });
          queryClient.invalidateQueries({
            queryKey: ["boards"],
          });
        }),
    onSuccess: () => {
      toast.success("Board slug updated successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to update board slug.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
