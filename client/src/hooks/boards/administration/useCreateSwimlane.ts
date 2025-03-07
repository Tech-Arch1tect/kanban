import { useMutation, useQueryClient } from "@tanstack/react-query";
import { swimlanesApi } from "../../../lib/api";
import { SwimlaneCreateSwimlaneRequest } from "../../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useCreateSwimlane = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (swimlane: SwimlaneCreateSwimlaneRequest) =>
      swimlanesApi
        .apiV1SwimlanesCreatePost({
          request: swimlane,
        })
        .then(() => {
          queryClient.invalidateQueries({
            queryKey: ["boardData", String(swimlane.boardId)],
          });
        }),
    onSuccess: () => {
      toast.success("Swimlane created successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to create swimlane.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
