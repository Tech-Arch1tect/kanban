import { useState } from "react";
import {
  ModelsColumn,
  ModelsSwimlane,
  ModelsTask,
} from "../../typescript-fetch-client";
import { useCreateTask } from "../../hooks/tasks/useCreateTask";
import { PlusIcon } from "@heroicons/react/24/solid";
import { Task } from "../Task/Task";
import { useMoveTask } from "../../hooks/tasks/useMoveTask";

export default function ColumnTasks({
  column,
  swimlane,
  tasks,
}: {
  column: ModelsColumn;
  swimlane: ModelsSwimlane;
  tasks: ModelsTask[];
}) {
  const { mutate: createTask } = useCreateTask();
  const { mutate: moveTask } = useMoveTask();

  const [isFormVisible, setFormVisible] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  const [hoveredPosition, setHoveredPosition] = useState<number | null>(null);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    createTask({
      title,
      description,
      status: "open",
      swimlaneId: swimlane.id,
      columnId: column.id,
      boardId: swimlane.boardId,
    });
    setTitle("");
    setDescription("");
    setFormVisible(false);
  };

  const handleDragOver = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    const targetElement = (event.target as HTMLElement).closest(".task");
    if (targetElement) {
      const pos = targetElement.getAttribute("data-position");
      if (pos !== null) {
        setHoveredPosition(Number(pos));
      }
    } else {
      setHoveredPosition(null);
    }
  };

  const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    const data = JSON.parse(event.dataTransfer.getData("text/plain"));
    const { taskId } = data;

    if (!taskId || !column.id || !swimlane.id) return;

    moveTask({
      taskId,
      columnId: column.id,
      swimlaneId: swimlane.id,
      position: hoveredPosition ?? 0,
    });

    setHoveredPosition(null);
  };

  return (
    <div
      className="bg-gray-100 rounded shadow min-h-[10px]"
      onDrop={handleDrop}
      onDragOver={handleDragOver}
    >
      {!isFormVisible && (
        <button
          className="flex items-center justify-center w-full p-2 bg-gray-200 rounded"
          onClick={() => setFormVisible(true)}
          aria-label="Add Task"
        >
          <PlusIcon className="w-6 h-6" />
        </button>
      )}

      {isFormVisible && (
        <form onSubmit={handleSubmit} className="flex flex-col space-y-2 mt-4">
          <input
            className="p-2 border rounded"
            placeholder="Task title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
          <textarea
            className="p-2 border rounded"
            placeholder="Task description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
          <div className="flex space-x-2">
            <button
              className="bg-blue-500 text-white p-2 rounded"
              type="submit"
            >
              Create
            </button>
            <button
              className="bg-gray-300 p-2 rounded"
              type="button"
              onClick={() => setFormVisible(false)}
            >
              Cancel
            </button>
          </div>
        </form>
      )}

      <div className="space-y-2">
        {tasks
          .filter(
            (task) =>
              task.swimlaneId === swimlane.id && task.columnId === column.id
          )
          .map((task) => (
            <Task key={task.id} task={task} />
          ))}
      </div>
    </div>
  );
}
