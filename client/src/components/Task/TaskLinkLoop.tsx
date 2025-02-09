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
    <>
      {links.map((link) => (
        <li
          key={link.id}
          className="flex items-center justify-between p-2 border border-gray-200 dark:border-gray-600 rounded bg-white dark:bg-gray-700"
        >
          <Link
            // @ts-ignore
            to={`/task/${link.dstTaskId}`}
            className="text-sm font-medium text-blue-500 hover:text-blue-700 dark:text-blue-400 dark:hover:text-blue-300"
          >
            <span>
              {linkTypeDisplay[link.linkType ?? ""] || link.linkType}:{" "}
              {link.dstTask?.title}
            </span>
          </Link>
          <button
            onClick={() => handleDelete(link.id!)}
            className="text-red-500 hover:text-red-700 dark:text-red-400 dark:hover:text-red-300"
            disabled={isDeleting}
            title="Delete link"
          >
            <TrashIcon className="w-4 h-4" />
          </button>
        </li>
      ))}
    </>
  );
};
