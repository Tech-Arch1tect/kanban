import { useState, useEffect } from "react";
import { ModelsFile, ModelsTask } from "../../../typescript-fetch-client";
import { UploadFile } from "./UploadFile";
import { useDeleteFile } from "../../../hooks/tasks/Files/useDeleteFile";
import { useDownloadFile } from "../../../hooks/tasks/Files/useDownloadFile";
import { useGetImage } from "../../../hooks/tasks/Files/useGetImage";
import { toast } from "react-toastify";
import {
  ChevronDownIcon,
  ChevronUpIcon,
  ArrowDownTrayIcon,
  TrashIcon,
  XMarkIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
} from "@heroicons/react/24/outline";
import { ImageThumbnail } from "./ImageThumbnail";

export const ViewFiles = ({ task }: { task: ModelsTask }) => {
  const [showFiles, setShowFiles] = useState(false);

  const [galleryIndex, setGalleryIndex] = useState<number | null>(null);
  const [currentImageSrc, setCurrentImageSrc] = useState<string | null>(null);

  const { mutate: deleteFile, isPending: isDeleting } = useDeleteFile();
  const { mutate: downloadFile, isPending: isDownloading } = useDownloadFile();
  const { mutate: getImage } = useGetImage();

  const imageFiles = task.files?.filter((file) => file.type === "image") || [];
  const nonImageFiles =
    task.files?.filter((file) => file.type !== "image") || [];

  const handleDelete = (file: ModelsFile) => {
    if (confirm(`Are you sure you want to delete "${file.name}"?`)) {
      deleteFile(file);
    }
  };

  const handleDownload = (fileId: number) => {
    downloadFile(fileId);
  };

  const openGallery = (index: number) => {
    const file = imageFiles[index];
    getImage(file.id as number, {
      onSuccess: (data) => {
        if (data.content) {
          setCurrentImageSrc(data.content);
          setGalleryIndex(index);
        } else {
          toast.error("File is not an image.");
        }
      },
    });
  };

  const navigateGallery = (newIndex: number) => {
    if (newIndex >= 0 && newIndex < imageFiles.length) {
      openGallery(newIndex);
    }
  };

  useEffect(() => {
    if (currentImageSrc !== null && galleryIndex !== null) {
      const handleKeyDown = (e: KeyboardEvent) => {
        if (e.key === "ArrowLeft" && galleryIndex > 0) {
          navigateGallery(galleryIndex - 1);
        } else if (
          e.key === "ArrowRight" &&
          galleryIndex < imageFiles.length - 1
        ) {
          navigateGallery(galleryIndex + 1);
        } else if (e.key === "Escape") {
          setCurrentImageSrc(null);
          setGalleryIndex(null);
        }
      };
      window.addEventListener("keydown", handleKeyDown);
      return () => window.removeEventListener("keydown", handleKeyDown);
    }
  }, [currentImageSrc, galleryIndex, imageFiles.length]);

  return (
    <div className="bg-white dark:bg-gray-800 rounded-lg shadow-sm">
      {/* Header */}
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-2xl font-bold text-gray-700 dark:text-gray-200 flex items-center space-x-2">
          <span>Files</span>
          <span className="bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-gray-200 px-2 py-1 rounded text-sm">
            {task.files?.length || 0}
          </span>
        </h2>
        <button
          onClick={() => setShowFiles((prev) => !prev)}
          className="text-blue-500 hover:text-blue-700"
          title={showFiles ? "Hide files" : "Show files"}
        >
          {showFiles ? (
            <ChevronUpIcon className="w-6 h-6" />
          ) : (
            <ChevronDownIcon className="w-6 h-6" />
          )}
        </button>
      </div>

      {showFiles && (
        <>
          {/* Gallery for image files */}
          {imageFiles.length > 0 && (
            <div className="mb-4">
              <h3 className="text-xl font-semibold text-gray-700 dark:text-gray-200 mb-2">
                Image Gallery
              </h3>
              <div className="grid grid-cols-2 md:grid-cols-3 gap-2">
                {imageFiles.map((file, index) => (
                  <ImageThumbnail
                    key={file.id}
                    file={file}
                    onClick={() => openGallery(index)}
                  />
                ))}
              </div>
            </div>
          )}

          {/* List for non-image files */}
          {nonImageFiles.length > 0 && (
            <div className="mb-4">
              <h3 className="text-xl font-semibold text-gray-700 dark:text-gray-200 mb-2">
                Other Files
              </h3>
              <div className="space-y-2">
                {nonImageFiles.map((file) => (
                  <div
                    key={file.id}
                    className="flex items-center justify-between p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-gray-50 dark:bg-gray-700"
                  >
                    <span className="text-sm text-gray-700 dark:text-gray-200">
                      {file.name}
                    </span>
                    <div className="flex items-center space-x-2">
                      <button
                        onClick={() => handleDownload(file.id ?? 0)}
                        disabled={isDownloading}
                        title={isDownloading ? "Downloading..." : "Download"}
                        className="p-2 rounded-md text-white bg-blue-500 hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed"
                      >
                        <ArrowDownTrayIcon className="w-5 h-5" />
                      </button>
                      <button
                        onClick={() => handleDelete(file)}
                        disabled={isDeleting}
                        title={isDeleting ? "Deleting..." : "Delete"}
                        className="p-2 rounded-md text-white bg-red-500 hover:bg-red-600 dark:bg-red-600 dark:hover:bg-red-700 disabled:bg-gray-400 disabled:cursor-not-allowed"
                      >
                        <TrashIcon className="w-5 h-5" />
                      </button>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}

          <UploadFile taskId={task.id ?? 0} />
        </>
      )}

      {/* Gallery Modal */}
      {currentImageSrc && galleryIndex !== null && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-75 z-50">
          <div className="relative max-w-4xl w-full p-4">
            {/* Close Button */}
            <button
              onClick={() => {
                setCurrentImageSrc(null);
                setGalleryIndex(null);
              }}
              className="absolute top-2 right-2 text-white"
              title="Close"
            >
              <XMarkIcon className="w-8 h-8" />
            </button>
            {/* Delete Button */}
            <button
              onClick={() => {
                handleDelete(imageFiles[galleryIndex]);
              }}
              className="absolute top-2 left-2 text-white"
              title="Delete image"
            >
              <TrashIcon className="w-8 h-8" />
            </button>
            {/* Left navigation arrow */}
            {galleryIndex > 0 && (
              <button
                onClick={() => navigateGallery(galleryIndex - 1)}
                className="absolute top-1/2 left-2 transform -translate-y-1/2 text-white bg-gray-800 bg-opacity-50 p-2 rounded-full"
                title="Previous image"
              >
                <ChevronLeftIcon className="w-6 h-6" />
              </button>
            )}
            {/* Right navigation arrow */}
            {galleryIndex < imageFiles.length - 1 && (
              <button
                onClick={() => navigateGallery(galleryIndex + 1)}
                className="absolute top-1/2 right-2 transform -translate-y-1/2 text-white bg-gray-800 bg-opacity-50 p-2 rounded-full"
                title="Next image"
              >
                <ChevronRightIcon className="w-6 h-6" />
              </button>
            )}
            <img
              src={currentImageSrc}
              alt={imageFiles[galleryIndex].name}
              className="max-h-[90vh] max-w-full mx-auto rounded-md shadow-lg"
            />
            {/* Image file name caption */}
            <p className="mt-4 text-center text-white text-lg">
              {imageFiles[galleryIndex].name}
            </p>
          </div>
        </div>
      )}
    </div>
  );
};
