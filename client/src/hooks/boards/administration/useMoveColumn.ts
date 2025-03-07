import { useMutation, useQueryClient } from "@tanstack/react-query";
import { columnsApi } from "../../../lib/api";
import { ColumnMoveColumnRequest } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";
export const useMoveColumn = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (column: ColumnMoveColumnRequest) =>
      columnsApi.apiV1ColumnsMovePost({ request: column }),
    onSuccess: (column) => {
      queryClient.invalidateQueries({
        queryKey: ["boardData", String(column.column?.boardId)],
      });
      toast.success("Column moved successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to move column.");
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
