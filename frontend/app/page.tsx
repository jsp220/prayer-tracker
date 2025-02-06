"use client";

import styles from "./sass/style.module.scss";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

export default function Home() {
    const router = useRouter();
    const [isLoading, setIsLoading] = useState(true);
    const [email, setEmail] = useState<string>("");
    const [warningMessage, setWarningMessage] = useState<string | null>(null);

    // Email validation regex
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    useEffect(() => {
        const storedEmail = localStorage.getItem("email");

        if (storedEmail) {
            setEmail(storedEmail);
            router.push("/prayers");
        } else {
            setIsLoading(false);
        }
    }, [router]);

    useEffect(() => {
        if (warningMessage) {
            const removeTimeout = setTimeout(() => {
                setWarningMessage(null);
            }, 6000);

            return () => {
                clearTimeout(removeTimeout);
            };
        }
    }, [warningMessage]);

    const handleSubmit = () => {
        if (!emailRegex.test(email)) {
            setWarningMessage("Please enter a valid email.");
            return;
        } else {
            localStorage.setItem("email", email);
            router.push("/prayers");
        }
    };

    if (isLoading) {
        return <h2>Loading...</h2>;
    }

    return (
        <div className={styles.loginPage}>
            <div className={warningMessage ? styles.warning : styles.hidden}>
                {warningMessage}
            </div>
            <input
                className={styles.loginInput}
                type="email"
                placeholder="Enter your email"
                onChange={(e) => setEmail(e.target.value)}
            />
            <button className={styles.loginButton} onClick={handleSubmit}>
                Login
            </button>
        </div>
    );
}
