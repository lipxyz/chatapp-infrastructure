import Cookies from "js-cookie";
import Image from "next/image";
import Link from "next/link";
import {usePathname, useRouter} from "next/navigation";
import {useState} from "react";

export default function LeftBarComponent() {
    const [countSecret, setCountSecret] = useState(0)
    const router = useRouter()

    const addCountSecret = () => {
        if (countSecret === 1) {
            alert('Не трогай меня!')
        }
        if (countSecret === 3) {
            alert('Мне много еще раз повторять?')
        }
        if (countSecret === 4) {
            alert('Лучше бы с таким интересом изучал кубер!')
        }
        if (countSecret === 5) {
            router.push('/secret')
        }
        setCountSecret(countSecret + 1)
    }
    const pathUrl = usePathname();
    return (
        <div className="layout-content-container flex flex-col w-80 text-white">
            <div className="flex h-full min-h-[700px] flex-col justify-between bg-slurm p-4">
                <div className="flex flex-col gap-4">
                    <Link href={'/auth/sign-up'}>
                        <div
                            className="flex gap-4 p-4 bg-gradient-to-r from-blue-100 via-blue-50 to-white rounded-lg shadow-lg relative overflow-hidden">
                            <div className="absolute inset-0">
                                <div className="bubbles absolute top-0 left-0 w-full h-full z-10">
                                    <div className="bubble animation-delay-1"></div>
                                    <div className="bubble animation-delay-2"></div>
                                    <div className="bubble animation-delay-3"></div>
                                </div>
                            </div>
                            <div className="flex flex-col relative z-20">
                                <h1 className="text-[#111418] text-xl font-semibold leading-tight">
                                    {Cookies.get("username") || "@guest"}
                                </h1>
                                <p className="text-[#637588] text-md font-medium leading-snug">
                                    {Cookies.get("email") || "@guest"}
                                </p>
                            </div>
                        </div>
                    </Link>

                    <div className="flex flex-col gap-2 border rounded-xl">
                        {/* Home */}
                        <Link href={"/"}>
                            <div
                                className={`flex items-center gap-3 px-3 py-2 rounded-xl ${
                                    pathUrl === "/" ? "bg-[#f0f2f4] text-[#2BD28A] font-bold" : ""
                                }`}
                            >
                                <div className="text-[#111418]">
                                    <svg
                                        xmlns="http://www.w3.org/2000/svg"
                                        width="24px"
                                        height="24px"
                                        fill="currentColor"
                                        viewBox="0 0 256 256"
                                    >
                                        {
                                            pathUrl === "/" ? <path
                                                    d="M224,115.55V208a16,16,0,0,1-16,16H168a16,16,0,0,1-16-16V168a8,8,0,0,0-8-8H112a8,8,0,0,0-8,8v40a16,16,0,0,1-16,16H48a16,16,0,0,1-16-16V115.55a16,16,0,0,1,5.17-11.78l80-75.48.11-.11a16,16,0,0,1,21.53,0,1.14,1.14,0,0,0,.11.11l80,75.48A16,16,0,0,1,224,115.55Z"
                                                    fill={'#2BD28A'}
                                                /> :
                                                <path
                                                    d="M224,115.55V208a16,16,0,0,1-16,16H168a16,16,0,0,1-16-16V168a8,8,0,0,0-8-8H112a8,8,0,0,0-8,8v40a16,16,0,0,1-16,16H48a16,16,0,0,1-16-16V115.55a16,16,0,0,1,5.17-11.78l80-75.48.11-.11a16,16,0,0,1,21.53,0,1.14,1.14,0,0,0,.11.11l80,75.48A16,16,0,0,1,224,115.55Z"
                                                    fill="white"
                                                />}
                                    </svg>
                                </div>
                                <p className="text-sm leading-normal">Home</p>
                            </div>
                        </Link>

                        {/* Slurmed */}
                        <Link href={"/tweets/slurmed"}>
                            <div
                                className={`flex items-center gap-3 px-3 py-2 rounded-md ${
                                    pathUrl === "/tweets/slurmed" ? "bg-[#f0f2f4]" : ""
                                }`}
                            >
                                <Image
                                    src={"/ilstr/idle-slurm.gif"}
                                    alt={"slurm"}
                                    width={24}
                                    height={24}
                                />
                                <p
                                    className={`text-sm leading-normal ${
                                        pathUrl === "/tweets/slurmed"
                                            ? "font-bold text-[#2BD28A]"
                                            : ""
                                    }`}
                                >
                                    Slurmed
                                </p>
                            </div>
                        </Link>

                        {/* Communities */}
                        <Link href={"/communities"}>
                            <div
                                className={`flex items-center gap-3 px-3 py-2 rounded-md ${
                                    pathUrl === "/communities" ? "bg-[#f0f2f4]" : ""
                                }`}
                            >
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    width="24px"
                                    height="24px"
                                    fill="currentColor"
                                    viewBox="0 0 256 256"
                                >
                                    {
                                        pathUrl === "/communities" ? <path
                                                d="M117.25,157.92a60,60,0,1,0-66.5,0A95.83,95.83,0,0,0,3.53,195.63a8,8,0,1,0,13.4,8.74,80,80,0,0,1,134.14,0,8,8,0,0,0,13.4-8.74A95.83,95.83,0,0,0,117.25,157.92ZM40,108a44,44,0,1,1,44,44A44.05,44.05,0,0,1,40,108Zm210.14,98.7a8,8,0,0,1-11.07-2.33A79.83,79.83,0,0,0,172,168a8,8,0,0,1,0-16,44,44,0,1,0-16.34-84.87,8,8,0,1,1-5.94-14.85,60,60,0,0,1,55.53,105.64,95.83,95.83,0,0,1,47.22,37.71A8,8,0,0,1,250.14,206.7Z"
                                                fill={'#2BD28A'}/> :
                                            <path
                                                d="M117.25,157.92a60,60,0,1,0-66.5,0A95.83,95.83,0,0,0,3.53,195.63a8,8,0,1,0,13.4,8.74,80,80,0,0,1,134.14,0,8,8,0,0,0,13.4-8.74A95.83,95.83,0,0,0,117.25,157.92ZM40,108a44,44,0,1,1,44,44A44.05,44.05,0,0,1,40,108Zm210.14,98.7a8,8,0,0,1-11.07-2.33A79.83,79.83,0,0,0,172,168a8,8,0,0,1,0-16,44,44,0,1,0-16.34-84.87,8,8,0,1,1-5.94-14.85,60,60,0,0,1,55.53,105.64,95.83,95.83,0,0,1,47.22,37.71A8,8,0,0,1,250.14,206.7Z"/>
                                    }
                                </svg>
                                <p className={`text-sm leading-normal ${
                                    pathUrl === "/communities"
                                        ? "font-bold text-[#2BD28A]"
                                        : ""
                                }`}>Communities</p>
                            </div>
                        </Link>

                        {/* More */}
                        <div className="flex items-center gap-3 px-3 py-2">
                            <button onClick={addCountSecret}>
                                <Image src={'/ilstr/idle-slurm.gif'} alt={'ilstr'} width={24} height={24}/>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
            {
                pathUrl === '/communities' ? null : (
                    <>
                        <div className="absolute right-44 top-0 m-4">
                            <Image
                                src="/ilstr/donut.png"
                                alt="donut"
                                width={150} // размер изображения по ширине
                                height={150} // размер изображения по высоте
                            />
                        </div>
                        <div className="absolute left-52 top-96 m-4">
                            <Image
                                src="/ilstr/circle.png"
                                alt="circle"
                                width={150} // размер изображения по ширине
                                height={150} // размер изображения по высоте
                            />
                        </div>
                        <div className="absolute right-20 bottom-20 m-4">
                            <Image
                                src="/ilstr/slurms.svg"
                                alt="slurms-bike"
                                width={350} // размер изображения по ширине
                                height={350} // размер изображения по высоте
                            />
                        </div>
                    </>
                )
            }
        </div>
    );
}
