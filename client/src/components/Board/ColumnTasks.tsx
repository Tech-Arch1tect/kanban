import { useState, useEffect, useRef } from "react";
import {
  ModelsColumn,
  ModelsSwimlane,
  ModelsTask,
} from "../../typescript-fetch-client";
import { useCreateTask } from "../../hooks/tasks/useCreateTask";
import { PlusIcon } from "@heroicons/react/24/solid";
import { Task } from "./Task/Task";
import { useMoveTask } from "../../hooks/tasks/useMoveTask";
import { useTaskDragDrop } from "../../hooks/tasks/useTaskDragDrop";

const TASKS_CHUNK_SIZE = 100;

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
  const [visibleTaskCount, setVisibleTaskCount] = useState(TASKS_CHUNK_SIZE);

  const loaderRef = useRef<HTMLDivElement | null>(null);

  const { handleDragOver, handleDrop, handleDragLeave, ...dragStates } =
    useTaskDragDrop({
      column,
      swimlane,
      moveTask,
      tasks,
    });

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

  const columnTasks = tasks.filter(
    (task) => task.swimlaneId === swimlane.id && task.columnId === column.id
  );

  const visibleTasks = columnTasks.slice(0, visibleTaskCount);

  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting) {
          setVisibleTaskCount((prev) =>
            Math.min(prev + TASKS_CHUNK_SIZE, columnTasks.length)
          );
        }
      },
      { threshold: 1.0 }
    );

    if (loaderRef.current) {
      observer.observe(loaderRef.current);
    }

    return () => {
      if (loaderRef.current) {
        observer.unobserve(loaderRef.current);
      }
    };
  }, [columnTasks.length]);

  return (
    <div
      className="bg-gray-100 rounded shadow min-h-[10px]"
      onDrop={handleDrop}
      onDragOver={handleDragOver}
      onDragLeave={handleDragLeave}
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
        <form
          onSubmit={handleSubmit}
          className="flex flex-col space-y-2 mt-4 p-2"
        >
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
        {visibleTasks.map((task) => (
          <Task key={task.id} task={task} {...dragStates} />
        ))}
        {visibleTaskCount < columnTasks.length && (
          <div
            ref={loaderRef}
            className="h-10 flex items-center justify-center text-gray-500"
          >
            Loading more tasks...
          </div>
        )}
      </div>
    </div>
  );
}
