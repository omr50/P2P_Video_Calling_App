import React, { useState, useMemo, useEffect } from "react";
import SearchBar from "../components/SearchBar";
import UserCard from "../components/UserCard";

const USERS = [
  { id: "1", name: "Alice Johnson", email: "alice@test.com", online: true },
  { id: "2", name: "Bob Smith", email: "bob@test.com", online: false },
  { id: "3", name: "Charlie Brown", email: "charlie@test.com", online: true },
];

const UserSearch: React.FC = () => {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState<any[]>([])


  useEffect(() => {
    if (!query || query.length < 3) {
        setResults([]);
        return;
    }

    const run = async () => {
      try {
        console.log("Calling?:", query);
        const res = await window.api.search(query);
        console.log(res);
        setResults(res);
      } catch (err) {
        console.error(err);
        setResults([]);
      }
    };

    run();
  }, [query]);

  return (
    <div>
      <h2>Search For Users</h2>
      <div className="max-w-md mx-auto space-y-3">
        <SearchBar value={query} onChange={setQuery} />

        <div className="border rounded-lg divide-y">
          {results.map((user: any) => (
            <UserCard key={user.id} user={user} />
          ))}

          {query && results.length === 0 && (
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
