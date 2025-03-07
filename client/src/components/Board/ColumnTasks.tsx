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
      className="bg-white dark:bg-gray-800 rounded-lg shadow-sm dark:shadow-md p-4 min-h-[200px]"
      onDrop={handleDrop}
      onDragOver={handleDragOver}
      onDragLeave={handleDragLeave}
    >
      {!isFormVisible && (
        <button
          className="flex items-center justify-center w-full p-3 bg-blue-50 dark:bg-blue-900 text-blue-600 dark:text-blue-400 rounded-lg hover:bg-blue-100 dark:hover:bg-blue-800 transition-colors"
          onClick={() => setFormVisible(true)}
          aria-label="Add Task"
        >
          <PlusIcon className="w-6 h-6" />
          <span className="ml-2">Add Task</span>
        </button>
      )}

      {isFormVisible && (
        <form onSubmit={handleSubmit} className="flex flex-col space-y-3 mt-4">
          <input
            className="p-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            placeholder="Task title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
          <textarea
            className="p-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            placeholder="Task description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
          <div className="flex space-x-3">
            <button
              className="bg-blue-500 dark:bg-blue-600 text-white p-2 rounded-lg hover:bg-blue-600 dark:hover:bg-blue-700 transition-colors"
              type="submit"
            >
              Create
            </button>
            <button
              className="bg-gray-200 dark:bg-gray-700 text-gray-700 dark:text-gray-200 p-2 rounded-lg hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
              type="button"
              onClick={() => setFormVisible(false)}
            >
              Cancel
            </button>
          </div>
        </form>
      )}

      <div className="space-y-3 mt-4">
        {visibleTasks.map((task) => (
          <Task key={task.id} task={task} {...dragStates} />
        ))}
        {visibleTaskCount < columnTasks.length && (
          <div
            ref={loaderRef}
            className="h-10 flex items-center justify-center text-gray-500 dark:text-gray-400"
          >
            Loading more tasks...
          </div>
        )}
      </div>
    </div>
  );
}
