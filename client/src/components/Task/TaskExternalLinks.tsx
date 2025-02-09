import { useCreateTaskExternalLink } from "../../hooks/tasks/useCreateTaskExternalLink";
import { useDeleteTaskExternalLink } from "../../hooks/tasks/useDeleteTaskExternalLink";
import { useUpdateTaskExternalLink } from "../../hooks/tasks/useUpdateTaskExternalLink";
import {
  ModelsTask,
  ModelsTaskExternalLink,
} from "../../typescript-fetch-client";
import { useState } from "react";
import {
  PencilIcon,
  CheckIcon,
  XMarkIcon,
  PlusIcon,
  TrashIcon,
} from "@heroicons/react/24/outline";

export const TaskExternalLinks = ({ task }: { task: ModelsTask }) => {
  const { mutate: createExternalLink } = useCreateTaskExternalLink();
  const { mutate: deleteExternalLink, isPending: isDeleting } =
    useDeleteTaskExternalLink();
  const { mutate: updateExternalLink, isPending: isUpdating } =
    useUpdateTaskExternalLink();

  const [isEditing, setIsEditing] = useState<number | null>(null);
  const [editedLink, setEditedLink] = useState<ModelsTaskExternalLink | null>(
    null
  );
  const [newLink, setNewLink] = useState({ title: "", url: "" });
  const [showNewLinkForm, setShowNewLinkForm] = useState(false);

  const handleCreate = () => {
    if (newLink.title && newLink.url) {
      createExternalLink({
        taskId: task.id!,
        title: newLink.title,
        url: newLink.url,
      });
      setNewLink({ title: "", url: "" });
    }
  };

  const handleEdit = (link: ModelsTaskExternalLink) => {
    setIsEditing(link.id!);
    setEditedLink({ ...link });
  };

  const handleUpdate = () => {
    if (editedLink) {
      updateExternalLink({
        id: editedLink.id!,
        title: editedLink.title,
        url: editedLink.url,
      });
      setIsEditing(null);
      setEditedLink(null);
    }
  };

  const handleDelete = (linkId: number) => {
    deleteExternalLink({ id: linkId });
  };

  return (
    <div className="space-y-4">
      {/* Header with toggle button */}
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold text-gray-700 dark:text-gray-200 flex items-center space-x-2">
          <span>Task External Links</span>
          <span className="bg-gray-200 text-gray-700 dark:bg-gray-700 dark:text-gray-200 px-2 py-1 rounded text-sm">
            {task.externalLinks?.length || 0}
          </span>
        </h2>
        <button
          onClick={() => setShowNewLinkForm((prev) => !prev)}
          className="flex-shrink-0 text-blue-500 hover:text-blue-700"
          title={showNewLinkForm ? "Close" : "Add External Link"}
        >
          {showNewLinkForm ? (
            <XMarkIcon className="w-6 h-6" />
          ) : (
            <PlusIcon className="w-6 h-6" />
          )}
        </button>
      </div>

      {/* New Link Creation Form */}
      {showNewLinkForm && (
        <div className="flex space-x-2">
          <input
            type="text"
            placeholder="Title"
            value={newLink.title}
            onChange={(e) => setNewLink({ ...newLink, title: e.target.value })}
            className="flex-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200"
          />
          <input
            type="url"
            placeholder="URL"
            value={newLink.url}
            onChange={(e) => setNewLink({ ...newLink, url: e.target.value })}
            className="flex-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200"
          />
          <button
            onClick={handleCreate}
            className="bg-blue-500 hover:bg-blue-700 text-white p-2 rounded-md flex items-center"
          >
            <PlusIcon className="w-5 h-5 mr-1" />
            Add Link
          </button>
        </div>
      )}

      {/* Existing External Links List */}
      <ul className="space-y-4">
        {(task.externalLinks?.length ?? 0) > 0 &&
          task.externalLinks?.map((link) => (
            <li
              key={link.id}
              className="flex justify-between items-center p-4 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700"
            >
              {isEditing === link.id ? (
                <div className="flex-1 space-x-2">
                  <input
                    type="text"
                    value={editedLink?.title || ""}
                    onChange={(e) =>
                      setEditedLink({ ...editedLink!, title: e.target.value })
                    }
                    className="p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200"
                  />
                  <input
                    type="url"
                    value={editedLink?.url || ""}
                    onChange={(e) =>
                      setEditedLink({ ...editedLink!, url: e.target.value })
                    }
                    className="p-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-200"
                  />
                </div>
              ) : (
                <a
                  href={link.url}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="flex-1 text-blue-500 hover:underline"
                >
                  <strong className="block">{link.title}</strong>
                  <span className="text-sm">{link.url}</span>
                </a>
              )}
              <div className="flex space-x-2">
                {isEditing === link.id ? (
                  <>
                    <button
                      onClick={handleUpdate}
                      disabled={isUpdating}
                      className="text-green-500 hover:text-green-700"
                      title="Save"
                    >
                      <CheckIcon className="w-5 h-5" />
                    </button>
                    <button
                      onClick={() => {
                        setIsEditing(null);
                        setEditedLink(null);
                      }}
                      className="text-gray-500 hover:text-gray-700"
                      title="Cancel"
                    >
                      <XMarkIcon className="w-5 h-5" />
                    </button>
                  </>
                ) : (
                  <>
                    <button
                      onClick={() => handleEdit(link)}
                      className="text-blue-500 hover:text-blue-700"
                      title="Edit"
                    >
                      <PencilIcon className="w-5 h-5" />
                    </button>
                    <button
                      onClick={() => handleDelete(link.id!)}
                      disabled={isDeleting}
                      className="text-red-500 hover:text-red-700"
                      title="Delete"
                    >
                      <TrashIcon className="w-5 h-5" />
                    </button>
                  </>
                )}
              </div>
            </li>
          ))}
      </ul>
    </div>
  );
};
