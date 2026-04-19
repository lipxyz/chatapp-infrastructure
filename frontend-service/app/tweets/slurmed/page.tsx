"use client"
import Cookies from "js-cookie";
import {useEffect, useState} from "react";
import TweetComponent from "@/components/TweetComponent";
import LeftBarComponent from "@/components/LeftBarComponent";
import AlertComponent from "@/components/AlertComponent";

export default function SlurmedPage() {
    const [tweets, setTweets] = useState<any[]>([]);
    const [token, setToken] = useState<any>();
    const [isClient, setIsClient] = useState(false)

    useEffect(() => {
        setIsClient(true)
    }, [])
    // Запрос данных с API
    useEffect(() => {
        setToken(Cookies.get("token"))
        fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/tweets/sorted`, {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json",
            },
        })
            .then((res) => {
                if (!res.ok) {
                    throw new Error("Failed to fetch tweets");
                }
                return res.json();
            })
            .then((data) => setTweets(data))
            .catch((error) => console.error("Error fetching tweets:", error));
    }, [isClient]);

    const handleTweetUpdate = (tweetId, action) => {

        setTweets(prevTweets =>
            prevTweets.map(tweet =>
                tweet.id === tweetId
                    ? { ...tweet, [action]: tweet[action] + 1 }
                    : tweet
            )
        );
    };
    return (
        <div className="relative flex size-full min-h-screen flex-col bg-slurm group/design-root overflow-x-hidden"
        >
            {token ? null : <AlertComponent type={'error'}/> }
            <div className="layout-container flex h-full grow flex-col">
                <div className="gap-1 px-6 flex flex-1 justify-center py-5">
                    <LeftBarComponent/>
                    <div className="layout-content-container flex flex-col max-w-[960px] flex-1">
                        <div className="flex flex-wrap justify-between gap-3 p-4"><p
                            className="text-[#2BD28A] tracking-light text-[32px] font-bold leading-tight min-w-72">Slurmed</p>
                        </div>

                        {/*TODO: TWEETS*/}
                        {tweets ? null : <p className={'text-white'}>📣 Кажется нет постов.. Советую посмотреть на backend и Redis</p>}
                        {tweets?.map((tweet) => isClient ? (
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
                        ) : null)}
                    </div>
                </div>
            </div>
        </div>);
}
