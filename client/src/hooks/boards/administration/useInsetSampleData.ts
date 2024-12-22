import { useMutation, useQueryClient } from "@tanstack/react-query";
import { sampleDataApi } from "../../../lib/api";

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
    onError: (error) => {
      console.error("Error inserting sample data:", error);
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
