import { useState, FormEvent } from "react";
import { useRenameBoard } from "../../../hooks/boards/administration/useRenameBoard";
import { useUpdateBoardSlug } from "../../../hooks/boards/administration/useUpdateBoardSlug";

interface BoardRenameAndSlugUpdateProps {
  boardId: number;
  currentName: string;
  currentSlug: string;
}

const BoardRenameAndSlugUpdate: React.FC<BoardRenameAndSlugUpdateProps> = ({
  boardId,
  currentName,
  currentSlug,
}) => {
  const [name, setName] = useState(currentName);
  const [slug, setSlug] = useState(currentSlug);

  const { mutate: renameBoard, isPending: isRenaming } = useRenameBoard();
  const { mutate: updateSlug, isPending: isUpdating } = useUpdateBoardSlug();

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    renameBoard({ id: boardId, name });
    updateSlug({ id: boardId, slug });
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="bg-white dark:bg-gray-800 p-6 rounded shadow space-y-4"
    >
      <h2 className="text-2xl font-semibold text-gray-800 dark:text-gray-200">
        Update Board Details
      </h2>
      <div>
        <label className="block text-gray-700 dark:text-gray-300 mb-1">
          Board Name
        </label>
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="w-full p-2 border border-gray-300 dark:border-gray-600 rounded"
          required
        />
      </div>
      <div>
        <label className="block text-gray-700 dark:text-gray-300 mb-1">
          Board Slug
        </label>
        <input
          type="text"
          value={slug}
          onChange={(e) => setSlug(e.target.value)}
          className="w-full p-2 border border-gray-300 dark:border-gray-600 rounded"
          required
        />
      </div>
      <button
        type="submit"
        disabled={isRenaming || isUpdating}
        className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:opacity-50"
      >
        {isRenaming || isUpdating ? "Updating..." : "Update Board"}
      </button>
    </form>
  );
};

export default BoardRenameAndSlugUpdate;
