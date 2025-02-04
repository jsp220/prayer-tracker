"use client"

import { useState } from "react";

export default function Test() {
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState("");

  const handleFetch = async () => {
    setLoading(true);
    try {
      const response = await fetch("http://localhost:8080/"); // Calls backend
      const result = await response.json();
      console.log("Response from backend:", result);
      setData(result.message);
    } catch (error) {
      console.error("Error fetching data:", error);
      setData("Error fetching data.")
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <h1>Test Page</h1>
      <button onClick={handleFetch} disabled={loading}>
        {loading ? "Loading..." : "Fetch Data"}
      </button>
      <div>{data}</div>
    </div>
  );
}