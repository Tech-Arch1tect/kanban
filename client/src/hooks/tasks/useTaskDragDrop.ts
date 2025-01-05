import { useState, DragEvent, useCallback } from "react";
import {
  ModelsColumn,
  ModelsSwimlane,
  ModelsTask,
  TaskMoveTaskRequest,
} from "../../typescript-fetch-client";
import { UseMutateFunction } from "@tanstack/react-query";

interface UseTaskDragDropProps {
  column: ModelsColumn;
  swimlane: ModelsSwimlane;
  moveTask: UseMutateFunction<void, unknown, TaskMoveTaskRequest>;
  tasks: ModelsTask[];
}

export function useTaskDragDrop({
  column,
  swimlane,
  moveTask,
  tasks,
}: UseTaskDragDropProps) {
  const [draggedTaskId, setDraggedTaskId] = useState<number | null>(null);
  const [hoveredTaskId, setHoveredTaskId] = useState<number | null>(null);
  const [hoveredPosition, setHoveredPosition] = useState<number | null>(null);

  const handleDragStart = useCallback(
    (event: DragEvent<HTMLDivElement>, taskId: number) => {
      event.dataTransfer.setData("text/plain", JSON.stringify({ taskId }));
      setDraggedTaskId(taskId);
    },
    []
  );

  const handleDragOver = useCallback((event: DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    const targetElement = (event.target as HTMLElement).closest(".task");
    if (targetElement) {
      const pos = targetElement.getAttribute("data-position");
      const hoveredId = targetElement.getAttribute("data-task-id");

      if (pos !== null && hoveredId !== null) {
        setHoveredTaskId(Number(hoveredId));
        setHoveredPosition(Number(pos));
      }
    } else {
      setHoveredTaskId(null);
      setHoveredPosition(null);
    }
  }, []);

  const handleDrop = useCallback(
    (event: DragEvent<HTMLDivElement>) => {
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

      setDraggedTaskId(null);
      setHoveredTaskId(null);
      setHoveredPosition(null);
    },
    [column.id, swimlane.id, hoveredPosition, moveTask]
  );

  const isDraggedTask = useCallback(
    (taskId: number) => draggedTaskId === taskId,
    [draggedTaskId]
  );
  const isHoveredTask = useCallback(
    (taskId: number) => hoveredTaskId === taskId,
    [hoveredTaskId]
  );

  return {
    draggedTaskId,
    hoveredTaskId,
    hoveredPosition,
    handleDragStart,
    handleDragOver,
    handleDrop,
    isDraggedTask,
    isHoveredTask,
  };
}
