import React, { useState } from "react";
import { useUploadFile } from "../../../hooks/tasks/Files/useUploadFile";

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
    <div className="upload-file-container p-4 bg-gray-100 rounded-lg shadow-sm">
      <h3 className="text-md font-medium text-gray-700 mb-3">Upload a File</h3>
      <div className="upload-file-input-group flex flex-col gap-3">
        <input
          type="file"
          onChange={handleFileChange}
          className="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border file:border-gray-300 file:bg-gray-50 file:text-gray-600 hover:file:bg-gray-100"
        />
        <button
          onClick={handleUpload}
          disabled={isPending}
          className={`upload-file-button px-4 py-2 rounded-md text-white ${
            isPending
              ? "bg-gray-400 cursor-not-allowed"
              : "bg-blue-500 hover:bg-blue-600"
          }`}
        >
          {isPending ? "Uploading..." : "Upload"}
        </button>
      </div>
      {isSuccess && (
        <p className="upload-success mt-2 text-sm text-green-600">
          File uploaded successfully!
        </p>
      )}
      {isError && (
        <p className="upload-error mt-2 text-sm text-red-600">
          Error: {error?.message || "Failed to upload file."}
        </p>
      )}
    </div>
  );
};
