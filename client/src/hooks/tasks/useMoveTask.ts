import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { TaskControllerMoveTaskRequest } from "../../typescript-fetch-client";

export const useMoveTask = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskControllerMoveTaskRequest) => {
      return await tasksApi
        .apiV1TasksMovePost({
          request: {
            taskId: task.taskId,
            columnId: task.columnId,
            swimlaneId: task.swimlaneId,
            position: task.position,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["tasks"] });
        });
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
