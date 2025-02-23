import { useState } from "react";
import { ModelsBoard } from "../../../typescript-fetch-client";
import { BoardSwimlaneDelete } from "./BoardSwimlaneDelete";
import { useDraggableSwimlanes } from "../../../hooks/boards/administration/useDraggableSwimlanes";
import { useRenameSwimlane } from "../../../hooks/boards/administration/useRenameSwimlane";
import { PencilIcon } from "@heroicons/react/24/outline";

export const BoardSwimlaneList = ({ board }: { board: ModelsBoard }) => {
  const { swimlanes, onDragStart, onDragOver, onDrop } = useDraggableSwimlanes(
    board.swimlanes ?? []
  );
  const { mutate: renameSwimlane } = useRenameSwimlane();
  const [editingSwimlaneId, setEditingSwimlaneId] = useState<number | null>(
    null
  );
  const [renameInput, setRenameInput] = useState<string>("");

  const handleRename = (swimlaneId: number) => {
    if (renameInput.trim()) {
      renameSwimlane({ id: swimlaneId, name: renameInput });
    }
    setEditingSwimlaneId(null);
  };

  return (
    <div className="border border-gray-200 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 shadow-sm">
      <ul className="divide-y divide-gray-200 dark:divide-gray-700">
        {swimlanes.map((swimlane) => (
          <li
            key={swimlane.id}
            className="flex items-center justify-between p-3 cursor-move hover:bg-gray-50 dark:hover:bg-gray-700 transition-colours"
            draggable
            onDragStart={(e) => onDragStart(e, swimlane.id!)}
            onDragOver={onDragOver}
            onDrop={(e) => onDrop(e, swimlane.id!)}
          >
            {editingSwimlaneId === swimlane.id ? (
              <input
                type="text"
                className="bg-gray-100 dark:bg-gray-600 rounded p-1 text-black dark:text-white"
                value={renameInput}
                onChange={(e) => setRenameInput(e.target.value)}
                onBlur={() => handleRename(swimlane.id!)}
                onKeyDown={(e) => {
                  if (e.key === "Enter") {
                    handleRename(swimlane.id!);
                  }
                }}
                autoFocus
              />
            ) : (
              <div className="flex items-center">
                <span className="font-medium text-gray-700 dark:text-gray-200">
                  {swimlane.name}
                </span>
                <button
                  onClick={() => {
                    setEditingSwimlaneId(swimlane.id!);
                    setRenameInput(swimlane.name as string);
                  }}
                  className="ml-2 text-sm text-blue-600 dark:text-blue-400"
                >
                  <PencilIcon className="w-4 h-4 text-yellow-500 dark:text-yellow-400" />
                </button>
              </div>
            )}
            <BoardSwimlaneDelete swimlane={swimlane} />
          </li>
        ))}
      </ul>
    </div>
  );
};
