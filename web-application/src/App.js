import UserSearch from "./UserSearch";
import { makeFetchHappen } from "./json-fetch";

import "./App.css";

export default function App() {
  async function onSearch(searchValue) {
    console.log({ searchValue });
    const users = await makeFetchHappen(searchValue);
    console.log({ users });
  }

  return (
    <div className="App">
      <header id="radical">
        <h1>Protobufs with web applications</h1>
      </header>
      <main>
        <UserSearch onSearch={onSearch} />
      </main>
    </div>
  );
}
