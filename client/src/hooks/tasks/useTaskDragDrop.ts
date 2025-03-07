import { useState, useCallback } from "react";
import {
  ModelsColumn,
  ModelsSwimlane,
  ServerApiControllersTaskMoveTaskRequest,
} from "../../typescript-fetch-client";
import { UseMutateFunction } from "@tanstack/react-query";
import useDebounce from "../useDebounce";

interface UseTaskDragDropProps {
  column: ModelsColumn;
  swimlane: ModelsSwimlane;
  moveTask: UseMutateFunction<
    void,
    unknown,
    ServerApiControllersTaskMoveTaskRequest
  >;
}

export function useTaskDragDrop({
  column,
  swimlane,
  moveTask,
}: UseTaskDragDropProps) {
  const [draggedTaskId, setDraggedTaskId] = useState<number | null>(null);
  const [rawHoveredTaskId, setRawHoveredTaskId] = useState<number | null>(null);
  const [rawHoveredPosition, setRawHoveredPosition] = useState<number | null>(
    null
  );

  const hoveredTaskId = useDebounce(rawHoveredTaskId, 5);
  const hoveredPosition = useDebounce(rawHoveredPosition, 5);

  const handleDragStart = useCallback(
    (event: React.DragEvent<HTMLDivElement>, taskId: number) => {
      event.dataTransfer.setData("text/plain", JSON.stringify({ taskId }));
      setDraggedTaskId(taskId);
    },
    []
  );

  const handleDragOver = useCallback(
    (event: React.DragEvent<HTMLDivElement>) => {
      event.preventDefault();
      const targetElement = (event.target as HTMLElement).closest(".task");
      if (targetElement) {
        const pos = targetElement.getAttribute("data-position");
        const hoveredId = targetElement.getAttribute("data-task-id");

        if (pos !== null && hoveredId !== null) {
          setRawHoveredTaskId(Number(hoveredId));
          setRawHoveredPosition(Number(pos));
        }
      } else {
        setRawHoveredTaskId(null);
        setRawHoveredPosition(null);
      }
    },
    []
  );

  const handleDragLeave = useCallback(() => {
    setRawHoveredTaskId(null);
    setRawHoveredPosition(null);
  }, []);

  const handleDrop = useCallback(
    (event: React.DragEvent<HTMLDivElement>) => {
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
      setRawHoveredTaskId(null);
      setRawHoveredPosition(null);
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
    handleDragLeave,
    handleDrop,
    isDraggedTask,
    isHoveredTask,
  };
}
