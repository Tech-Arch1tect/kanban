import { useEffect, useState } from "react";
import { useGetImage } from "../../../hooks/tasks/Files/useGetImage";
import { ModelsFile } from "../../../typescript-fetch-client";

export const ImageThumbnail = ({
  file,
  onClick,
}: {
  file: ModelsFile;
  onClick: () => void;
}) => {
  const [thumbnailSrc, setThumbnailSrc] = useState<string | null>(null);
  const { mutate: getImage } = useGetImage();

  useEffect(() => {
    getImage(file.id as number, {
      onSuccess: (data) => {
        if (data.content) {
          setThumbnailSrc(data.content);
        }
      },
    });
  }, [file.id, getImage]);

  return (
    <div
      onClick={onClick}
      className="cursor-pointer rounded-md overflow-hidden relative"
    >
      {thumbnailSrc ? (
        <img
          src={thumbnailSrc}
          alt={file.name}
          className="object-cover w-full h-32"
        />
      ) : (
        <div className="w-full h-32 flex items-center justify-center bg-gray-200 dark:bg-gray-700 text-gray-500">
          Loading...
        </div>
      )}
    </div>
  );
};
