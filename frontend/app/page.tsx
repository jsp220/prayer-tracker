import Image from "next/image";
import styles from "./sass/style.module.scss";
import Link from "next/link";
import Button from "./components/Button";

export default function Home() {
    return (
        <main className={styles.main}>
            <Image
                className={styles.logo}
                src="/next.svg"
                alt="Next.js logo"
                width={180}
                height={38}
                priority
            />
            <ol>
                <li>
                    <Link href="/test">
                        <Button label="Test" className="nav-buttons" />
                    </Link>
                </li>
            </ol>

            <div className={styles.ctas}>
                <a
                    className={styles.primary}
                    href="https://vercel.com/new?utm_source=create-next-app&utm_medium=appdir-template&utm_campaign=create-next-app"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    <Image
                        className={styles.logo}
                        src="/vercel.svg"
                        alt="Vercel logomark"
                        width={20}
                        height={20}
                    />
                    Deploy now
                </a>
                <a
                    href="https://nextjs.org/docs?utm_source=create-next-app&utm_medium=appdir-template&utm_campaign=create-next-app"
                    target="_blank"
                    rel="noopener noreferrer"
                    className={styles.secondary}
                >
                    Read our docs
                </a>
            </div>
        </main>
    );
}
