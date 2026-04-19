import Image from "next/image";
import { useState } from "react";
import Cookies from "js-cookie";

export default function TweetComponent({ id, username, tweetContent, retweets, likes, slurmed, onUpdate }) {
    const [hasLiked, setHasLiked] = useState(false);
    const [hasSlurmed, setHasSlurmed] = useState(false);
    const [hasShared, setHasShared] = useState(false);

    const handleAction = async (action: "like" | "share" | "slurm") => {
        if ((action === "like" && hasLiked) || (action === "slurm" && hasSlurmed)) {
            return; // Не отправляем повторный запрос
        }

        try {
            const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/tweets/${id}/${action}`, {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${Cookies.get("token")}`,
                    "Content-Type": "application/json",
                },
            });

            if (!response.ok) throw new Error(`Ошибка при ${action}`);

            // Обновляем UI
            onUpdate(id, action);

            // Фиксируем состояние
            if (action === "like") setHasLiked(true);
            if (action === "slurm") setHasSlurmed(true);
            if (action === "share") setHasShared(true);
        } catch (error) {
            console.error(`Ошибка при ${action}:`, error);
        }
    };

    return (
        <div className="bg-white rounded-2xl shadow-lg border border-gray-200 p-4 mb-4 z-10">
            <div className="flex items-center justify-between mb-2">
                <div className="flex items-center gap-4">
                    <div className="flex flex-col">
                        <p className="text-[#111418] text-base font-medium leading-normal">{username}</p>
                        <p className="text-[#637588] text-sm font-normal">
                            {retweets} Retweets • {slurmed} Slurmed • {likes} Likes
                        </p>
                    </div>
                </div>
            </div>

            <p className="text-[#111418] text-base font-normal leading-normal mb-3 break-words">
                {tweetContent}
            </p>

            <div className="flex justify-between items-center border-t border-gray-200 pt-2">
                {/* SLURM */}
                <button onClick={() => handleAction("slurm")} className="flex items-center gap-2">
                    <Image src={'/ilstr/idle-slurm.gif'} alt={'gif'} width={24} height={24}/>
                    <p className={`text-sm font-bold ${hasSlurmed ? "text-green-500" : "text-[#637588]"}`}>{slurmed}</p>
                </button>

                {/* LIKE */}
                <button onClick={() => handleAction("like")} className="flex items-center gap-2">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24px" height="24px" fill="currentColor"
                         viewBox="0 0 256 256"
                         className={hasLiked ? "text-red-500" : "text-[#637588]"}
                    >
                        <path d="M178,32c-20.65,0-38.73,8.88-50,23.89C116.73,40.88,98.65,32,78,32A62.07,62.07,0,0,0,16,94c0,70,103.79,126.66,108.21,129a8,8,0,0,0,7.58,0C136.21,220.66,240,164,240,94A62.07,62.07,0,0,0,178,32ZM128,206.8C109.74,196.16,32,147.69,32,94A46.06,46.06,0,0,1,78,48c19.45,0,35.78,10.36,42.6,27a8,8,0,0,0,14.8,0c6.82-16.67,23.15-27,42.6-27a46.06,46.06,0,0,1,46,46C224,147.61,146.24,196.15,128,206.8Z"></path>
                    </svg>
                    <p className={`text-sm font-bold ${hasLiked ? "text-red-500" : "text-[#637588]"}`}>{likes}</p>
                </button>

                {/* SHARE */}
                <button onClick={() => handleAction("share")} className="flex items-center gap-2">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24px" height="24px" fill="currentColor"
                         viewBox="0 0 256 256"
                         className={hasShared ? "text-green-500" : "text-[#637588]"}
                    >
                        <path d="M229.66,109.66l-48,48a8,8,0,0,1-11.32-11.32L204.69,112H165a88,88,0,0,0-85.23,66,8,8,0,0,1-15.5-4A103.94,103.94,0,0,1,165,96h39.71L170.34,61.66a8,8,0,0,1,11.32-11.32l48,48A8,8,0,0,1,229.66,109.66ZM192,208H40V88a8,8,0,0,0-16,0V208a16,16,0,0,0,16,16H192a8,8,0,0,0,0-16Z"></path>
                    </svg>
                    <p className={`text-sm font-bold ${hasShared ? "text-green-500" : "text-[#637588]"}`}>{retweets}</p>
                </button>
            </div>
        </div>
    );
}
