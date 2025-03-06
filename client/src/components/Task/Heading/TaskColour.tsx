import { useState } from "react";
import { ModelsTask } from "../../../typescript-fetch-client";
import { TaskColourPicker } from "../../Utility/TaskColourPicker";
import { colourToClass } from "../../Utility/colourToClass";

export const TaskColour = ({ task }: { task: ModelsTask }) => {
  const [showPicker, setShowPicker] = useState(false);
  const togglePicker = () => setShowPicker((prev) => !prev);

  const badgeClass =
    colourToClass[task.colour as keyof typeof colourToClass] ||
    "bg-gray-500 dark:bg-gray-700";

  return (
    <div className="relative inline-block">
      <span className="font-medium text-gray-600 dark:text-gray-400 w-24 inline-block">
        Colour:
      </span>
      <div
        className={`w-6 h-6 rounded-full ${badgeClass} cursor-pointer border-2 border-white dark:border-gray-800 inline-block`}
        onClick={togglePicker}
      />

      {showPicker && (
        <div className="absolute z-10 mt-2">
          <TaskColourPicker task={task} />
        </div>
      )}
    </div>
  );
};
