import { useState } from "react";

export default function Input({ onSearch }) {
  const [searchInput, setSearchInput] = useState("");

  return (
    <div style={{ display: "flex", gridGap: 12 }}>
      <input
        placeholder="User name"
        value={searchInput}
        onChange={(evt) => setSearchInput(evt.target.value)}
      />
      <button onClick={() => onSearch(searchInput)}>Search users</button>
      <button onClick={() => onSearch("")}>All users</button>
    </div>
  );
}
