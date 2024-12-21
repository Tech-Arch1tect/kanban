import { useQuery } from "@tanstack/react-query";

import { tasksApi } from "../../lib/api";

export const useTaskData = ({ id }: { id: number }) => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["task", id],
    queryFn: () => tasksApi.apiV1TasksGetIdGet({ id }),
  });
  return { data, isLoading, error };
};
