import styles from "../page.module.css";
import Link from "next/link";
import RegisterForm from "@/components/RegisterForm";

export default function RegisterPage() {
  return (
    <div className={styles.main}>
      <div className="relative">
        <RegisterForm />
        <Link
          href="/"
          className="absolute font-medium text-white right-0 hover:underline"
        >
          Back to login
        </Link>
      </div>
    </div>
  );
}
