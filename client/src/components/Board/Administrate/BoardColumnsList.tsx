import { ModelsBoard } from "../../../typescript-fetch-client";
import { BoardColumnDelete } from "./BoardColumnDelete";
import { useDraggableColumns } from "../../../hooks/boards/administration/useDraggableColumns";

export const BoardColumnsList = ({ board }: { board: ModelsBoard }) => {
  const { columns, onDragStart, onDragOver, onDrop } = useDraggableColumns(
    board.columns ?? []
  );

  return (
    <div className="border border-gray-200 rounded-md bg-white">
      <ul className="divide-y divide-gray-200">
        {columns.map((column) => (
          <li
            key={column.id}
            className="flex items-center justify-between p-2 cursor-move"
            draggable
            onDragStart={(e) => onDragStart(e, column.id!)}
            onDragOver={onDragOver}
            onDrop={(e) => onDrop(e, column.id!)}
          >
            <span className="font-medium text-gray-700">{column.name}</span>
            <BoardColumnDelete column={column} />
          </li>
        ))}
      </ul>
    </div>
  );
};
