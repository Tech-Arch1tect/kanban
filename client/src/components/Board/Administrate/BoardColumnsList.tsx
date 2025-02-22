import { ModelsBoard } from "../../../typescript-fetch-client";
import { useState } from "react";
import { useRenameColumn } from "../../../hooks/boards/administration/useRenameColumn";
import { BoardColumnDelete } from "./BoardColumnDelete";
import { useDraggableColumns } from "../../../hooks/boards/administration/useDraggableColumns";
import { PencilIcon } from "@heroicons/react/24/outline";

export const BoardColumnsList = ({ board }: { board: ModelsBoard }) => {
  const { columns, onDragStart, onDragOver, onDrop } = useDraggableColumns(
    board.columns ?? []
  );
  const { mutate: renameColumn } = useRenameColumn();
  const [editingColumnId, setEditingColumnId] = useState<number | null>(null);
  const [renameInput, setRenameInput] = useState<string>("");

  const handleRename = (columnId: number) => {
    if (renameInput.trim()) {
      renameColumn({ id: columnId, name: renameInput });
    }
    setEditingColumnId(null);
  };

  return (
    <div className="border border-gray-200 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 shadow-sm">
      <ul className="divide-y divide-gray-200 dark:divide-gray-700">
        {columns.map((column) => (
          <li
            key={column.id}
            className="flex items-center justify-between p-3 cursor-move hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
            draggable
            onDragStart={(e) => onDragStart(e, column.id!)}
            onDragOver={onDragOver}
            onDrop={(e) => onDrop(e, column.id!)}
          >
            {editingColumnId === column.id ? (
              <input
                type="text"
                className="bg-gray-100 dark:bg-gray-600 rounded p-1 dark:text-white"
                value={renameInput}
                onChange={(e) => setRenameInput(e.target.value)}
                onBlur={() => handleRename(column.id!)}
                onKeyDown={(e) => {
                  if (e.key === "Enter") {
                    handleRename(column.id!);
                  }
                }}
                autoFocus
              />
            ) : (
              <div className="flex items-center">
                <span className="font-medium text-gray-700 dark:text-gray-200">
                  {column.name}
                </span>
                <button
                  onClick={() => {
                    setEditingColumnId(column.id!);
                    setRenameInput(column.name as string);
                  }}
                  className="ml-2 text-sm text-blue-600 dark:text-blue-400"
                >
                  <PencilIcon className="w-4 h-4 text-yellow-500 dark:text-yellow-400" />
                </button>
              </div>
            )}
            <BoardColumnDelete column={column} />
          </li>
        ))}
      </ul>
    </div>
  );
};
