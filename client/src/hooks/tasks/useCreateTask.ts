import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { TaskControllerCreateTaskRequest } from "../../typescript-fetch-client/models";

export const useCreateTask = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskControllerCreateTaskRequest) => {
      return await tasksApi
        .apiV1TasksCreatePost({
          request: {
            boardId: task.boardId,
            title: task.title,
            description: task.description,
            status: task.status,
            swimlaneId: task.swimlaneId,
            columnId: task.columnId,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["tasks"] });
        });
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
