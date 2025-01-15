import { Link } from "@tanstack/react-router";
import { ModelsTaskLinks } from "../../typescript-fetch-client";
import { useDeleteTaskLink } from "../../hooks/tasks/useDeleteTaskLink";
import { TrashIcon } from "@heroicons/react/24/solid";
export const TaskLinkLoop = ({ links }: { links: ModelsTaskLinks[] }) => {
  const { mutate: deleteTaskLink, isPending: isDeleting } = useDeleteTaskLink();

  const linkTypeDisplay: { [key: string]: string } = {
    depends_on: "Depends on",
    blocks: "Blocks",
    fixes: "Fixes",
    blocked_by: "Blocked by",
    fixed_by: "Fixed by",
    depended_on_by: "Depended on by",
  };

  const handleDelete = (linkId: number) => {
    deleteTaskLink({ linkId });
  };

  return (
    <div className="mx-auto">
      <div className="space-y-4 mt-4">
        {links && links.length > 0 ? (
          links.map((link) => (
            <div
              key={link.id}
              className="flex items-start justify-between p-4 border border-gray-300 rounded-md bg-white"
            >
              <div>
                <Link
                  to={`/task/${link.dstTaskId}`}
                  className="text-sm font-medium text-blue-500 hover:text-blue-700"
                >
                  <strong className="block text-sm text-gray-700 inline-block">
                    {linkTypeDisplay[link.linkType ?? ""] || link.linkType}:
                  </strong>
                  <span className="ml-2 text-sm text-gray-700 inline-block">
                    {link.dstTask?.title}
                  </span>
                </Link>
              </div>
              <button
                onClick={() => handleDelete(link.id!)}
                className="text-red-500 hover:text-red-700 text-sm"
                disabled={isDeleting}
              >
                <TrashIcon className="w-4 h-4" />
              </button>
            </div>
          ))
        ) : (
          <div className="text-gray-500 text-center">
            No task links available.
          </div>
        )}
      </div>
    </div>
  );
};
