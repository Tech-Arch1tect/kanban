import { useMutation, useQueryClient } from "@tanstack/react-query";
import { columnsApi } from "../../../lib/api";
import { ModelsColumn } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useDeleteColumn = () => {
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: (column: ModelsColumn) =>
      columnsApi.apiV1ColumnsDeletePost({
        request: column,
      }),
    onSuccess: ({ column }) => {
      queryClient.invalidateQueries({
        queryKey: ["boardData", String(column?.boardId)],
      });
      toast.success("Column deleted successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to delete column.");
    },
  });

  const deleteColumn = (column: ModelsColumn) => {
    if (window.confirm(`Are you sure you want to delete this column?`)) {
      mutation.mutate(column);
    }
  };

  return {
    deleteColumn,
    isDeleting: mutation.isPending,
    isError: mutation.isError,
    error: mutation.error,
    isSuccess: mutation.isSuccess,
  };
};
