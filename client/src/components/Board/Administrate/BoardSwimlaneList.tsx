import { ModelsBoard } from "../../../typescript-fetch-client";
import { BoardSwimlaneDelete } from "./BoardSwimlaneDelete";
import { useDraggableSwimlanes } from "../../../hooks/boards/administration/useDraggableSwimlanes";

export const BoardSwimlaneList = ({ board }: { board: ModelsBoard }) => {
  const { swimlanes, onDragStart, onDragOver, onDrop } = useDraggableSwimlanes(
    board.swimlanes ?? []
  );

  return (
    <div className="border border-gray-200 rounded-md bg-white">
      <ul className="divide-y divide-gray-200">
        {swimlanes.map((swimlane) => (
          <li
            key={swimlane.id}
            className="flex items-center justify-between p-2 cursor-move"
            draggable
            onDragStart={(e) => onDragStart(e, swimlane.id!)}
            onDragOver={onDragOver}
            onDrop={(e) => onDrop(e, swimlane.id!)}
          >
            <span className="font-medium text-gray-700">{swimlane.name}</span>
            <BoardSwimlaneDelete swimlane={swimlane} />
          </li>
        ))}
      </ul>
    </div>
  );
};
