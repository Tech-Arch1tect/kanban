import { useMutation, useQueryClient } from "@tanstack/react-query";
import { columnsApi } from "../../../lib/api";
import { ModelsColumn } from "../../../typescript-fetch-client";

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
