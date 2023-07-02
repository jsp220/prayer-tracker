import { useState } from "react";
import Header from "./Header";

export default function HomePage() {
  const names = ["John Doe", "Jane Smith", "Joon Park"];

  const [likes, setLikes] = useState(0);

  const handleClick = function(): void {
    setLikes(likes + 1);
  }

  return (
    <>
     <Header title="Test" />
     <ul>
      {names.map(name => (
        <li key={name}>{name}</li>
      ))}
     </ul>

     <button onClick={handleClick}>Like ({likes})</button>
    </>
  );
}