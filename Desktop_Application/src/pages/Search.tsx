import React, { useState, useMemo } from "react";
import SearchBar from "../components/SearchBar";
import UserCard from "../components/UserCard";

const USERS = [
  { id: "1", name: "Alice Johnson", email: "alice@test.com", online: true },
  { id: "2", name: "Bob Smith", email: "bob@test.com", online: false },
  { id: "3", name: "Charlie Brown", email: "charlie@test.com", online: true },
];

const UserSearch: React.FC = () => {
  const [query, setQuery] = useState("");

  const filteredUsers = useMemo(await () => {
    if (!query) return [];
    let res = []
    if (query.length >= 3) {
      console.log("Calling?")
      res = async window.api.search(query)
      console.log(res)
    }
    return res.filter((user: any) =>
      user.email.toLowerCase().includes(query.toLowerCase())
    );
  }, [query]);

  return (
    <div>
      <h2>Search For Users</h2>
      <div className="max-w-md mx-auto space-y-3">
        <SearchBar value={query} onChange={setQuery} />

        <div className="border rounded-lg divide-y">
          {filteredUsers.map((user: any) => (
            <UserCard key={user.id} user={user} />
          ))}

          {query && filteredUsers.length === 0 && (
            <div className="p-4 text-sm text-gray-500 text-center">
              No users found
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default UserSearch;
