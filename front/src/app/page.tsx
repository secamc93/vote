"use client";
import { useRouter } from 'next/navigation';
import styles from "./page.module.css";

export default function Home() {
  const router = useRouter();

  const handleClick = () => {
    router.push("/group");
  };

  return (
    <button onClick={handleClick} className={styles.welcome}>
      Bienvenido a Voting
    </button>
  );
}
