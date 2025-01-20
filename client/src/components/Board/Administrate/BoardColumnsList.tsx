import { ModelsBoard } from "../../../typescript-fetch-client";
import { BoardColumnDelete } from "./BoardColumnDelete";
import { useDraggableColumns } from "../../../hooks/boards/administration/useDraggableColumns";

export const BoardColumnsList = ({ board }: { board: ModelsBoard }) => {
  const { columns, onDragStart, onDragOver, onDrop } = useDraggableColumns(
    board.columns ?? []
  );

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
            <span className="font-medium text-gray-700 dark:text-gray-200">
              {column.name}
            </span>

            <BoardColumnDelete column={column} />
          </li>
        ))}
      </ul>
    </div>
  );
};
