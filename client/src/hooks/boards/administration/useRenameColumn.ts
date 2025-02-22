import { useMutation, useQueryClient } from "@tanstack/react-query";
import { columnsApi } from "../../../lib/api";
import { ColumnRenameColumnRequest } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useRenameColumn = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (column: ColumnRenameColumnRequest) =>
      columnsApi
        .apiV1ColumnsRenamePost({
          request: column,
        })
        .then((response) => {
          queryClient.invalidateQueries({
            queryKey: ["boardData", String(response?.column?.boardId)],
          });
        }),
    onSuccess: () => {
      toast.success("Column renamed successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to rename column.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
