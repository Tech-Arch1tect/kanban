import { useMutation } from "@tanstack/react-query";
import { tasksApi } from "../../../lib/api";
import { toast } from "react-toastify";

export const useDownloadFile = () => {
  const { mutate, error, isError, isSuccess, data, isPending } = useMutation({
    mutationFn: async (id: number) => {
      return await tasksApi.apiV1TasksDownloadFileIdGet({
        fileId: id,
      });
    },
    onSuccess: (data) => {
      if (!data.content) {
        toast.error("File not found.");
        return;
      }

      const binaryString = window.atob(data.content);
      const binaryLength = binaryString.length;
      const bytes = new Uint8Array(binaryLength);

      for (let i = 0; i < binaryLength; i++) {
        bytes[i] = binaryString.charCodeAt(i);
      }

      const blob = new Blob([bytes], { type: data.file?.type });
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url;
      a.download = data.file?.name || "file";
      a.click();
      window.URL.revokeObjectURL(url);
      toast.success("File downloaded successfully!");
    },
    onError: (error: Error) => {
      toast.error(error.message || "Failed to download file.");
    },
  });

  return { mutate, error, isError, isSuccess, data, isPending };
};
