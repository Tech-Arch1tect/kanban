import { useMutation, useQueryClient } from "@tanstack/react-query";
import { swimlanesApi } from "../../../lib/api";
import { SwimlaneControllerMoveSwimlaneRequest } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useMoveSwimlane = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (swimlane: SwimlaneControllerMoveSwimlaneRequest) =>
      swimlanesApi.apiV1SwimlanesMovePost({ request: swimlane }),
    onSuccess: (swimlane) => {
      queryClient.invalidateQueries({
        queryKey: ["boardData", String(swimlane.swimlane?.boardId)],
      });
      toast.success("Swimlane moved successfully!");
    },
    onError: (error: any) => {
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
