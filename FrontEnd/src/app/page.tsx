"use client";
import { useState, FormEvent } from "react";
import { useRouter } from 'next/navigation';
import Image from "next/image";
import styles from "./page.module.css";

export default function Home() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const router = useRouter();

  const handleLogin = async (e: FormEvent<HTMLFormElement>): Promise<void> => {
    e.preventDefault();
    setError("");
    setSuccessMessage("");

    try {
      const response = await fetch("http://localhost:8080/api/v1/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      const data = await response.json();

      if (response.ok) {
        setSuccessMessage("Login successful!");
        localStorage.setItem("token", data.token);
        localStorage.setItem("userID", data.userID); // Save userID
        router.push("/user");
      } else {
        setError(data.message || "Login failed. Please check your credentials.");
      }
    } catch (error) {
      setError("An error occurred while logging in. Please try again later.");
    }
  };

  return (
    <div className={styles.main}>
      <div className="w-[400px] h-[400px] rounded-lg bg-blue-500 flex flex-col justify-center items-center shadow-2xl p-5">
        <Image
          src="/img/teenoilogo.jpg"
          alt="logo"
          width={100}
          height={100}
          priority
        />
        <h1 className="text-xl font-bold mb-4 text-white">Login</h1>
        <form
          onSubmit={handleLogin}
          className="flex flex-col w-full max-w-xs"
        >
          <input
            className="w-full h-[30px] mb-3 px-2 rounded-md text-sm"
            type="text"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
          <input
            className="w-full h-[30px] mb-4 px-2 rounded-md text-sm"
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <button
            type="submit"
            className="w-full h-[35px] bg-blue-400 rounded-md font-semibold transition-transform transform hover:scale-105 active:scale-95 focus:outline-none"
          >
            Login
          </button>
          {error && <div className="mt-3 text-red-400">{error}</div>}
          {successMessage && <div className="mt-3 text-green-400">{successMessage}</div>}
        </form>
      </div>
    </div>
  );
}
