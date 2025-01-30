import { ModelsColumn } from "../../typescript-fetch-client";

export default function BoardColumns({ columns }: { columns: ModelsColumn[] }) {
  return (
    <div className="mb-2">
      {/* Empty State Handling */}
      {columns.length === 0 ? (
        <p className="text-gray-500 dark:text-gray-400 text-center py-4">
          No columns available.
        </p>
      ) : (
        <div
          className="grid gap-4"
          style={{
            gridTemplateColumns: `repeat(auto-fit, minmax(200px, 1fr))`,
          }}
        >
          {columns.map((column) => (
            <div
              key={column.id}
              className="flex items-center justify-center font-semibold text-lg text-gray-700 dark:text-gray-200 bg-white dark:bg-gray-800 p-3 rounded-lg shadow-sm dark:shadow-md transition-all hover:shadow-md dark:hover:shadow-lg hover:bg-gray-100 dark:hover:bg-gray-700"
            >
              {column.name}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
