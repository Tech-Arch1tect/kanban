import React, { useState } from "react";
import { useUploadFile } from "../../../hooks/tasks/Files/useUploadFile";
import { CloudArrowUpIcon, ArrowPathIcon } from "@heroicons/react/24/outline";

export const UploadFile = ({ taskId }: { taskId: number }) => {
  const [file, setFile] = useState<File | null>(null);
  const [fileName, setFileName] = useState("");
  const { mutate, isError, isSuccess, error, isPending } = useUploadFile();

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const selectedFile = e.target.files[0];
      setFile(selectedFile);
      setFileName(selectedFile.name);
    }
  };

  const handleUpload = () => {
    if (!file || !taskId) {
      alert("Please select a file and ensure a task ID is provided.");
      return;
    }

    const reader = new FileReader();
    reader.onload = (event) => {
      const fileContent = new Uint8Array(event.target?.result as ArrayBuffer);
      const fileContentArray = Array.from(fileContent);

      mutate({
        file: fileContentArray,
        name: fileName,
        taskId,
      });
    };

    reader.readAsArrayBuffer(file);
  };

  return (
    <div className="mt-4">
      {/* Header with upload icon inline */}
      <div className="flex items-center justify-between mb-3">
        <h2 className="text-2xl font-bold text-gray-700 dark:text-gray-200">
          Upload a File
        </h2>
        <label
          htmlFor={`file-upload-${taskId}`}
          className="cursor-pointer p-2 rounded-md bg-blue-500 hover:bg-blue-600 text-white"
          title="Select file"
        >
          <CloudArrowUpIcon className="w-5 h-5" />
        </label>
        <input
          id={`file-upload-${taskId}`}
          type="file"
          onChange={handleFileChange}
          className="hidden"
        />
      </div>
      <div className="flex flex-col items-start space-y-2">
        {fileName && (
          <span className="text-sm text-gray-700 dark:text-gray-200">
            {fileName}
          </span>
        )}
        {file && (
          <button
            onClick={handleUpload}
            disabled={isPending}
            title={isPending ? "Uploading..." : "Upload"}
            className="p-2 rounded-md text-white bg-green-500 hover:bg-green-600 disabled:bg-gray-400 disabled:cursor-not-allowed"
          >
            {isPending ? (
              <ArrowPathIcon className="w-5 h-5 animate-spin" />
            ) : (
              <CloudArrowUpIcon className="w-5 h-5" />
            )}
          </button>
        )}
        {isSuccess && (
          <p className="text-sm text-green-600 dark:text-green-400">
            File uploaded successfully!
          </p>
        )}
        {isError && (
          <p className="text-sm text-red-600 dark:text-red-400">
            Error: {error?.message || "Failed to upload file."}
          </p>
        )}
      </div>
    </div>
  );
};
