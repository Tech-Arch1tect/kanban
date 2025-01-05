import { useState, useCallback, useEffect } from "react";
import { ModelsSwimlane } from "../../../typescript-fetch-client";
import { useMoveSwimlane } from "./useMoveSwimlane";

type UseDraggableSwimlanesReturn = {
  swimlanes: ModelsSwimlane[];
  onDragStart: (e: React.DragEvent<HTMLLIElement>, swimlaneId: number) => void;
  onDragOver: (e: React.DragEvent<HTMLLIElement>) => void;
  onDrop: (e: React.DragEvent<HTMLLIElement>, targetSwimlaneId: number) => void;
  setSwimlanes: React.Dispatch<React.SetStateAction<ModelsSwimlane[]>>;
};

export const useDraggableSwimlanes = (
  initialSwimlanes: ModelsSwimlane[]
): UseDraggableSwimlanesReturn => {
  const [swimlanes, setSwimlanes] = useState(initialSwimlanes);
  const [draggedSwimlaneId, setDraggedSwimlaneId] = useState<number | null>(
    null
  );
  const { mutate: moveSwimlane } = useMoveSwimlane();

  useEffect(() => {
    setSwimlanes(initialSwimlanes);
  }, [initialSwimlanes]);

  const onDragStart = useCallback(
    (e: React.DragEvent<HTMLLIElement>, swimlaneId: number) => {
      setDraggedSwimlaneId(swimlaneId);
      e.dataTransfer.effectAllowed = "move";
    },
    []
  );

  const onDragOver = useCallback((e: React.DragEvent<HTMLLIElement>) => {
    e.preventDefault();
  }, []);

  const onDrop = useCallback(
    (e: React.DragEvent<HTMLLIElement>, targetSwimlaneId: number) => {
      e.preventDefault();
      if (draggedSwimlaneId == null) return;

      const draggedIndex = swimlanes.findIndex(
        (s) => s.id === draggedSwimlaneId
      );
      const targetIndex = swimlanes.findIndex((s) => s.id === targetSwimlaneId);
      if (draggedIndex === targetIndex) return;

      const updatedSwimlanes = [...swimlanes];
      const [draggedItem] = updatedSwimlanes.splice(draggedIndex, 1);
      updatedSwimlanes.splice(targetIndex, 0, draggedItem);
      setSwimlanes(updatedSwimlanes);

      const direction = draggedIndex < targetIndex ? "before" : "after";
      moveSwimlane({
        id: draggedItem.id!,
        relativeId: targetSwimlaneId,
        direction: direction,
      });

      setDraggedSwimlaneId(null);
    },
    [draggedSwimlaneId, swimlanes, moveSwimlane]
  );

  return {
    swimlanes,
    onDragStart,
    onDragOver,
    onDrop,
    setSwimlanes,
  };
};
