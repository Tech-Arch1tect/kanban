import { useMutation, useQueryClient } from "@tanstack/react-query";
import { boardsApi } from "../../lib/api";
import {
  ApiV1BoardsCreatePostRequest,
  BoardCreateBoardResponse,
} from "../../typescript-fetch-client";
import { toast } from "react-toastify";
export const useCreateBoard = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation<
    BoardCreateBoardResponse,
    Error,
    ApiV1BoardsCreatePostRequest
  >({
    mutationFn: async (board) => boardsApi.apiV1BoardsCreatePost(board),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["boards"] });
      toast.success("Board created successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to create board.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
