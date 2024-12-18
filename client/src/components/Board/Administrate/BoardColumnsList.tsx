import { ModelsBoard } from "../../../typescript-fetch-client";
import { useDeleteColumn } from "../../../hooks/boards/administration/useDeleteColumn";
import { BoardColumnDelete } from "./BoardColumnDelete";
import { useDraggableColumns } from "../../../hooks/boards/administration/useDraggableColumns";

export const BoardColumnsList = ({ board }: { board: ModelsBoard }) => {
  const { columns, onDragStart, onDragOver, onDrop } = useDraggableColumns(
    board.columns ?? []
  );
  return (
    <div>
      <h2 className="text-xl font-bold">Columns</h2>
      <ul>
        {columns.map((column) => (
          <li
            key={column.id}
            className="flex items-center gap-2"
            draggable
            onDragStart={(e) => onDragStart(e, column.id!)}
            onDragOver={onDragOver}
            onDrop={(e) => onDrop(e, column.id!)}
          >
            <span>{column.name}</span>
            <BoardColumnDelete column={column} />
          </li>
        ))}
      </ul>
    </div>
  );
};
