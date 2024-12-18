import { useMutation, useQueryClient } from "@tanstack/react-query";
import { swimlanesApi } from "../../../lib/api";
import { SwimlaneControllerMoveSwimlaneRequest } from "../../../typescript-fetch-client";

export const useMoveSwimlane = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (swimlane: SwimlaneControllerMoveSwimlaneRequest) =>
      swimlanesApi.apiV1SwimlanesMovePost({ request: swimlane }),
    onSuccess: (swimlane) => {
      queryClient.invalidateQueries({
        queryKey: ["boardData", String(swimlane.swimlane?.boardId)],
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
