'use client'
import {useEffect, useState} from "react";
import LeftBarComponent from "@/components/LeftBarComponent";
import Cookies from "js-cookie";
import AlertComponent from "@/components/AlertComponent";

export default function CommunityPage() {
    const [users, setUsers] = useState([]);
    const [isClient, setIsClient] = useState(false)

    useEffect(() => {
        setIsClient(true)
    }, [])

    useEffect(() => {
        fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/users`, {
            headers: {
                Authorization: `Bearer ${Cookies.get("token")}`
            }
        })
            .then(res => res.json())
            .then(data => setUsers(data))
            .catch(err => console.error("Ошибка загрузки пользователей", err));
    }, [isClient]);

    return (
        <>
            {Cookies.get("token") ? null : <AlertComponent type={'error'}/>}
            <div className="flex min-h-screen bg-gray-100">
                <LeftBarComponent/>
                <div className="flex-1 p-6">
                    <div className="flex flex-wrap justify-between gap-3 p-4"><p
                        className="text-[#2BD28A] tracking-light text-[32px] font-bold leading-tight min-w-72">Communities</p>
                    </div>
                    <div className="bg-white p-4 rounded-lg shadow">
                        <ul>
                            {
                                users?.length > 0 ? (
                                    users.map(user => (
                                        <li key={user.id} className="p-3 border-b last:border-b-0">
                                            <p className="text-lg text-black font-semibold">{user.username}</p>
                                            <p className="text-sm text-gray-500">Зарегистрирован: {new Date(user.created_at).toLocaleDateString()}</p>
                                        </li>
                                    ))
                                ) : 'Произошла ошибка, подробнее смотрите в "Network"'
                            }
                        </ul>
                    </div>
                </div>
            </div>
        </>
    );
}
