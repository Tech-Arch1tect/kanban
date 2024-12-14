import { useMutation, useQueryClient } from "@tanstack/react-query";
import { boardApi } from "../../lib/api";
import {
  ApiV1BoardsCreatePostRequest,
  BoardControllerCreateBoardResponse,
} from "../../typescript-fetch-client";

export const useCreateBoard = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation<
    BoardControllerCreateBoardResponse,
    Error,
    ApiV1BoardsCreatePostRequest
  >({
    mutationFn: async (board) => boardApi.apiV1BoardsCreatePost(board),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["boards"] });
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
