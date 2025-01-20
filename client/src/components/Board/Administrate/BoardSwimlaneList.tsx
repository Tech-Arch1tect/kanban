import { ModelsBoard } from "../../../typescript-fetch-client";
import { BoardSwimlaneDelete } from "./BoardSwimlaneDelete";
import { useDraggableSwimlanes } from "../../../hooks/boards/administration/useDraggableSwimlanes";

export const BoardSwimlaneList = ({ board }: { board: ModelsBoard }) => {
  const { swimlanes, onDragStart, onDragOver, onDrop } = useDraggableSwimlanes(
    board.swimlanes ?? []
  );

  return (
    <div className="border border-gray-200 dark:border-gray-700 rounded-md bg-white dark:bg-gray-800 shadow-sm">
      <ul className="divide-y divide-gray-200 dark:divide-gray-700">
        {swimlanes.map((swimlane) => (
          <li
            key={swimlane.id}
            className="flex items-center justify-between p-3 cursor-move hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
            draggable
            onDragStart={(e) => onDragStart(e, swimlane.id!)}
            onDragOver={onDragOver}
            onDrop={(e) => onDrop(e, swimlane.id!)}
          >
            <span className="font-medium text-gray-700 dark:text-gray-200">
              {swimlane.name}
            </span>

            <BoardSwimlaneDelete swimlane={swimlane} />
          </li>
        ))}
      </ul>
    </div>
  );
};
