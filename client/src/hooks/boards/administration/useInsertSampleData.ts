import { useMutation, useQueryClient } from "@tanstack/react-query";
import { sampleDataApi } from "../../../lib/api";
import { toast } from "react-toastify";

export const useInsertSampleData = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation<
    void,
    Error,
    { boardId: number; numTasks: number; numComments: number }
  >({
    mutationFn: ({ boardId, numTasks, numComments }) =>
      sampleDataApi
        .apiV1SampleDataInsertPost({
          request: {
            boardId: boardId,
            numTasks: numTasks,
            numComments: numComments,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({
            queryKey: ["tasks", String(boardId)],
          });
        }),
    onSuccess: () => {
      toast.success("Sample data inserted successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to insert sample data.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
