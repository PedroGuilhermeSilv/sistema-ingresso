

import { EventModel } from "@/models";
import { EventCard } from "../components/EvendCard";
import { Title } from "../components/Title";

export async function getEvents(): Promise<EventModel[]> {
    const response = await fetch('http://localhost:8080/events', {
        cache: "no-store",


    })
    return (await response.json()).events;
}
export default async function HomePage() {
    const events: EventModel[] = await getEvents();
    return (
        <main className="mt-10 flex flex-col">
            <Title>Eventos disponíveis</Title>
            <div className="mt-8 sm:grid sm:grid-cols-auto-fit-cards flex flex-wrap justify-center gap-x-2 gap-y-4">
                {events.map((event) => (
                    <EventCard key={event.id} event={event} />
                ))}
            </div>
        </main>
    );
}