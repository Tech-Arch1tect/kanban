import { useMutation, useQueryClient } from "@tanstack/react-query";
import { columnsApi } from "../../../lib/api";
import { ColumnControllerCreateColumnRequest } from "../../../typescript-fetch-client";

export const useCreateColumn = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (column: ColumnControllerCreateColumnRequest) =>
      columnsApi
        .apiV1ColumnsCreatePost({
          request: column,
        })
        .then(() => {
          queryClient.invalidateQueries({
            queryKey: ["boardData", String(column.boardId)],
          });
        }),
    onError: (error) => {
      console.error("Error creating column:", error);
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
