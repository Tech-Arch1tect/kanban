import { ModelsColumn } from "../../typescript-fetch-client";

export default function BoardColumns({ columns }: { columns: ModelsColumn[] }) {
  return (
    <div
      className="grid gap-4"
      style={{
        gridTemplateColumns: `repeat(${columns?.length ?? 0}, minmax(200px, 1fr))`,
      }}
    >
      {columns?.map((column) => (
        <div
          key={column.id}
          className="font-semibold text-lg text-gray-700 text-center bg-white p-1 rounded-lg shadow-sm"
        >
          {column.name}
        </div>
      ))}
    </div>
  );
}
