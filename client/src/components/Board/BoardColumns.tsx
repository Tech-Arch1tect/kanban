import { ModelsColumn } from "../../typescript-fetch-client";

export default function BoardColumns({ columns }: { columns: ModelsColumn[] }) {
  return (
    <div
      className="grid gap-4 mb-4"
      style={{
        gridTemplateColumns: `repeat(${columns?.length ?? 0}, minmax(200px, 1fr))`,
      }}
    >
      {columns?.map((column) => (
        <div key={column.id} className="font-semibold text-center">
          {column.name}
        </div>
      ))}
    </div>
  );
}
