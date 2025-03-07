import { useMutation, useQueryClient } from "@tanstack/react-query";
import { columnsApi } from "../../../lib/api";
import { ColumnCreateColumnRequest } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useCreateColumn = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (column: ColumnCreateColumnRequest) =>
      columnsApi
        .apiV1ColumnsCreatePost({
          request: column,
        })
        .then(() => {
          queryClient.invalidateQueries({
            queryKey: ["boardData", String(column.boardId)],
          });
        }),
    onSuccess: () => {
      toast.success("Column created successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to create column.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
