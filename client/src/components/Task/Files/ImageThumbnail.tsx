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
  const { data: image } = useGetImage(file.id as number);

  useEffect(() => {
    if (image) {
      setThumbnailSrc(image.content || null);
    }
  }, [image]);

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
