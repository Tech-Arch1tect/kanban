import { useState } from "react";
import {
  ModelsColumn,
  ModelsSwimlane,
  ModelsTask,
} from "../../typescript-fetch-client";
import { useCreateTask } from "../../hooks/tasks/useCreateTask";
import { PlusIcon } from "@heroicons/react/24/solid";
import { Task } from "../Task/Task";

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
  const [isFormVisible, setFormVisible] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

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

  return (
    <div className="bg-gray-100 rounded shadow min-h-[10px]">
      {!isFormVisible && (
        <button
          className="flex items-center justify-center w-4 h-4 bg-blue-500 text-white rounded hover:bg-blue-600 transition mt-[-10px] float-right"
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
