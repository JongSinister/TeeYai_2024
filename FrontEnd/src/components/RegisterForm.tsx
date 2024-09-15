"use client";
import { useState, FormEvent } from "react";
import { useRouter } from "next/navigation";

export default function RegisterForm() {
  const [name, setUserName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const router = useRouter();

  const getRole = async (token: string) => {
    try {
      const response = await fetch("http://localhost:8080/api/v1/auth/me", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error("Failed to fetch user data.");
      }

      const userData = await response.json();
      return userData.Role;
    } catch (e) {
      console.error(e);
      return undefined;
    }
  };

  const handleRegister = async (
    e: FormEvent<HTMLFormElement>
  ): Promise<void> => {
    e.preventDefault();
    try {
      const response = await fetch(
        "http://localhost:8080/api/v1/auth/register",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ name, email, password, role: "user" }),
        }
      );

      const data = await response.json();
      if (response.ok) {
        console.log("User registered successfully!");
        localStorage.setItem("token", data.token);
        localStorage.setItem("userID", data.userID);
        const role = await getRole(data.token);
        if (role === "admin") {
          router.push("/admin");
        } else if (role === "user") {
          router.push("/user");
        } else {
          console.error("Invalid role.");
        }
      }

      console.log(data);
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <div className="w-[400px] h-[400px] rounded-lg bg-blue-500 flex flex-col justify-center items-center shadow-2xl p-5">
      <h1 className="text-xl font-bold mb-4 text-white">Register</h1>
      <form onSubmit={handleRegister} className="flex flex-col w-full max-w-xs">
        <input
          className="w-full h-[30px] mb-3 px-2 rounded-md text-sm"
          type="text"
          placeholder="Username"
          value={name}
          onChange={(e) => setUserName(e.target.value)}
          required
        />
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
          Register
        </button>
      </form>
    </div>
  );
}
