import MainNavbar from "../components/MainNavbar";
import { privateRoute } from "../components/privateRoute";
import Link from "next/link";
import styles from "../styles/Home.module.css";

function Home() {
  return (
    <div className={styles.container}>
      <MainNavbar />
      <Link href="/login">
        <a>Login</a>
      </Link>
      <Link href="/form">
        <a>Request Instance</a>
      </Link>
    </div>
  );
}

export default privateRoute(Home);
