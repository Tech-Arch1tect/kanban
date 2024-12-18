import { useMutation, useQueryClient } from "@tanstack/react-query";
import { columnsApi } from "../../../lib/api";
import { ColumnControllerMoveColumnRequest } from "../../../typescript-fetch-client/models/ColumnControllerMoveColumnRequest";

export const useMoveColumn = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (column: ColumnControllerMoveColumnRequest) =>
      columnsApi.apiV1ColumnsMovePost({ request: column }),
    onSuccess: (column) => {
      queryClient.invalidateQueries({
        queryKey: ["boardData", String(column.column?.boardId)],
      });
    },
  });

  return {
    mutate,
    error,
    isError,
    isSuccess,
    data,
    isPending,
  };
};
