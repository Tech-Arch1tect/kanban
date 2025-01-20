import { ModelsFile, ModelsTask } from "../../../typescript-fetch-client";
import { UploadFile } from "./UploadFile";
import { useDeleteFile } from "../../../hooks/tasks/Files/useDeleteFile";
import { useDownloadFile } from "../../../hooks/tasks/Files/useDownloadFile";
import { useGetImage } from "../../../hooks/tasks/Files/useGetImage";
import { toast } from "react-toastify";
import { useState } from "react";

export const ViewFiles = ({ task }: { task: ModelsTask }) => {
  const { mutate: deleteFile, isPending: isDeleting } = useDeleteFile();
  const { mutate: downloadFile, isPending: isDownloading } = useDownloadFile();
  const [image, setImage] = useState<string | null>(null);
  const {
    mutate: getImage,
    isPending: isFetchingImage,
    data: imageData,
  } = useGetImage();

  const handleDelete = (file: ModelsFile) => {
    if (confirm(`Are you sure you want to delete the file "${file.name}"?`)) {
      deleteFile(file);
    }
  };

  const handleDownload = (fileId: number) => {
    downloadFile(fileId);
  };

  const handleViewImage = (fileId: number) => {
    getImage(fileId, {
      onSuccess: (data) => {
        if (data.content) {
          setImage(`${data.content}`);
        } else {
          toast.error("File is not an image.");
        }
      },
    });
  };

  return (
    <div className="view-files-container bg-white dark:bg-gray-800 rounded-lg shadow-sm">
      <h2 className="text-lg font-semibold text-gray-800 dark:text-gray-200 mb-4">
        Task Files
      </h2>
      <div className="files-list space-y-2">
        {task.files && task.files.length > 0 ? (
          task.files.map((file) => (
            <div
              key={file.id}
              className="file-item flex items-center justify-between p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-gray-50 dark:bg-gray-700"
            >
              <span className="text-sm text-gray-700 dark:text-gray-200">
                {file.name}
              </span>
              <div className="flex items-center space-x-3">
                <button
                  onClick={() => handleDownload(file.id ?? 0)}
                  disabled={isDownloading}
                  className={`px-3 py-1 text-sm rounded-md text-white ${
                    isDownloading
                      ? "bg-gray-400 dark:bg-gray-600 cursor-not-allowed"
                      : "bg-blue-500 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-700"
                  }`}
                >
                  {isDownloading ? "Downloading..." : "Download"}
                </button>
                {file.type === "image" && (
                  <button
                    onClick={() => handleViewImage(file.id ?? 0)}
                    disabled={isFetchingImage}
                    className={`px-3 py-1 text-sm rounded-md text-white ${
                      isFetchingImage
                        ? "bg-gray-400 dark:bg-gray-600 cursor-not-allowed"
                        : "bg-green-500 hover:bg-green-600 dark:bg-green-600 dark:hover:bg-green-700"
                    }`}
                  >
                    {isFetchingImage ? "Opening..." : "View Image"}
                  </button>
                )}
                <button
                  onClick={() => handleDelete(file)}
                  disabled={isDeleting}
                  className={`px-3 py-1 text-sm rounded-md text-white ${
                    isDeleting
                      ? "bg-gray-400 dark:bg-gray-600 cursor-not-allowed"
                      : "bg-red-500 hover:bg-red-600 dark:bg-red-600 dark:hover:bg-red-700"
                  }`}
                >
                  {isDeleting ? "Deleting..." : "Delete"}
                </button>
              </div>
            </div>
          ))
        ) : (
          <p className="text-gray-500 dark:text-gray-400">
            No files uploaded yet.
          </p>
        )}
      </div>
      <div className="mt-6">
        <UploadFile taskId={task.id ?? 0} />
      </div>
      {image && (
        <img
          src={image}
          alt="Task Image"
          className="mt-4 rounded-md shadow-sm"
        />
      )}
    </div>
  );
};
