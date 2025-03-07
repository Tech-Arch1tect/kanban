import { useMutation, useQueryClient } from "@tanstack/react-query";
import { swimlanesApi } from "../../../lib/api";
import { SwimlaneMoveSwimlaneRequest } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useMoveSwimlane = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (swimlane: SwimlaneMoveSwimlaneRequest) =>
      swimlanesApi.apiV1SwimlanesMovePost({ request: swimlane }),
    onSuccess: (swimlane) => {
      queryClient.invalidateQueries({
        queryKey: ["boardData", String(swimlane.swimlane?.boardId)],
      });
      toast.success("Swimlane moved successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to move swimlane.");
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
