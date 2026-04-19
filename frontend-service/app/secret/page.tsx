'use client'

import Image from "next/image";

export default function secretPage() {

    return (
        <div className="flex justify-center items-center min-h-screen bg-slurm">
            <div className="flex w-full max-w-4xl p-12 rounded-lg shadow-lg bg-white">
                {/* Form Section (Left Side) */}
                <div className="w-1/2 flex flex-col justify-center space-y-6">
                    <div className="flex justify-center items-center">
                        <Image src="/thebestteamlead.png" alt="Profile"
                             className="object-cover  border-[#170F63]" width={256} height={256} />
                    </div>
                    <div className="text-center">
                        <h2 className="text-xl font-semibold text-[#170F63]">Вячеслав Федосеев</h2>
                        <p className="text-md text-[#637588]">TeamLead DevOps</p>
                    </div>
                </div>

                {/* Bio Section (Right Side) */}
                <div className="w-1/2 flex flex-col justify-center items-start space-y-4">
                    <h3 className="text-xl font-semibold text-[#170F63]">О себе</h3>
                    <p className="text-md text-[#637588] leading-relaxed">
                        Работал админом еще в 2007, с 2015 занимался организацией эксплуатации высоконагруженных
                        сервисов в энтерпрайзе, руководил инженерами. На себе почувствовал, зачем нужен DevOps и с 2020
                        TeamLead DevOps. Автор курсов, ментор и наставник для DevOps-инженеров.
                    </p>
                </div>
            </div>
        </div>

    );
}
