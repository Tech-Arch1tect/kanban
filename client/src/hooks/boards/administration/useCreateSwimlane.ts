import { useMutation, useQueryClient } from "@tanstack/react-query";
import { swimlanesApi } from "../../../lib/api";
import { SwimlaneControllerCreateSwimlaneRequest } from "../../../typescript-fetch-client";

export const useCreateSwimlane = () => {
  const queryClient = useQueryClient();
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: (swimlane: SwimlaneControllerCreateSwimlaneRequest) =>
      swimlanesApi
        .apiV1SwimlanesCreatePost({
          request: swimlane,
        })
        .then(() => {
          queryClient.invalidateQueries({
            queryKey: ["boardData", String(swimlane.boardId)],
          });
        }),
    onError: (error) => {
      console.error("Error creating swimlane:", error);
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
