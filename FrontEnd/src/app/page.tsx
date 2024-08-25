import Image from "next/image";
import styles from "./page.module.css";

export default function Home() {
  return (
    <div className={styles.main}>
    <div className="w-[400px] h-[400px] rounded-lg bg-blue-500 flex justify-center items-center shadow-2xl">
      <form
        action="/user"
        method="post"
        className="flex flex-col justify-center items-center"
      >
        <Image src="/img/teenoilogo.jpg" alt="logo" width={100} height={100} />
        <input
          className="w-[200px] h-[30px] my-2 px-2 rounded-md text-sm"
          type="text"
        />
        <input
          className="w-[200px] h-[30px] my-2 px-2 rounded-md text-sm"
          type="password"
        />
        <button className="w-[100px] h-[30px] bg-blue-300 rounded-md my-5 font-semibold">
          Login
        </button>
      </form>
    </div>
    </div>
  );
}
