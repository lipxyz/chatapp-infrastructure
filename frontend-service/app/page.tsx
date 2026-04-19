"use client";
import Cookies from "js-cookie";
import {useEffect, useState} from "react";
import TweetComponent from "@/components/TweetComponent";
import LeftBarComponent from "@/components/LeftBarComponent";
import {useRouter, useSearchParams} from 'next/navigation'
import AlertComponent from "@/components/AlertComponent";

export default function Home() {
    const [tweets, setTweets] = useState<any[]>([]);
    const [tweetText, setTweetText] = useState("");
    const [isClient, setIsClient] = useState(false);
    const [dataSending, setDataSending] = useState(false)
    const token = Cookies.get("token");
    const username = Cookies.get("username");
    const userId = Cookies.get("userId");
    const searchParams = useSearchParams()
    const mode = searchParams.get('mode')
    const router = useRouter()

    useEffect(() => {
        setIsClient(true);
    }, []);

    useEffect(() => {

        if (isClient && mode !== 'guest' && !(username && token && userId)) {
            Cookies.set("mode", "guest", {expires: 7, path: "/"});
            if (Cookies.get("mode") == "guest") {
                return
            }
            router.push('/auth/sign-up')
        }
    }, [isClient]);

    // Запрос данных с API
    useEffect(() => {
        const fetchTweets = async () => {
            try {
                const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/tweets`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                if (!response.ok) throw new Error("Ошибка загрузки твитов");
                const data = await response.json();
                setTweets(data);
            } catch (error) {
                console.error("Ошибка при загрузке твитов:", error);
            }
        };
        fetchTweets();
    }, [token, dataSending]);

    // Функция отправки твита
    const handleTweetSubmit = async () => {
        if (!tweetText.trim()) return;
        try {
            const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/tweets`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify({
                    text: tweetText,
                    userId,
                    username,
                }),
            });

            if (!response.ok) throw new Error("Ошибка отправки твита");

            const newTweet = await response.json();
            setTweets(prevTweets => [newTweet, ...(prevTweets || [])]); // Добавляем новый твит в начало списка
            setDataSending(!dataSending)
            setTweetText(""); // Очищаем инпут
        } catch (error) {
            console.error("Ошибка при отправке твита:", error);
        }
    };
    const handleTweetUpdate = (tweetId, action) => {

        if (action === "slurm") {
            setTweets(prevTweets =>
                prevTweets.map(tweet =>
                    tweet.id === tweetId
                        ? {...tweet, [action + 'ed']: tweet[action + 'ed'] + 1}
                        : tweet
                )
            );
        } else {
            setTweets(prevTweets =>
                prevTweets.map(tweet =>
                    tweet.id === tweetId
                        ? {...tweet, [action]: tweet[action] + 1}
                        : tweet
                )
            );
        }

    };

    return (
        <>
            {mode == 'guest' || !(username && token && userId) ?
                <AlertComponent type={'error'}/> : null}
            <div className="relative flex size-full min-h-screen flex-col bg-slurm group/design-root overflow-x-hidden">
                <div className="layout-container flex h-full grow flex-col">
                    <div className="gap-1 px-6 flex flex-1 justify-center py-5">
                        <div className="layout-content-container flex flex-col w-80">
                            <LeftBarComponent/>
                        </div>
                        <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
                            <div className="flex flex-wrap justify-between gap-3 p-4">
                                <p className="text-[#2BD28A] tracking-light text-[32px] font-bold leading-tight min-w-72">
                                    SlurmTalks
                                </p>
                            </div>
                            <div className="flex items-center px-4 py-3 gap-3 @container">
                                <label className="flex flex-col min-w-40 h-full flex-1">
                                    <div className="flex w-full flex-1 items-stretch rounded-xl h-full">
                                        <div
                                            className="flex border border-[#dce0e5] bg-white justify-end pl-[15px] pr-[15px] pt-[15px] rounded-l-xl border-r-0">
                                            <div
                                                className="bg-center bg-no-repeat aspect-square bg-cover rounded-full size-10 shrink-0"></div>
                                        </div>
                                        <div className="flex flex-1 flex-col">
                                        <textarea
                                            disabled={!username || !token || !userId}
                                            placeholder={
                                                username && token && userId
                                                    ? "What happened? 👀"
                                                    : "You should register 😒"
                                            }
                                            className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-xl text-[#111418] focus:outline-0 focus:ring-0 border border-[#dce0e5] bg-white focus:border-[#dce0e5] h-auto placeholder:text-[#bbc4ce] rounded-l-none border-l-0  pr-16 rounded-b-none border-b-0 text-base font-normal leading-normal pt-[22px]"
                                            value={tweetText}
                                            onChange={(e) => setTweetText(e.target.value)}
                                            onKeyDown={(e) => {
                                                if (e.key === "Enter" && !e.shiftKey) {
                                                    e.preventDefault();
                                                    void handleTweetSubmit();
                                                }
                                            }}
                                        ></textarea>
                                            <div
                                                className="flex border border-[#dce0e5] bg-white justify-end pr-[15px] rounded-br-xl border-l-0 border-t-0 px-[15px] pb-[15px]">
                                                <button
                                                    onClick={handleTweetSubmit}
                                                >
                                                    <svg
                                                        xmlns="http://www.w3.org/2000/svg"
                                                        width="24px"
                                                        height="24px"
                                                        viewBox="0 0 24 24"
                                                        fill="currentColor"
                                                        className="transition-colors duration-200 hover:fill-green-500"
                                                    >
                                                        <path d="M2 12L22 2 15 22 11 13 2 12z"/>
                                                    </svg>
                                                </button>
                                            </div>
                                        </div>
                                    </div>
                                </label>
                            </div>

                            {/* Отображение твитов */}
                            {tweets?.map((tweet, index) =>
                                isClient ? (
                                    <TweetComponent
                                        key={tweet.id}
                                        id={tweet.id}  // Передаем ID твита
                                        username={tweet.username || "Anonymous"}
                                        tweetContent={tweet.text}
                                        retweets={tweet.share}
                                        likes={tweet.like}
                                        slurmed={tweet.slurmed}
                                        onUpdate={handleTweetUpdate}
                                    />

                                ) : null
                            )}

                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}
