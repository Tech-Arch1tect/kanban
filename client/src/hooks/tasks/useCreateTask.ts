import { useMutation, useQueryClient } from "@tanstack/react-query";
import { tasksApi } from "../../lib/api";
import { TaskCreateTaskRequest } from "../../typescript-fetch-client";
import { toast } from "react-toastify";

export const useCreateTask = () => {
  const queryClient = useQueryClient();

  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (task: TaskCreateTaskRequest) => {
      return await tasksApi
        .apiV1TasksCreatePost({
          request: {
            parentTaskId: task.parentTaskId,
            boardId: task.boardId,
            title: task.title,
            description: task.description,
            status: task.status,
            swimlaneId: task.swimlaneId,
            columnId: task.columnId,
            assigneeId: task.assigneeId,
          },
        })
        .then(() => {
          queryClient.invalidateQueries({ queryKey: ["tasks"] });
          if (task.parentTaskId) {
            queryClient.invalidateQueries({
              queryKey: ["task", task.parentTaskId],
            });
          }
        });
    },
    onSuccess: () => {
      toast.success("Task created successfully!");
    },
    onError: (error: any) => {
      toast.error(error.message || "Failed to create task.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
