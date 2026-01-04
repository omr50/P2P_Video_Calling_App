import React from "react";

interface User {
  id: string;
  name: string;
  email: string;
  online: boolean;
}

interface UserCardProps {
  user: User;
}

const UserCard: React.FC<UserCardProps> = ({ user }) => {
  const initial = user.name.charAt(0).toUpperCase();

  return (
    <div className="flex items-center gap-4 p-3 hover:bg-gray-600 rounded-lg cursor-pointer">
      {/* Avatar */}
      <div className="relative">
        <div className="w-12 h-12 rounded-full bg-blue-600 flex items-center justify-center text-white font-semibold text-lg">
          {initial}
        </div>

        {/* Status Indicator */}
        <span
          className={`absolute bottom-0 right-0 w-3.5 h-3.5 rounded-full border-2 border-white ${
            user.online ? "bg-green-500" : "bg-gray-400"
          }`}
        />
      </div>

      {/* User Info */}
      <div className="flex flex-col">
        <span className="font-medium text-orange-300">{user.name}</span>
        <span className="text-sm text-blue-500">{user.email}</span>
      </div>
    </div>
  );
};

export default UserCard;
