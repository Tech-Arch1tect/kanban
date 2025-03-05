import { useQuery } from "@tanstack/react-query";
import { tasksApi } from "../../../lib/api";
import { toast } from "react-toastify";

export const useGetImage = (id: number) => {
  const q = useQuery({
    queryKey: ["image", id],
    queryFn: () => tasksApi.apiV1TasksGetImageFileIdGet({ fileId: id }),
    staleTime: 10 * 60 * 1000, // 10 minutes
    gcTime: 60 * 60 * 1000, // 1 hour
  });

  if (q.error) {
    toast.error(q.error.message || "Failed to download image.");
  }

  return q;
};
