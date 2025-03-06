import { useState } from "react";
import {
  ModelsTask,
  TaskUpdateTaskColourRequestColourEnum,
} from "../../typescript-fetch-client";
import { useUpdateTaskColour } from "../../hooks/tasks/useUpdateTaskColour";

const tailwindColours = [
  "slate",
  "gray",
  "zinc",
  "neutral",
  "stone",
  "red",
  "orange",
  "amber",
  "yellow",
  "lime",
  "green",
  "emerald",
  "teal",
  "cyan",
  "sky",
  "blue",
  "indigo",
  "violet",
  "purple",
  "fuchsia",
  "pink",
  "rose",
];

export const TaskColourPicker = ({ task }: { task: ModelsTask }) => {
  const { mutate } = useUpdateTaskColour();
  const [selectedColour, setSelectedColour] = useState(task.colour);

  const handleColourChange = (colour: string) => {
    setSelectedColour(colour as TaskUpdateTaskColourRequestColourEnum);
    mutate({
      taskId: task.id as number,
      colour: colour as TaskUpdateTaskColourRequestColourEnum,
    });
  };

  return (
    <div className="flex flex-wrap gap-2">
      {tailwindColours.map((colour) => {
        const baseClass = "w-8 h-8 rounded-full border-2";
        const activeClass =
          selectedColour === colour
            ? "border-black dark:border-white"
            : "border-transparent";
        const lightModeClass = `bg-${colour}-500`;
        const darkModeClass = `dark:bg-${colour}-700`;

        return (
          <button
            key={colour}
            onClick={() => handleColourChange(colour)}
            className={`${baseClass} ${activeClass} ${lightModeClass} ${darkModeClass}`}
            title={colour}
          />
        );
      })}
    </div>
  );
};
