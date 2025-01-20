import React, { useState } from "react";
import { toast } from "react-toastify";
import { useAddOrInvite } from "../../../hooks/boards/useAddOrInvite";
import { BoardAddOrInviteUserToBoardRequestRoleEnum } from "../../../typescript-fetch-client";

export const BoardAddOrInvite = ({ boardId }: { boardId: number }) => {
  const { mutate, isError, isSuccess, isPending } = useAddOrInvite();
  const [email, setEmail] = useState("");
  const [role, setRole] = useState("member");

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!email) {
      toast.error("Please enter an email address.");
      return;
    }

    mutate({
      request: {
        boardId,
        email,
        role: role as BoardAddOrInviteUserToBoardRequestRoleEnum,
      },
    });
  };

  return (
    <div className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-sm dark:shadow-md space-y-6">
      <h3 className="text-2xl font-semibold text-gray-900 dark:text-gray-100">
        Invite User to Board
      </h3>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="flex flex-col gap-2">
          <label
            htmlFor="email"
            className="text-gray-700 dark:text-gray-300 font-medium"
          >
            Email:
          </label>
          <input
            type="email"
            id="email"
            className="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>

        <div className="flex flex-col gap-2">
          <label
            htmlFor="role"
            className="text-gray-700 dark:text-gray-300 font-medium"
          >
            Role:
          </label>
          <select
            id="role"
            className="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            value={role}
            onChange={(e) => setRole(e.target.value)}
          >
            <option value="admin">Admin</option>
            <option value="member">Member</option>
            <option value="reader">Reader</option>
          </select>
        </div>

        <button
          type="submit"
          className={`px-4 py-2 text-white rounded-lg transition-colors ${
            isPending
              ? "bg-gray-400 dark:bg-gray-600 cursor-not-allowed"
              : "bg-blue-600 dark:bg-blue-700 hover:bg-blue-700 dark:hover:bg-blue-800"
          }`}
          disabled={isPending}
        >
          {isPending ? "Sending Invite..." : "Send Invite"}
        </button>
      </form>

      {isError && (
        <p className="text-red-600 dark:text-red-400">
          Failed to send invite. Please try again.
        </p>
      )}
      {isSuccess && (
        <p className="text-green-600 dark:text-green-400">
          Invite sent successfully!
        </p>
      )}
    </div>
  );
};
