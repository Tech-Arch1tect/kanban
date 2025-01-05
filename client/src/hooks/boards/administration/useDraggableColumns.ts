import { useState, useCallback, useEffect } from "react";
import { ModelsColumn } from "../../../typescript-fetch-client";
import { useMoveColumn } from "./useMoveColumn";

type UseDraggableColumnsReturn = {
  columns: ModelsColumn[];
  onDragStart: (e: React.DragEvent<HTMLLIElement>, columnId: number) => void;
  onDragOver: (e: React.DragEvent<HTMLLIElement>) => void;
  onDrop: (e: React.DragEvent<HTMLLIElement>, targetColumnId: number) => void;
  setColumns: React.Dispatch<React.SetStateAction<ModelsColumn[]>>;
};

export const useDraggableColumns = (
  initialColumns: ModelsColumn[]
): UseDraggableColumnsReturn => {
  const [columns, setColumns] = useState(initialColumns);
  const [draggedColumnId, setDraggedColumnId] = useState<number | null>(null);
  const { mutate: moveColumn } = useMoveColumn();

  useEffect(() => {
    setColumns(initialColumns);
  }, [initialColumns]);

  const onDragStart = useCallback(
    (e: React.DragEvent<HTMLLIElement>, columnId: number) => {
      setDraggedColumnId(columnId);
      e.dataTransfer.effectAllowed = "move";
    },
    []
  );

  const onDragOver = useCallback((e: React.DragEvent<HTMLLIElement>) => {
    e.preventDefault();
  }, []);

  const onDrop = useCallback(
    (e: React.DragEvent<HTMLLIElement>, targetColumnId: number) => {
      e.preventDefault();
      if (draggedColumnId == null) return;

      const draggedIndex = columns.findIndex((c) => c.id === draggedColumnId);
      const targetIndex = columns.findIndex((c) => c.id === targetColumnId);
      if (draggedIndex === targetIndex) return;

      const updatedColumns = [...columns];
      const [draggedItem] = updatedColumns.splice(draggedIndex, 1);
      updatedColumns.splice(targetIndex, 0, draggedItem);
      setColumns(updatedColumns);

      const direction = draggedIndex < targetIndex ? "before" : "after";
      moveColumn({
        id: draggedItem.id!,
        relativeId: targetColumnId,
        direction: direction,
      });

      setDraggedColumnId(null);
    },
    [draggedColumnId, columns, moveColumn]
  );

  return {
    columns,
    onDragStart,
    onDragOver,
    onDrop,
    setColumns,
  };
};
