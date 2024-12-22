import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { TaskControllerEditTaskRequest } from "../../typescript-fetch-client";

export const useEditTask = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskControllerEditTaskRequest) => {
      return await tasksApi
        .apiV1TasksEditPost({
          request: {
            id: task.id,
            title: task.title,
            description: task.description,
            status: task.status,
            assigneeId: task.assigneeId,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["tasks"] });
          queryClient.invalidateQueries({ queryKey: ["task", task.id] });
        });
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
