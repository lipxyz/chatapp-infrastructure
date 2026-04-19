import {useState} from "react";
import Link from "next/link";

const alertVariants = {
    error: {icon: "⚠️", color: "bg-red-100 text-black"},
    info: {icon: "ℹ️", color: "bg-blue-100 text-blue-800"},
};

export default function AlertComponent({type = "info"}) {
    const [visible, setVisible] = useState(true);


    if (!visible) return null;

    const {icon, color} = alertVariants[type] || alertVariants.info;

    return (
        <div className={`flex items-center p-4 shadow-md ${color}`}>
            <span className="mr-3">{icon}</span>
            <span className="flex-1">
        <span>Вы сейчас в гостевом режиме. Выполните <Link href="/auth/sign-up"><span
            className={'underline text-blue-800'}>регистрацию</span></Link> или <Link href="/auth/sign-in"><span
            className={'underline text-blue-800'}>вход</span></Link></span>
      </span>
            <button onClick={() => setVisible(false)} className="ml-3">✖️</button>
        </div>
    );
}