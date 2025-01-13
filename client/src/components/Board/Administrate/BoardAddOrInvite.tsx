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
    <div className="bg-white p-6 rounded space-y-6">
      <h3 className="text-2xl font-semibold">Invite User to Board</h3>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="flex flex-col gap-2">
          <label htmlFor="email" className="text-gray-700 font-medium">
            Email:
          </label>
          <input
            type="email"
            id="email"
            className="px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>

        <div className="flex flex-col gap-2">
          <label htmlFor="role" className="text-gray-700 font-medium">
            Role:
          </label>
          <select
            id="role"
            className="px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
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
          className={`px-4 py-2 text-white rounded ${
            isPending
              ? "bg-gray-400 cursor-not-allowed"
              : "bg-blue-600 hover:bg-blue-700"
          }`}
          disabled={isPending}
        >
          {isPending ? "Sending Invite..." : "Send Invite"}
        </button>
      </form>

      {isError && (
        <p className="text-red-600">Failed to send invite. Please try again.</p>
      )}
      {isSuccess && <p className="text-green-600">Invite sent successfully!</p>}
    </div>
  );
};
