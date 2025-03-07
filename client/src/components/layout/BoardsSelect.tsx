import React, { useState } from "react";
import Select, {
  SingleValue,
  StylesConfig,
  CSSObjectWithLabel,
} from "react-select";
import { ModelsBoard } from "../../typescript-fetch-client";
import { useNavigate } from "@tanstack/react-router";
import { useBoards } from "../../hooks/boards/useBoards";

interface Option {
  value: string;
  label: string;
}

const BoardsSelect: React.FC = () => {
  const navigate = useNavigate();
  const { boards, error, isLoading } = useBoards();
  const [selectedOption] = useState<SingleValue<Option>>(null);

  const handleChange = (option: SingleValue<Option>) => {
    navigate({ to: `/boards/${option?.value}` });
  };

  const options: Option[] =
    boards && boards.boards
      ? boards.boards.map((board: ModelsBoard) => ({
          value: board.slug ?? "error",
          label: board.name ?? "error",
        }))
      : [];

  if (isLoading) return <p>Loading boards...</p>;
  if (error) return <p>Error loading boards: {error.message}</p>;

  const customStyles: StylesConfig<Option, false> = {
    option: (provided: CSSObjectWithLabel) => ({
      ...provided,
      "&:hover": {
        backgroundColor: "#60a5fa",
        color: "#ffffff",
      },
    }),
  };

  return (
    <Select<Option, false>
      value={selectedOption}
      onChange={handleChange}
      options={options}
      placeholder="Select a board..."
      isClearable
      className="w-96"
      styles={customStyles}
    />
  );
};

export default BoardsSelect;
