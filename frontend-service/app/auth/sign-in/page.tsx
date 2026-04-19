'use client'
import {useState} from "react";
import Image from "next/image";
import {useRouter} from "next/navigation";
import Cookies from "js-cookie";
import Link from "next/link";

export default function SignInPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const router = useRouter()

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({username, password}),
            });

            if (!response.ok) {
                throw new Error("Failed to sign in");
            }

            const data = await response.json();  // Получаем данные из ответа

            // Сохраняем токен и userId в куки
            Cookies.set("token", data.token, {expires: 7, path: "/"});  // Кука с токеном, срок истечения - 7 дней
            Cookies.set("userId", data.userId, {expires: 7, path: "/"});  // Кука с userId
            Cookies.set("username", data.username, {expires: 7, path: "/"});  // Кука с userId
            Cookies.set("email", data.email, {expires: 7, path: "/"});  // Кука с userId

            // Успех - обработка редиректа или сообщения об успехе
            alert("User logged in successfully!");
            router.push("/");  // Переход на главную страницу
        } catch (error) {
            setError(error.message);
        }
    };

    return (
        <div className="flex justify-center items-center min-h-screen bg-slurm">
            <div className="flex w-full max-w-4xl p-12 rounded-lg">
                {/* Form Section (Left Side) */}
                <div className="w-1/2 flex flex-col justify-center">
                    <h2 className="text-4xl font-bold text-[#2BD28A] mb-8 text-center">Sign In</h2>

                    <form onSubmit={handleSubmit} className="w-full">
                        <div className="mb-4">
                            <label className="block text-[#2BD28A]" htmlFor="username">
                                Username
                            </label>
                            <input
                                type="text"
                                id="username"
                                className="text-white bg-inherit w-full p-3 border border-green-300 rounded-md"
                                value={username}
                                onChange={(e) => setUsername(e.target.value)}
                            />
                        </div>

                        <div className="mb-6">
                            <label className="block text-[#2BD28A]" htmlFor="password">
                                Password
                            </label>
                            <input
                                type="password"
                                id="password"
                                className="text-white bg-inherit w-full p-3 border border-green-300 rounded-md"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                        </div>

                        {error && <p className="text-red-500 mb-4">{error}</p>}

                        <div className="flex justify-between gap-4">
                            <Link href={'/auth/sign-up'}
                                  className="text-center text-[#2BD28A] w-1/2 border py-3 rounded-md hover:border-green-300 hover:bg-[#2BD28A] hover:text-white">
                                <button type="button"
                                >
                                    Register now!
                                </button>
                            </Link>

                            <button type="submit" className="w-1/2 bg-[#2BD28A] text-white py-3 rounded-md font-bold">
                                Let&apos;s dive
                            </button>
                        </div>
                        <div className="flex justify-between mx-auto gap-4 mt-5">
                            <Link href={'/'}
                                  className={'text-center w-full border text-white py-3 rounded-md hover:border-green-400 hover:text-white hover:bg-[#2BD28A]'}>
                                <button type="submit">
                                    Log in as a guest
                                </button>
                            </Link>
                        </div>
                    </form>
                </div>

                {/* Image Section (Right Side) */}
                <div className="w-1/2 flex justify-center items-center relative">
                    <Image
                        src="/ilstr/k-cap.png"
                        alt="k"
                        width={400} // Adjust size of the image
                        height={256} // Adjust size of the image
                        className="absolute top-1/2 transform -translate-y-1/2"
                    />
                </div>
            </div>
        </div>
    );
}
