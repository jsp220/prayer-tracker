"use client";

import styles from "../sass/style.module.scss";
import { useState } from "react";

export default function Test() {
    const [loading, setLoading] = useState(false);
    const [data, setData] = useState("");

    const handleAmyFetch = async () => {
        setLoading(true);
        try {
            const response = await fetch("http://localhost:8080/amy"); // Calls backend
            const result = await response.json();
            console.log("Response from backend :", result);
            setData(result.message);
        } catch (error) {
            console.error("Error fetching data:", error);
            setData("Error fetching data.");
        } finally {
            setLoading(false);
        }
    };

    const handleEmmaFetch = async () => {
        setLoading(true);
        try {
            const response = await fetch("http://localhost:8080/emma"); // Calls backend
            const result = await response.json();
            console.log("Response from backend :", result);
            setData(result.message);
        } catch (error) {
            console.error("Error fetching data:", error);
            setData("Error fetching data.");
        } finally {
            setLoading(false);
        }
    };

    const handleJayFetch = async () => {
        setLoading(true);
        try {
            const response = await fetch("http://localhost:8080/jay"); // Calls backend
            const result = await response.json();
            console.log("Response from backend :", result);
            setData(result.message);
        } catch (error) {
            console.error("Error fetching data:", error);
            setData("Error fetching data.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <>
            <h2>Test Page</h2>
            <button onClick={handleAmyFetch} disabled={loading}>
                {loading ? "Loading..." : "Amy"}
            </button>
            <button onClick={handleEmmaFetch} disabled={loading}>
                {loading ? "Loading..." : "Emma"}
            </button>
            <button onClick={handleJayFetch} disabled={loading}>
                {loading ? "Loading..." : "Jay"}
            </button>
            <div className={styles.text}>{data}</div>
        </>
    );
}
